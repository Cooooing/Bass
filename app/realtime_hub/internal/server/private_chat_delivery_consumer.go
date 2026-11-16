package server

import (
	"context"
	"fmt"
	"time"

	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	commonevent "common/proto/gen/common/enums/event"

	"realtime_hub/internal/biz/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

// PrivateChatDeliveryConsumer 消费私聊事件并路由给消息双方。
type PrivateChatDeliveryConsumer struct {
	nats         *nats.Client
	realtime     *usecase.RealtimeUsecase
	subscription *nats.Subscription
	cancel       context.CancelFunc
}

func (c *PrivateChatDeliveryConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	subscription, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectIMPrivateChatMessageCreated.String(), c.handle)
	if err != nil {
		cancel()
		return err
	}
	c.subscription = subscription
	return nil
}

func (c *PrivateChatDeliveryConsumer) Stop(context.Context) error {
	if c.cancel != nil {
		c.cancel()
	}
	if c.subscription != nil {
		err := c.subscription.Unsubscribe()
		c.subscription = nil
		return err
	}
	return nil
}

func (c *PrivateChatDeliveryConsumer) handle(ctx context.Context, message *nats.Message) error {
	event := &commonevent.Event{}
	if err := protojson.Unmarshal(message.Data, event); err != nil {
		return fmt.Errorf("decode private chat delivery event: %w", err)
	}
	payload := event.GetImPrivateChatMessageCreated()
	if event.GetSubject() != commonevent.EventSubject_EVENT_SUBJECT_IM_PRIVATE_CHAT_MESSAGE_CREATED || payload == nil {
		return nil
	}
	if payload.GetMessageId() <= 0 || payload.GetSenderId() <= 0 || payload.GetReceiverId() <= 0 {
		return nil
	}
	createdAt := time.Time{}
	if payload.GetCreatedAt() != nil {
		createdAt = payload.GetCreatedAt().AsTime()
	} else if event.GetTimestamp() != nil {
		createdAt = event.GetTimestamp().AsTime()
	}
	return c.realtime.DeliverPrivateChatMessage(ctx, "bbs", payload.GetSenderId(), payload.GetReceiverId(), payload.GetMessageId(), int32(payload.GetMessageType()), payload.GetContent(), createdAt)
}
