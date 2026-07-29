package repo

import (
	"context"
	"scheduler/internal/biz/model"
)

type DelayedTaskVersionRepo interface {
	Create(ctx context.Context, task *model.DelayedTask) (*model.DelayedTaskVersion, error)
	Get(ctx context.Context, req *DelayedTaskVersionGetReq) (*model.DelayedTaskVersion, error)
	List(ctx context.Context, req *DelayedTaskVersionGetReq) ([]*model.DelayedTaskVersion, error)
}

type DelayedTaskVersionGetReq struct {
	ID             *int64
	IDs            []int64
	DelayedTaskID  *int64
	DelayedTaskIDs []int64
	Version        *int64
	Versions       []int64
}
