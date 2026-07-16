package repo

import (
	"context"
	"scheduler/internal/biz/model"
)

type TaskAlert interface {
	Alert(ctx context.Context, req *TaskAlertReq) (*TaskAlertResponse, error)
}

type TaskAlertReq struct {
	Task   *model.Task
	Record *model.TaskExecutionRecord
	Reason string
}

type TaskAlertResponse struct{}
