package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
)

type TaskVersionRepo interface {
	Get(ctx context.Context, req *TaskVersionGetReq) (*model.TaskVersion, error)
	List(ctx context.Context, req *TaskVersionGetReq) ([]*model.TaskVersion, error)
	Map(ctx context.Context, req *TaskVersionGetReq) (map[int64]*model.TaskVersion, error)
	Count(ctx context.Context, req *TaskVersionGetReq) (int, error)
	Page(ctx context.Context, req *TaskVersionPageReq) (*TaskVersionPageResp, error)

	Create(ctx context.Context, task *model.Task) (*model.TaskVersion, error)
}

type TaskVersionGetReq struct {
	ID       *int64
	IDs      []int64
	TaskID   *int64
	TaskIDs  []int64
	Version  *int64
	Versions []int64
}

type TaskVersionPageReq struct {
	Page *common.PageReq
	TaskVersionGetReq
}

type TaskVersionPageResp struct {
	Rows []*model.TaskVersion
	Page *common.PageResp
}
