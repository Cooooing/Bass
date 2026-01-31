package biz

import (
	"common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
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
	asynqCache *util.AsynqCache
	producer   *client.Producer
}

func NewNodePowTaskHandler(baseDomain *base.BaseDomain, nodeDomain *NodeDomain, nodeRepo repo.NodeRepo, nodeCache cache.NodeCache, asynqCache *util.AsynqCache, producer *client.Producer) *NodePowTaskHandler {
	return &NodePowTaskHandler{
		BaseDomain: baseDomain,
		nodeDomain: nodeDomain,
		nodeRepo:   nodeRepo,
		nodeCache:  nodeCache,
		asynqCache: asynqCache,
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

		// 判断任务是否存在
		if version, err := h.asynqCache.GetAsynqTaskVersion(ctx, data.TaskName); err != nil {
			return err
		} else if version != data.Version {
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

		// 下一次任务
		rc, ok1 := asynq.GetRetryCount(ctx)
		mrc, ok2 := asynq.GetMaxRetry(ctx)
		if ok1 && ok2 && rc <= mrc {
			data.Delay = true
			err = h.producer.EnqueueContextTask(ctx, data)
			if err != nil {
				return err
			}
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
