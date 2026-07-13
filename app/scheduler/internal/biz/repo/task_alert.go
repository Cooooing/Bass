package repo

import (
	"context"
	"scheduler/internal/biz/model"
)

type TaskAlert interface {
	Alert(ctx context.Context, task *model.Task, record *model.TaskExecutionRecord, reason string) error
}
