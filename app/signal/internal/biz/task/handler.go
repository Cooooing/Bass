package task

import (
	"context"

	"github.com/hibiken/asynq"
)

type Handler interface {
	Name() TaskName
	Handler() asynq.HandlerFunc
	ErrHandler(ctx context.Context, task *asynq.Task, err error)
}
