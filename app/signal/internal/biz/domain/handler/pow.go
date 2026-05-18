package handler

import (
	"common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util/task"
	"context"
	"encoding/json"
	"errors"
	"signal/internal/biz/cache"
	"signal/internal/biz/domain"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type PowHandler struct {
	log        *log.Helper
	nodeDomain *domain.NodeDomain
	nodeRepo   repo.NodeRepo
	nodeCache  cache.NodeCache
	asynqCache *task.AsynqCache
	producer   *client.Producer
}

func NewPowHandler(
	logger log.Logger,
	nodeDomain *domain.NodeDomain,
	nodeRepo repo.NodeRepo,
	nodeCache cache.NodeCache,
	asynqCache *task.AsynqCache,
	producer *client.Producer,
) *PowHandler {
	return &PowHandler{
		log:        log.NewHelper(logger),
		nodeDomain: nodeDomain,
		nodeRepo:   nodeRepo,
		nodeCache:  nodeCache,
		asynqCache: asynqCache,
		producer:   producer,
	}
}

func (h *PowHandler) Name() constant.TaskName {
	return constant.TaskSignalNodePow
}

func (h *PowHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		h.log.Debugf("task %s: %s", task.Type(), string(task.Payload()))
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
		data.Delay = true
		err = h.producer.EnqueueContextTask(ctx, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *PowHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {
	go func() {
		data := new(commonModel.Task)
		err = json.Unmarshal(task.Payload(), data)
		if err != nil {
			h.log.Errorf("task error handler %s failed: %v", task.Type(), err)
			return
		}
		node := new(model.Node)
		if err = json.Unmarshal(data.Data, node); err != nil {
			h.log.Errorf("task error handler %s failed: %v", task.Type(), err)
			return
		}
		if err := h.nodeDomain.Unregister(ctx, node.Key); err != nil {
			h.log.Errorf("task error handler %s failed: %v", task.Type(), err)
		}
	}()
}
