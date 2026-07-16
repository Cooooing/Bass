package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
)

type TaskVersionRepo interface {
	Get(ctx context.Context, req *TaskVersionGetReq) (*TaskVersionGetResponse, error)
	List(ctx context.Context, req *TaskVersionGetReq) (*TaskVersionListResponse, error)
	Map(ctx context.Context, req *TaskVersionGetReq) (*TaskVersionMapResponse, error)
	Count(ctx context.Context, req *TaskVersionGetReq) (*TaskVersionCountResponse, error)
	Page(ctx context.Context, req *TaskVersionPageReq) (*TaskVersionPageResponse, error)

	Create(ctx context.Context, req *TaskVersionCreateReq) (*TaskVersionCreateResponse, error)
}

type TaskVersionGetReq struct {
	ID       *int64
	IDs      []int64
	TaskID   *int64
	TaskIDs  []int64
	Version  *int64
	Versions []int64
}

type TaskVersionGetResponse struct {
	Row *model.TaskVersion
}

type TaskVersionListResponse struct {
	Rows []*model.TaskVersion
}

type TaskVersionMapResponse struct {
	Rows map[int64]*model.TaskVersion
}

type TaskVersionCountResponse struct {
	Count int
}

type TaskVersionPageReq struct {
	Page *common.PageRequest
	TaskVersionGetReq
}

type TaskVersionPageResponse struct {
	Rows []*model.TaskVersion
	Page *common.PageResponse
}

type TaskVersionCreateReq struct {
	Task *model.Task
}

type TaskVersionCreateResponse struct {
	Row *model.TaskVersion
}
