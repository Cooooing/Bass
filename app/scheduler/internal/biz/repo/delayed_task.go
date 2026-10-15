package repo

import (
	commonmodel "common/pkg/model"
	"context"

	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
)

type DelayedTaskRepo interface {
	Get(ctx context.Context, req *DelayedTaskGetReq) (*model.DelayedTask, error)
	List(ctx context.Context, req *DelayedTaskGetReq) ([]*model.DelayedTask, error)
	Page(ctx context.Context, req *DelayedTaskPageReq) (*DelayedTaskPageResp, error)
	MapByHandlerName(ctx context.Context, handlerNames []schedulerenum.TaskHandlerName) (map[schedulerenum.TaskHandlerName]*model.DelayedTask, error)
	Upsert(ctx context.Context, row *model.DelayedTask) (*model.DelayedTask, error)
	Lock(ctx context.Context, id int64) error
}

type DelayedTaskGetReq struct {
	ID           *int64
	IDs          []int64
	HandlerNames []schedulerenum.TaskHandlerName
	HandlerName  *schedulerenum.TaskHandlerName
	Title        *string
	Enabled      *bool
}

type DelayedTaskPageReq struct {
	Page *commonmodel.PageReq
	DelayedTaskGetReq
}

type DelayedTaskPageResp struct {
	Rows []*model.DelayedTask
	Page *commonmodel.PageResp
}
