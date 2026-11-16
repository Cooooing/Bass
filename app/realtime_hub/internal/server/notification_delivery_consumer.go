package server

import (
	"context"
	"fmt"

	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	commonevent "common/proto/gen/common/enums/event"

	"realtime_hub/internal/biz/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

// NotificationDeliveryConsumer 消费站内信事件并路由 notification 下行消息。
type NotificationDeliveryConsumer struct {
	nats         *nats.Client
	realtime     *usecase.RealtimeUsecase
	subscription *nats.Subscription
	cancel       context.CancelFunc
}

func (c *NotificationDeliveryConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	subscription, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectNotifyStationMessagesCreated.String(), c.handle)
	if err != nil {
		cancel()
		return err
	}
	c.subscription = subscription
	return nil
}

func (c *NotificationDeliveryConsumer) Stop(context.Context) error {
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

func (c *NotificationDeliveryConsumer) handle(ctx context.Context, message *nats.Message) error {
	event := &commonevent.Event{}
	if err := protojson.Unmarshal(message.Data, event); err != nil {
		return fmt.Errorf("decode notification delivery event: %w", err)
	}
	payload := event.GetNotifyStationMessagesCreated()
	if event.GetSubject() != commonevent.EventSubject_EVENT_SUBJECT_NOTIFY_STATION_MESSAGES_CREATED || payload == nil {
		return nil
	}
	if payload.GetNotificationId() <= 0 {
		return nil
	}
	for _, receiverID := range payload.GetReceiverIds() {
		if receiverID <= 0 {
			continue
		}
		if err := c.realtime.DeliverNotification(ctx, "bbs", receiverID, payload.GetNotificationId()); err != nil {
			return err
		}
	}
	return nil
}
