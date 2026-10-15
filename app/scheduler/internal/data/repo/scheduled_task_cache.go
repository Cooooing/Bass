package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"errors"
	"time"

	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

type ScheduledTaskCacheRepo struct {
	redis *commonclient.RedisClient
	key   string
	ttl   time.Duration
}

func NewScheduledTaskCacheRepo(redisClient *commonclient.RedisClient) bizrepo.ScheduledTaskCacheRepo {
	return &ScheduledTaskCacheRepo{redis: redisClient, key: "scheduler:scheduled_task:definitions", ttl: 24 * time.Hour}
}

func (r *ScheduledTaskCacheRepo) GetScheduledTask(ctx context.Context, handlerName string) (*model.ScheduledTask, error) {
	value, err := r.redis.Client.HGet(ctx, r.key, handlerName).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	row := &model.ScheduledTask{}
	if err = json.Unmarshal([]byte(value), row); err != nil {
		return nil, err
	}
	if row.Timeout <= 0 {
		_ = r.redis.Client.HDel(ctx, r.key, handlerName).Err()
		return nil, nil
	}
	return row, nil
}

func (r *ScheduledTaskCacheRepo) SetScheduledTask(ctx context.Context, row *model.ScheduledTask) error {
	data, err := json.Marshal(row)
	if err != nil {
		return err
	}
	if err = r.redis.Client.HSet(ctx, r.key, row.HandlerName.String(), data).Err(); err != nil {
		return err
	}
	return r.redis.Client.Expire(ctx, r.key, r.ttl).Err()
}

func (r *ScheduledTaskCacheRepo) DeleteScheduledTask(ctx context.Context, handlerName string) error {
	return r.redis.Client.HDel(ctx, r.key, handlerName).Err()
}
