package handler

import (
	"common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"encoding/json"
	"errors"
	"signal/internal/biz/base"
	"signal/internal/biz/cache"
	"signal/internal/biz/domain"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type SessionHandler struct {
	*base.BaseDomain
	nodeDomain   *domain.NodeDomain
	nodeRepo     repo.NodeRepo
	nodeCache    cache.NodeCache
	sessionCache cache.SessionCache
	asynqCache   *util.AsynqCache
	producer     *client.Producer
}

func NewSessionHandler(baseDomain *base.BaseDomain, nodeDomain *domain.NodeDomain, nodeRepo repo.NodeRepo, nodeCache cache.NodeCache, sessionCache cache.SessionCache, asynqCache *util.AsynqCache, producer *client.Producer) *SessionHandler {
	return &SessionHandler{
		BaseDomain:   baseDomain,
		nodeDomain:   nodeDomain,
		nodeRepo:     nodeRepo,
		nodeCache:    nodeCache,
		sessionCache: sessionCache,
		asynqCache:   asynqCache,
		producer:     producer,
	}
}

func (h *SessionHandler) Name() constant.TaskName {
	return constant.TaskSignalNodeSession
}

func (h *SessionHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		h.Log.Debugf("task %s: %s", task.Type(), string(task.Payload()))
		data := new(commonModel.Task)
		err := json.Unmarshal(task.Payload(), data)
		if err != nil {
			return err
		}
		node := new(model.Node)
		if err := json.Unmarshal(data.Data, node); err != nil {
			return err
		}

		// 幂等
		exists, _ := h.nodeCache.ExistsNodeRank(ctx, node.Key)
		if !exists {
			return nil
		}
		if version, err := h.asynqCache.GetAsynqTaskVersion(ctx, data.TaskName); err != nil {
			if errors.Is(err, redis.Nil) {
				return nil
			}
			return err
		} else if version != data.Version {
			return nil
		}
		err = h.asynqCache.SetAsynqTaskExpire(ctx, data.TaskName, data.Interval*2)
		if err != nil {
			return err
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
		data.Delay = true
		err = h.producer.EnqueueContextTask(ctx, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *SessionHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {
	go func() {
		data := new(commonModel.Task)
		err = json.Unmarshal(task.Payload(), data)
		if err != nil {
			h.Log.Errorf("task error handler %s failed: %v", task.Type(), err)
			return
		}
		node := new(model.Node)
		if err = json.Unmarshal(data.Data, node); err != nil {
			h.Log.Errorf("task error handler %s failed: %v", task.Type(), err)
			return
		}
		if err := h.nodeDomain.Unregister(ctx, node.Key); err != nil {
			h.Log.Errorf("task error handler %s failed: %v", task.Type(), err)
		}
	}()
}
