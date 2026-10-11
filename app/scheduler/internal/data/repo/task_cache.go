package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type taskCache[T any] struct {
	redis *commonclient.RedisClient
	key   string
	ttl   time.Duration
	new   func() T
	keyOf func(T) string
	valid func(T) bool
}

func (r *taskCache[T]) get(ctx context.Context, taskKey string) (T, error) {
	var zero T
	value, err := r.redis.Client.HGet(ctx, r.key, taskKey).Result()
	if errors.Is(err, redis.Nil) {
		return zero, nil
	}
	if err != nil {
		return zero, err
	}
	row := r.new()
	if err = json.Unmarshal([]byte(value), row); err != nil {
		return zero, err
	}
	if !r.valid(row) {
		_ = r.redis.Client.HDel(ctx, r.key, taskKey).Err()
		return zero, nil
	}
	return row, nil
}

func (r *taskCache[T]) set(ctx context.Context, row T) error {
	data, err := json.Marshal(row)
	if err != nil {
		return err
	}
	if err = r.redis.Client.HSet(ctx, r.key, r.keyOf(row), data).Err(); err != nil {
		return err
	}
	return r.redis.Client.Expire(ctx, r.key, r.ttl).Err()
}
