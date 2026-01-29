package biz

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"connector/internal/biz/base"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

type NodeRegisterTaskHandler struct {
	*base.BaseDomain
	producer *client.Producer
}

func NewNodeRegisterTaskHandler(baseDomain *base.BaseDomain, producer *client.Producer) *NodeRegisterTaskHandler {
	return &NodeRegisterTaskHandler{
		BaseDomain: baseDomain,
		producer:   producer,
	}
}

func (h *NodeRegisterTaskHandler) Name() constant.TaskName {
	return constant.TaskConnectorRegister
}

func (h *NodeRegisterTaskHandler) Handler() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		h.Log.Infof("task %s: %s", task.Type(), string(task.Payload()))
		data := new(model.Task)
		err := json.Unmarshal(task.Payload(), data)
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

func (h *NodeRegisterTaskHandler) ErrHandler(ctx context.Context, task *asynq.Task, err error) {
	h.Log.Errorf("task %s failed: %v", task.Type(), err)
}
