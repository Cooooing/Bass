package repo

import (
	"context"
	"scheduler/internal/biz/model"
)

type DelayedTaskCacheRepo interface {
	GetDelayedTask(ctx context.Context, title string) (*model.DelayedTask, error)
	SetDelayedTask(ctx context.Context, row *model.DelayedTask) error
	DeleteDelayedTask(ctx context.Context, title string) error
}
