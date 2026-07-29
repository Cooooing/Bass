package repo

import (
	"context"
	"scheduler/internal/biz/model"
)

type ScheduledTaskVersionRepo interface {
	Create(ctx context.Context, task *model.ScheduledTask) (*model.ScheduledTaskVersion, error)
	Get(ctx context.Context, req *ScheduledTaskVersionGetReq) (*model.ScheduledTaskVersion, error)
	List(ctx context.Context, req *ScheduledTaskVersionGetReq) ([]*model.ScheduledTaskVersion, error)
}

type ScheduledTaskVersionGetReq struct {
	ID               *int64
	IDs              []int64
	ScheduledTaskID  *int64
	ScheduledTaskIDs []int64
	Version          *int64
	Versions         []int64
}
