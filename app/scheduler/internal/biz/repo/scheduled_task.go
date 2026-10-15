package repo

import (
	commonmodel "common/pkg/model"
	"context"

	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
)

type ScheduledTaskRepo interface {
	Get(ctx context.Context, req *ScheduledTaskGetReq) (*model.ScheduledTask, error)
	List(ctx context.Context, req *ScheduledTaskGetReq) ([]*model.ScheduledTask, error)
	Page(ctx context.Context, req *ScheduledTaskPageReq) (*ScheduledTaskPageResp, error)
	MapByHandlerName(ctx context.Context, handlerNames []schedulerenum.TaskHandlerName) (map[schedulerenum.TaskHandlerName]*model.ScheduledTask, error)
	Upsert(ctx context.Context, row *model.ScheduledTask) (*model.ScheduledTask, error)
	Lock(ctx context.Context, id int64) error
}

type ScheduledTaskGetReq struct {
	ID           *int64
	IDs          []int64
	HandlerNames []schedulerenum.TaskHandlerName
	HandlerName  *schedulerenum.TaskHandlerName
	Title        *string
	Enabled      *bool
}

type ScheduledTaskPageReq struct {
	Page *commonmodel.PageReq
	ScheduledTaskGetReq
}

type ScheduledTaskPageResp struct {
	Rows []*model.ScheduledTask
	Page *commonmodel.PageResp
}
