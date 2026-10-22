package server

import (
	"context"
	"fmt"
	"time"

	"common/pkg/client"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"realtime_hub/internal/biz/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

// PrivateChatDeliveryConsumer 消费私聊事件并路由给消息双方。
type PrivateChatDeliveryConsumer struct {
	nats     client.Subscriber
	realtime *usecase.RealtimeUsecase
	cancel   context.CancelFunc
}

func (c *PrivateChatDeliveryConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	_, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectIMPrivateChatMessageCreated.String(), c.handle)
	return err
}

func (c *PrivateChatDeliveryConsumer) Stop(context.Context) error {
	if c.cancel != nil {
		c.cancel()
	}
	return nil
}

func (c *PrivateChatDeliveryConsumer) handle(ctx context.Context, message *client.Message) error {
	event := &commonenums.Event{}
	if err := protojson.Unmarshal(message.Data, event); err != nil {
		return fmt.Errorf("decode private chat delivery event: %w", err)
	}
	payload := event.GetImPrivateChatMessageCreated()
	if event.GetType() != commonenums.EventType_EVENT_TYPE_IM_PRIVATE_CHAT_MESSAGE_CREATED || payload == nil {
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
