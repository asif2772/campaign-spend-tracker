package config

import "os"

type Config struct {
	Port          string
	PostgresURL   string
	RedisAddr     string
	RedisPassword string
}

func Load() Config {
	return Config{
		Port:          getEnv("PORT", "8080"),
		PostgresURL:   getEnv("POSTGRES_URL", "postgres://postgres:postgres@localhost:5432/campaign?sslmode=disable"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
