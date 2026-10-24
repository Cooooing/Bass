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

// GlobalChatDeliveryConsumer 消费全局聊天室事件并广播给本机 node。
type GlobalChatDeliveryConsumer struct {
	nats     client.Subscriber
	realtime *usecase.RealtimeUsecase
	cancel   context.CancelFunc
}

func (c *GlobalChatDeliveryConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	if _, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectIMGlobalChatMessageCreated.String(), c.handle); err != nil {
		return err
	}
	_, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectIMGlobalChatBarrageCreated.String(), c.handle)
	return err
}

func (c *GlobalChatDeliveryConsumer) Stop(context.Context) error {
	if c.cancel != nil {
		c.cancel()
	}
	return nil
}

func (c *GlobalChatDeliveryConsumer) handle(ctx context.Context, message *client.Message) error {
	event := &commonenums.Event{}
	if err := protojson.Unmarshal(message.Data, event); err != nil {
		return fmt.Errorf("decode global chat delivery event: %w", err)
	}
	createdAt := time.Time{}
	if event.GetTimestamp() != nil {
		createdAt = event.GetTimestamp().AsTime()
	}
	if payload := event.GetImGlobalChatMessageCreated(); event.GetType() == commonenums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_MESSAGE_CREATED && payload != nil && payload.GetMessageId() > 0 {
		if payload.GetCreatedAt() != nil {
			createdAt = payload.GetCreatedAt().AsTime()
		}
		c.realtime.DeliverGlobalChatMessage("bbs", payload.GetMessageId(), payload.GetSenderId(), int32(payload.GetMessageType()), payload.GetContent(), createdAt)
		return nil
	}
	if payload := event.GetImGlobalChatBarrageCreated(); event.GetType() == commonenums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_BARRAGE_CREATED && payload != nil && payload.GetBarrageId() > 0 {
		if payload.GetCreatedAt() != nil {
			createdAt = payload.GetCreatedAt().AsTime()
		}
		c.realtime.DeliverGlobalChatBarrage("bbs", payload.GetBarrageId(), payload.GetUserId(), payload.GetContent(), createdAt)
	}
	return nil
}
