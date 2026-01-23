package biz

import (
	"common/pkg/cutil/base"
	"encoding/json"
	"signal/internal/biz/model"
	"signal/internal/biz/task"

	"github.com/hibiken/asynq"
)

type Producer struct {
	client *asynq.Client
}

func NewProducer(client *task.AsynqClient) *Producer {
	return &Producer{
		client: client.Client,
	}
}

func (p *Producer) EnqueueTask(data *model.Task) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	opts := make([]asynq.Option, 0)
	if data.TaskId != "" {
		opts = append(opts, asynq.TaskID(data.TaskId))
	}
	_, err = p.client.Enqueue(
		asynq.NewTask(
			data.TaskName,
			marshal,
			opts...,
		),
		asynq.MaxRetry(data.MaxRetry),
		asynq.ProcessIn(base.If(data.Delay, data.Interval, 0)),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) EnqueueTasks(data []*model.Task) error {
	for _, d := range data {
		err := p.EnqueueTask(d)
		if err != nil {
			return err
		}
	}
	return nil
}
