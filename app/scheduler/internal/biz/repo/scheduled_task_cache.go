package repo

import (
	"context"
	"scheduler/internal/biz/model"
)

type ScheduledTaskCacheRepo interface {
	GetScheduledTask(ctx context.Context, title string) (*model.ScheduledTask, error)
	SetScheduledTask(ctx context.Context, row *model.ScheduledTask) error
	DeleteScheduledTask(ctx context.Context, title string) error
}
