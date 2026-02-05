package handler

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	"connector/internal/biz/base"
	"connector/internal/biz/domain"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

type RegisterHandler struct {
	*base.BaseDomain
	asynqCache *util.AsynqCache
	producer   *client.Producer
	server     *domain.ServerDomain
}

func NewRegisterHandler(baseDomain *base.BaseDomain, producer *client.Producer, asynqCache *util.AsynqCache, server *domain.ServerDomain) *RegisterHandler {
	return &RegisterHandler{
		BaseDomain: baseDomain,
		asynqCache: asynqCache,
		producer:   producer,
		server:     server,
	}
}

func (h *RegisterHandler) Name() constant.TaskName {
	return constant.TaskConnectorRegister
}

func (h *RegisterHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		h.Log.Infof("task %s: %s", task.Type(), string(task.Payload()))
		data := new(model.Task)
		err := json.Unmarshal(task.Payload(), data)
		if err != nil {
			return err
		}

		// 判断任务是否存在
		if version, err := h.asynqCache.GetAsynqTaskVersion(ctx, data.TaskName); err != nil {
			return err
		} else if version != data.Version {
			return nil
		}
		err = h.asynqCache.SetAsynqTaskExpire(ctx, data.TaskName, data.Interval*2)
		if err != nil {
			return err
		}

		// 处理任务
		err = h.server.Register()
		if err != nil {
			h.Log.Errorf("failed to register node: %v", err)
		}

		// 下一次任务，不管成功失败，常驻任务
		data.Delay = true
		err = h.producer.EnqueueContextTask(ctx, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *RegisterHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {}
