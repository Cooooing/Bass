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

type NodeSessionTaskHandler struct {
	*base.BaseDomain
	nodeDomain   *NodeDomain
	nodeRepo     repo.NodeRepo
	nodeCache    cache.NodeCache
	sessionCache cache.SessionCache
	producer     *client.Producer
}

func NewNodeSessionTaskHandler(baseDomain *base.BaseDomain, nodeDomain *NodeDomain, nodeRepo repo.NodeRepo, nodeCache cache.NodeCache, sessionCache cache.SessionCache, producer *client.Producer) *NodeSessionTaskHandler {
	return &NodeSessionTaskHandler{
		BaseDomain:   baseDomain,
		nodeDomain:   nodeDomain,
		nodeRepo:     nodeRepo,
		nodeCache:    nodeCache,
		sessionCache: sessionCache,
		producer:     producer,
	}
}

func (h *NodeSessionTaskHandler) Name() constant.TaskName {
	return constant.TaskSignalNodeSession
}

func (h *NodeSessionTaskHandler) Handler() asynq.HandlerFunc {
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
		if version, err := h.nodeCache.GetAsynqTaskVersion(ctx, data.TaskName); err != nil {
			return err
		} else if version != data.Version {
			return nil
		}
		// 获取在线会话 id
		sessionIds, err := h.nodeDomain.Session(node)
		if err != nil {
			return err
		}
		// 更新缓存
		err = h.nodeCache.UpdateNodeConnections(ctx, node.Key, int64(len(sessionIds)))
		if err != nil {
			return err
		}
		err = h.sessionCache.SetNodeSessionIds(ctx, node.Key, sessionIds)
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

func (h *NodeSessionTaskHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {
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
