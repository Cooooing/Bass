package repo

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"context"
	"encoding/json"
	bizrepo "scheduler/internal/biz/repo"
)

var _ bizrepo.TaskEventBus = (*TaskEventBus)(nil)

type TaskEventBus struct {
	natsClient *commonClient.NatsClient
}

func NewTaskEventBus(
	natsClient *commonClient.NatsClient,
) bizrepo.TaskEventBus {
	return &TaskEventBus{
		natsClient: natsClient,
	}
}

func (b *TaskEventBus) PublishTaskChanged(
	ctx context.Context,
	message *bizrepo.TaskChangedMessage,
) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if err := b.natsClient.Publish(ctx, constant.SchedulerTaskChangedSubject, &commonClient.Message{
		Data: data,
	}); err != nil {
		return err
	}
	return nil
}

func (b *TaskEventBus) PublishExecutionCanceled(
	ctx context.Context,
	message *bizrepo.TaskExecutionCanceledMessage,
) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if err := b.natsClient.Publish(ctx, constant.SchedulerTaskExecutionCanceledSubject, &commonClient.Message{
		Data: data,
	}); err != nil {
		return err
	}
	return nil
}

func (b *TaskEventBus) SubscribeTaskChanged(
	ctx context.Context,
) (<-chan bizrepo.TaskChangedMessage, error) {
	ch := make(chan bizrepo.TaskChangedMessage, 16)
	_, err := b.natsClient.Subscribe(ctx, constant.SchedulerTaskChangedSubject, func(_ context.Context, msg *commonClient.Message) error {
		var payload bizrepo.TaskChangedMessage
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			return err
		}
		ch <- payload
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (b *TaskEventBus) SubscribeExecutionCanceled(
	ctx context.Context,
) (<-chan bizrepo.TaskExecutionCanceledMessage, error) {
	ch := make(chan bizrepo.TaskExecutionCanceledMessage, 16)
	_, err := b.natsClient.Subscribe(ctx, constant.SchedulerTaskExecutionCanceledSubject, func(_ context.Context, msg *commonClient.Message) error {
		var payload bizrepo.TaskExecutionCanceledMessage
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			return err
		}
		ch <- payload
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ch, nil
}
