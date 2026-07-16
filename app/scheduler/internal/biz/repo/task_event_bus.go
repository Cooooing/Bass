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
	PublishTaskChanged(ctx context.Context, req *PublishTaskChangedReq) (*PublishTaskChangedResponse, error)
	PublishExecutionCanceled(ctx context.Context, req *PublishExecutionCanceledReq) (*PublishExecutionCanceledResponse, error)
	SubscribeTaskChanged(ctx context.Context, req *SubscribeTaskChangedReq) (*SubscribeTaskChangedResponse, error)
	SubscribeExecutionCanceled(ctx context.Context, req *SubscribeExecutionCanceledReq) (*SubscribeExecutionCanceledResponse, error)
}

type PublishTaskChangedReq struct {
	Message *TaskChangedMessage
}

type PublishTaskChangedResponse struct{}

type PublishExecutionCanceledReq struct {
	Message *TaskExecutionCanceledMessage
}

type PublishExecutionCanceledResponse struct{}

type SubscribeTaskChangedReq struct{}

type SubscribeTaskChangedResponse struct {
	Messages <-chan TaskChangedMessage
}

type SubscribeExecutionCanceledReq struct{}

type SubscribeExecutionCanceledResponse struct {
	Messages <-chan TaskExecutionCanceledMessage
}
