package repo

import "context"

type TaskChangedMessage struct {
	TaskID  int64 `json:"task_id"`
	Version int64 `json:"version"`
}

type TaskExecutionCanceledMessage struct {
	ExecutionRecordID int64 `json:"execution_record_id"`
}

type TaskEventBus interface {
	PublishTaskChanged(ctx context.Context, msg *TaskChangedMessage) error
	PublishExecutionCanceled(ctx context.Context, msg *TaskExecutionCanceledMessage) error
	SubscribeTaskChanged(ctx context.Context) (<-chan TaskChangedMessage, error)
	SubscribeExecutionCanceled(ctx context.Context) (<-chan TaskExecutionCanceledMessage, error)
}
