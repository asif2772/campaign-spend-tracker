package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) IncrementDailySpend(
	ctx context.Context,
	tenantID int64,
	campaignID int64,
	amountMicros int64,
) (int64, error) {

	key := buildDailySpendKey(
		tenantID,
		campaignID,
		time.Now().UTC(),
	)

	total, err := r.client.IncrBy(
		ctx,
		key,
		amountMicros,
	).Result()

	if err != nil {
		return 0, fmt.Errorf("increment daily spend: %w", err)
	}

	if total == amountMicros {
		if err := r.client.Expire(
			ctx,
			key,
			48*time.Hour,
		).Err(); err != nil {
			return 0, fmt.Errorf("set redis ttl: %w", err)
		}
	}

	return total, nil
}

func (r *RedisRepository) GetDailySpend(
	ctx context.Context,
	tenantID int64,
	campaignID int64,
) (int64, error) {

	key := buildDailySpendKey(
		tenantID,
		campaignID,
		time.Now().UTC(),
	)

	total, err := r.client.Get(ctx, key).Int64()

	if err == redis.Nil {
		return 0, nil
	}

	if err != nil {
		return 0, fmt.Errorf("get daily spend: %w", err)
	}

	return total, nil
}

func (r *RedisRepository) Health(
	ctx context.Context,
) error {

	if err := r.client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis health check: %w", err)
	}

	return nil
}
