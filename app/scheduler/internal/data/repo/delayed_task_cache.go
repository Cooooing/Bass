package repo

import (
	commonclient "common/pkg/client"
	"context"
	"time"

	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
)

type DelayedTaskCacheRepo struct {
	cache *taskCache[*model.DelayedTask]
}

func NewDelayedTaskCacheRepo(redisClient *commonclient.RedisClient) bizrepo.DelayedTaskCacheRepo {
	return &DelayedTaskCacheRepo{cache: &taskCache[*model.DelayedTask]{redis: redisClient, key: "scheduler:delayed_task:definitions", ttl: 24 * time.Hour, new: func() *model.DelayedTask { return &model.DelayedTask{} }, keyOf: func(row *model.DelayedTask) string { return row.TaskKey }, valid: func(row *model.DelayedTask) bool { return row.Timeout > 0 }}}
}

func (r *DelayedTaskCacheRepo) GetDelayedTask(ctx context.Context, taskKey string) (*model.DelayedTask, error) {
	return r.cache.get(ctx, taskKey)
}

func (r *DelayedTaskCacheRepo) SetDelayedTask(ctx context.Context, row *model.DelayedTask) error {
	return r.cache.set(ctx, row)
}

func (r *DelayedTaskCacheRepo) DeleteDelayedTask(ctx context.Context, taskKey string) error {
	return r.cache.redis.Client.HDel(ctx, r.cache.key, taskKey).Err()
}
