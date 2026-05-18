package consumer

import (
	"common/api/gen/common/enums"
	"common/pkg/client"
	"context"
	"notify/internal/biz/domain/handler"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/proto"
)

// Consumer NATS 事件消费者
type Consumer struct {
	log        *log.Helper
	natsClient *client.NatsClient
	dispatcher *handler.Dispatcher
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewConsumer(
	logger log.Logger,
	natsClient *client.NatsClient,
	dispatcher *handler.Dispatcher,
) *Consumer {
	return &Consumer{
		log:        log.NewHelper(logger),
		natsClient: natsClient,
		dispatcher: dispatcher,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)
	for _, subject := range []string{"content.>", "user.>"} {
		_, err := c.natsClient.Subscribe(c.ctx, subject, c.handleMessage)
		if err != nil {
			c.log.Errorf("subscribe %s failed: %v", subject, err)
			continue
		}
		c.log.Infof("subscribed to %s", subject)
	}
	return nil
}

func (c *Consumer) Stop(_ context.Context) error {
	c.cancel()
	return nil
}

func (c *Consumer) handleMessage(ctx context.Context, msg *client.Message) error {
	var event enums.Event
	if err := proto.Unmarshal(msg.Data, &event); err != nil {
		c.log.Errorf("unmarshal event failed: subject=%s err=%v", msg.Subject, err)
		return nil
	}

	c.log.Infof("received event: type=%v receivers=%d subject=%s",
		event.Type, len(event.ReceiverIds), msg.Subject)

	return c.dispatcher.Dispatch(ctx, &event)
}
