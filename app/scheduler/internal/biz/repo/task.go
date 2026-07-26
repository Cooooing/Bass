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
	Page(ctx context.Context, req *TaskPageReq) (*TaskPageResp, error)

	MapByTitle(ctx context.Context, titles []string) (map[string]*model.Task, error)
	Upsert(ctx context.Context, row *model.Task) (*model.Task, error)
	Lock(ctx context.Context, id int64) error
}

type TaskGetReq struct {
	ID      *int64
	IDs     []int64
	Name    *string
	Title   *string
	Titles  []string
	Enabled *bool
}

type TaskPageReq struct {
	Page *common.PageReq
	TaskGetReq
}

type TaskPageResp struct {
	Rows []*model.Task
	Page *common.PageResp
}
