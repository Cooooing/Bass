package biz

import (
	"context"
	"encoding/json"
	"signal/internal/biz/base"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/biz/task"

	"github.com/hibiken/asynq"
)

type NodePingTaskHandler struct {
	*base.BaseDomain
	nodeDomain *NodeDomain
	nodeRepo   repo.NodeRepo
	producer   *Producer
}

func NewNodePingTaskHandler(baseDomain *base.BaseDomain, nodeDomain *NodeDomain, nodeRepo repo.NodeRepo, producer *Producer) *NodePingTaskHandler {
	return &NodePingTaskHandler{
		BaseDomain: baseDomain,
		nodeDomain: nodeDomain,
		nodeRepo:   nodeRepo,
		producer:   producer,
	}
}

func (h *NodePingTaskHandler) Name() task.TaskName {
	return task.TaskNodePing
}

func (h *NodePingTaskHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		data := new(model.Task)
		err := json.Unmarshal(task.Payload(), data)
		if err != nil {
			return err
		}
		node := new(model.Node)
		if err := json.Unmarshal(data.Data, node); err != nil {
			return err
		}
		pingMs, err := h.nodeDomain.Ping(node)
		if err != nil {
			return err
		}
		err = h.nodeRepo.UpdatePing(ctx, node.Key, pingMs)
		if err != nil {
			return err
		}
		data.Delay = true
		err = h.producer.EnqueueTask(data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *NodePingTaskHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {
	h.Log.Errorf("task %s failed: %v", task.Type(), err)
}
