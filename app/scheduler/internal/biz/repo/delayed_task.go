package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
)

type DelayedTaskRepo interface {
	Get(ctx context.Context, req *DelayedTaskGetReq) (*model.DelayedTask, error)
	List(ctx context.Context, req *DelayedTaskGetReq) ([]*model.DelayedTask, error)
	Page(ctx context.Context, req *DelayedTaskPageReq) (*DelayedTaskPageResp, error)
	MapByTaskKey(ctx context.Context, taskKeys []string) (map[string]*model.DelayedTask, error)
	Upsert(ctx context.Context, row *model.DelayedTask) (*model.DelayedTask, error)
	Lock(ctx context.Context, id int64) error
}

type DelayedTaskGetReq struct {
	ID          *int64
	IDs         []int64
	TaskKey     *string
	TaskKeys    []string
	HandlerName *schedulerenum.TaskHandlerName
	Title       *string
	Enabled     *bool
}

type DelayedTaskPageReq struct {
	Page *common.PageReq
	DelayedTaskGetReq
}

type DelayedTaskPageResp struct {
	Rows []*model.DelayedTask
	Page *common.PageResp
}
