package task

import (
	"context"

	"github.com/hibiken/asynq"
)

type NodePingTaskHandler struct {
}

func NewNodePingTaskHandler() *NodePingTaskHandler {
	return &NodePingTaskHandler{}
}

func (h *NodePingTaskHandler) Name() TaskName {
	return TaskNodePing
}

func (h *NodePingTaskHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		return nil
	}
}
