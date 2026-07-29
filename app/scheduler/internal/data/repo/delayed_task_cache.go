package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"errors"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"time"

	"github.com/redis/go-redis/v9"
)

type DelayedTaskCacheRepo struct {
	redis *commonclient.RedisClient
	key   string
	ttl   time.Duration
}

func NewDelayedTaskCacheRepo(
	redisClient *commonclient.RedisClient,
) bizrepo.DelayedTaskCacheRepo {
	return &DelayedTaskCacheRepo{
		redis: redisClient,
		key:   "Scheduler:DelayedTasks",
		ttl:   24 * time.Hour,
	}
}

func (r *DelayedTaskCacheRepo) GetDelayedTask(ctx context.Context, title string) (*model.DelayedTask, error) {
	value, err := r.redis.Client.HGet(ctx, r.key, title).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	row := &model.DelayedTask{}
	if err = json.Unmarshal([]byte(value), row); err != nil {
		return nil, err
	}
	if row.Timeout <= 0 {
		_ = r.redis.Client.HDel(ctx, r.key, title).Err()
		return nil, nil
	}
	return row, nil
}

func (r *DelayedTaskCacheRepo) SetDelayedTask(ctx context.Context, row *model.DelayedTask) error {
	data, err := json.Marshal(row)
	if err != nil {
		return err
	}
	if err = r.redis.Client.HSet(ctx, r.key, row.Title, data).Err(); err != nil {
		return err
	}
	return r.redis.Client.Expire(ctx, r.key, r.ttl).Err()
}

func (r *DelayedTaskCacheRepo) DeleteDelayedTask(ctx context.Context, title string) error {
	return r.redis.Client.HDel(ctx, r.key, title).Err()
}
