package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asif2772/campaign-spend-tracker/internal/handler"
	"github.com/asif2772/campaign-spend-tracker/internal/publisher"
	"github.com/asif2772/campaign-spend-tracker/internal/repository"
	"github.com/asif2772/campaign-spend-tracker/internal/service"
	"github.com/asif2772/campaign-spend-tracker/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func main() {

	cfg := config.Load()
	fmt.Println("Postgres URL:", cfg.PostgresURL)
	fmt.Println("Redis Addr:", cfg.RedisAddr)

	ctx := context.Background()

	// PostgreSQL

	pool, err := pgxpool.New(ctx, cfg.PostgresURL)

	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}

	fmt.Println("PostgreSQL connected")

	defer pool.Close()

	// Redis

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis connection failed: %v", err)
	}

	fmt.Println("Redis connected")

	defer redisClient.Close()

	// Publisher

	pub := publisher.New(100)

	pub.Start()

	defer func() {

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)

		defer cancel()

		_ = pub.Shutdown(shutdownCtx)

	}()

	// Repository

	postgresRepository := repository.NewPostgresRepository(pool)

	redisRepository := repository.NewRedisRepository(redisClient)

	// Service

	spendService := service.NewSpendService(

		postgresRepository,

		redisRepository,

		pub,
	)

	// Handler

	spendHandler := handler.NewSpendHandler(spendService)

	healthHandler := handler.NewHealthHandler(

		postgresRepository,

		redisRepository,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthHandler.Health)

	mux.HandleFunc("/spend", spendHandler.CreateSpend)

	mux.HandleFunc("/spend/", spendHandler.GetDailySpend)

	server := &http.Server{

		Addr: ":" + cfg.Port,

		Handler: mux,
	}

	go func() {

		fmt.Printf("Server started on :%s\n", cfg.Port)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatal(err)

		}

	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(

		stop,

		syscall.SIGINT,

		syscall.SIGTERM,
	)

	<-stop

	shutdownCtx, cancel := context.WithTimeout(

		context.Background(),

		10*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {

		log.Println(err)

	}

	fmt.Println("shutdown complete")

}
