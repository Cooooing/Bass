package repo

import (
	commonclient "common/pkg/client"
	"context"
	"time"

	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
)

type ScheduledTaskCacheRepo struct {
	cache *taskCache[*model.ScheduledTask]
}

func NewScheduledTaskCacheRepo(redisClient *commonclient.RedisClient) bizrepo.ScheduledTaskCacheRepo {
	return &ScheduledTaskCacheRepo{cache: &taskCache[*model.ScheduledTask]{redis: redisClient, key: "scheduler:scheduled_task:definitions", ttl: 24 * time.Hour, new: func() *model.ScheduledTask { return &model.ScheduledTask{} }, keyOf: func(row *model.ScheduledTask) string { return row.TaskKey }, valid: func(row *model.ScheduledTask) bool { return row.Timeout > 0 }}}
}

func (r *ScheduledTaskCacheRepo) GetScheduledTask(ctx context.Context, taskKey string) (*model.ScheduledTask, error) {
	return r.cache.get(ctx, taskKey)
}

func (r *ScheduledTaskCacheRepo) SetScheduledTask(ctx context.Context, row *model.ScheduledTask) error {
	return r.cache.set(ctx, row)
}

func (r *ScheduledTaskCacheRepo) DeleteScheduledTask(ctx context.Context, taskKey string) error {
	return r.cache.redis.Client.HDel(ctx, r.cache.key, taskKey).Err()
}
