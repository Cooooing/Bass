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

// GlobalChatDeliveryConsumer 消费全局聊天室事件并广播给本机 node。
type GlobalChatDeliveryConsumer struct {
	nats          *nats.Client
	realtime      *usecase.RealtimeUsecase
	subscriptions []*nats.Subscription
	cancel        context.CancelFunc
}

func (c *GlobalChatDeliveryConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	messageSubscription, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectIMGlobalChatMessageCreated.String(), c.handle)
	if err != nil {
		cancel()
		return err
	}
	barrageSubscription, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectIMGlobalChatBarrageCreated.String(), c.handle)
	if err != nil {
		cancel()
		_ = messageSubscription.Unsubscribe()
		return err
	}
	topicSubscription, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectIMGlobalChatTopicUpdated.String(), c.handle)
	if err != nil {
		cancel()
		_ = barrageSubscription.Unsubscribe()
		_ = messageSubscription.Unsubscribe()
		return err
	}
	c.subscriptions = []*nats.Subscription{messageSubscription, barrageSubscription, topicSubscription}
	return nil
}

func (c *GlobalChatDeliveryConsumer) Stop(context.Context) error {
	if c.cancel != nil {
		c.cancel()
	}
	var stopErr error
	for _, subscription := range c.subscriptions {
		if err := subscription.Unsubscribe(); err != nil && stopErr == nil {
			stopErr = err
		}
	}
	c.subscriptions = nil
	return stopErr
}

func (c *GlobalChatDeliveryConsumer) handle(ctx context.Context, message *nats.Message) error {
	event := &commonevent.Event{}
	if err := protojson.Unmarshal(message.Data, event); err != nil {
		return fmt.Errorf("decode global chat delivery event: %w", err)
	}
	createdAt := time.Time{}
	if event.GetTimestamp() != nil {
		createdAt = event.GetTimestamp().AsTime()
	}
	if payload := event.GetImGlobalChatMessageCreated(); event.GetSubject() == commonevent.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_MESSAGE_CREATED && payload != nil && payload.GetMessageId() > 0 {
		if payload.GetCreatedAt() != nil {
			createdAt = payload.GetCreatedAt().AsTime()
		}
		c.realtime.DeliverGlobalChatMessage("bbs", payload.GetMessageId(), payload.GetSenderId(), int32(payload.GetMessageType()), payload.GetContent(), createdAt)
		return nil
	}
	if payload := event.GetImGlobalChatBarrageCreated(); event.GetSubject() == commonevent.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_BARRAGE_CREATED && payload != nil && payload.GetBarrageId() > 0 {
		if payload.GetCreatedAt() != nil {
			createdAt = payload.GetCreatedAt().AsTime()
		}
		c.realtime.DeliverGlobalChatBarrage("bbs", payload.GetBarrageId(), payload.GetUserId(), payload.GetContent(), createdAt)
		return nil
	}
	if payload := event.GetImGlobalChatTopicUpdated(); event.GetSubject() == commonevent.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_TOPIC_UPDATED && payload != nil && payload.GetTopicId() > 0 {
		if payload.GetCreatedAt() != nil {
			createdAt = payload.GetCreatedAt().AsTime()
		}
		c.realtime.DeliverGlobalChatTopic("bbs", payload.GetTopicId(), payload.GetUserId(), payload.GetContent(), createdAt)
	}
	return nil
}
