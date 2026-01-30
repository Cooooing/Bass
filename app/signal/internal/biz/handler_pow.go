package biz

import (
	"common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"context"
	"encoding/json"
	"signal/internal/biz/base"
	"signal/internal/biz/cache"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"

	"github.com/hibiken/asynq"
)

type NodePowTaskHandler struct {
	*base.BaseDomain
	nodeDomain *NodeDomain
	nodeRepo   repo.NodeRepo
	nodeCache  cache.NodeCache
	producer   *client.Producer
}

func NewNodePowTaskHandler(baseDomain *base.BaseDomain, nodeDomain *NodeDomain, nodeRepo repo.NodeRepo, nodeCache cache.NodeCache, producer *client.Producer) *NodePowTaskHandler {
	return &NodePowTaskHandler{
		BaseDomain: baseDomain,
		nodeDomain: nodeDomain,
		nodeRepo:   nodeRepo,
		nodeCache:  nodeCache,
		producer:   producer,
	}
}

func (h *NodePowTaskHandler) Name() constant.TaskName {
	return constant.TaskSignalNodePow
}

func (h *NodePowTaskHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		h.Log.Infof("task %s: %s", task.Type(), string(task.Payload()))
		data := new(commonModel.Task)
		err := json.Unmarshal(task.Payload(), data)
		if err != nil {
			return err
		}
		node := new(model.Node)
		if err := json.Unmarshal(data.Data, node); err != nil {
			return err
		}

		// 判断节点是否在线
		ok, err := h.nodeCache.ExistsNodeRank(ctx, node.Key)
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
		err = h.nodeCache.UpdateNodePowCost(ctx, node.Key, powMs)
		if err != nil {
			return err
		}
		// 更新节点分数
		err = h.nodeCache.UpdateScore(ctx, node.Key)
		if err != nil {
			return err
		}

		return nil
	}
}

func (h *NodePowTaskHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {
	h.Log.Errorf("task %s failed: %v", task.Type(), err)
	data := new(commonModel.Task)
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
