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

type NodePowTaskHandler struct {
	*base.BaseDomain
	nodeDomain *NodeDomain
	nodeRepo   repo.NodeRepo
	producer   *Producer
}

func NewNodePowTaskHandler(baseDomain *base.BaseDomain, nodeDomain *NodeDomain, nodeRepo repo.NodeRepo, producer *Producer) *NodePowTaskHandler {
	return &NodePowTaskHandler{
		BaseDomain: baseDomain,
		nodeDomain: nodeDomain,
		nodeRepo:   nodeRepo,
		producer:   producer,
	}
}

func (h *NodePowTaskHandler) Name() task.TaskName {
	return task.TaskNodePow
}

func (h *NodePowTaskHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		h.Log.Infof("task %s: %s", task.Type(), string(task.Payload()))
		data := new(model.Task)
		err := json.Unmarshal(task.Payload(), data)
		if err != nil {
			return err
		}
		node := new(model.Node)
		if err := json.Unmarshal(data.Data, node); err != nil {
			return err
		}

		// 判断节点是否在线
		ok, err := h.nodeRepo.IsOnline(ctx, node.Key)
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
		// pow
		powMs, err := h.nodeDomain.Pow(node)
		if err != nil {
			return err
		}
		err = h.nodeRepo.UpdatePowCost(ctx, node.Key, powMs)
		if err != nil {
			return err
		}
		err = h.nodeRepo.UpdateScore(ctx, node)
		if err != nil {
			return err
		}

		// 下一次任务
		data.Delay = true
		err = h.producer.EnqueueTask(data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *NodePowTaskHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {
	h.Log.Errorf("task %s failed: %v", task.Type(), err)
	data := new(model.Task)
	err = json.Unmarshal(task.Payload(), data)
	if err != nil {
		h.Log.Errorf("task error handler %s failed: %v", task.Type(), err)
		return
	}
	node := new(model.Node)
	if err := json.Unmarshal(data.Data, node); err != nil {
		h.Log.Errorf("task error handler %s failed: %v", task.Type(), err)
		return
	}
	err = h.nodeDomain.Unregister(ctx, node.Key)
	if err != nil {
		h.Log.Errorf("task error handler %s failed: %v", task.Type(), err)
		return
	}
}
