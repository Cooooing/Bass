package consumer

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	"common/pkg/util"
	"context"
	"notify/internal/biz/usecase"

	"log/slog"
)

// Consumer 是 NATS 事件消费者。
type Consumer struct {
	log          *util.LogHelper
	natsClient   *client.NatsClient
	eventUsecase *usecase.EventUsecase
	subjects     usecase.EventSubjects
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewConsumer(
	logger *slog.Logger,
	natsClient *client.NatsClient,
	eventUsecase *usecase.EventUsecase,
	subjects usecase.EventSubjects,
) *Consumer {
	return &Consumer{
		log:          util.NewLogHelper(logger),
		natsClient:   natsClient,
		eventUsecase: eventUsecase,
		subjects:     subjects,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)
	queueGroup := string(commonenum.EventQueueGroupNotify)
	for _, subject := range c.subjects {
		subjectName := string(subject)
		_, err := c.natsClient.QueueSubscribe(c.ctx, subjectName, queueGroup, func(ctx context.Context, msg *client.Message) error {
			if msg == nil {
				return nil
			}
			return c.eventUsecase.HandleMessage(ctx, msg.Subject, msg.Data)
		})
		if err != nil {
			c.log.Errorf("queue subscribe %s[%s] failed: %v", subjectName, queueGroup, err)
			continue
		}
	}
	return nil
}

func (c *Consumer) Stop(_ context.Context) error {
	if c.cancel == nil {
		return nil
	}
	c.cancel()
	return nil
}
