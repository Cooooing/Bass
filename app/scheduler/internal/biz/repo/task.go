package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
)

type TaskRepo interface {
	Get(ctx context.Context, req *TaskGetReq) (*TaskGetResponse, error)
	List(ctx context.Context, req *TaskGetReq) (*TaskListResponse, error)
	Map(ctx context.Context, req *TaskGetReq) (*TaskMapResponse, error)
	Count(ctx context.Context, req *TaskGetReq) (*TaskCountResponse, error)
	Page(ctx context.Context, req *TaskPageReq) (*TaskPageResponse, error)

	Upsert(ctx context.Context, req *TaskUpsertReq) (*TaskUpsertResponse, error)
	Lock(ctx context.Context, req *TaskLockReq) (*TaskLockResponse, error)
}

type TaskGetReq struct {
	ID      *int64
	IDs     []int64
	Name    *string
	Title   *string
	Enabled *bool
}

type TaskGetResponse struct {
	Row *model.Task
}

type TaskListResponse struct {
	Rows []*model.Task
}

type TaskMapResponse struct {
	Rows map[int64]*model.Task
}

type TaskCountResponse struct {
	Count int
}

type TaskPageReq struct {
	Page *common.PageRequest
	TaskGetReq
}

type TaskPageResponse struct {
	Rows []*model.Task
	Page *common.PageResponse
}

type TaskUpsertReq struct {
	Row *model.Task
}

type TaskUpsertResponse struct {
	Row *model.Task
}

type TaskLockReq struct {
	ID int64
}

type TaskLockResponse struct{}
