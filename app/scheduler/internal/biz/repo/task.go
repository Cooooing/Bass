package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
)

type TaskRepo interface {
	Get(ctx context.Context, req *TaskGetReq) (*model.Task, error)
	List(ctx context.Context, req *TaskGetReq) ([]*model.Task, error)
	Map(ctx context.Context, req *TaskGetReq) (map[int64]*model.Task, error)
	Count(ctx context.Context, req *TaskGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *TaskGetReq) ([]*model.Task, *common.PageReply, error)

	Upsert(ctx context.Context, row *model.Task) (*model.Task, error)
	Lock(ctx context.Context, id int64) error
}

type TaskGetReq struct {
	ID      *int64
	IDs     []int64
	Name    *string
	Title   *string
	Enabled *bool
}
