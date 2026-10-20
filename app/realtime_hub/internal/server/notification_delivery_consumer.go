package server

import (
	"context"
	"fmt"

	"common/pkg/client"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"realtime_hub/internal/biz/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

// NotificationDeliveryConsumer 消费站内信事件并路由 notification 下行消息。
type NotificationDeliveryConsumer struct {
	nats     client.Subscriber
	realtime *usecase.RealtimeUsecase
	cancel   context.CancelFunc
}

func (c *NotificationDeliveryConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	_, err := c.nats.Subscribe(runCtx, commonenum.EventSubjectNotifyStationMessagesCreated.String(), c.handle)
	return err
}

func (c *NotificationDeliveryConsumer) Stop(context.Context) error {
	if c.cancel != nil {
		c.cancel()
	}
	return nil
}

func (c *NotificationDeliveryConsumer) handle(ctx context.Context, message *client.Message) error {
	event := &commonenums.Event{}
	if err := protojson.Unmarshal(message.Data, event); err != nil {
		return fmt.Errorf("decode notification delivery event: %w", err)
	}
	payload := event.GetNotifyStationMessagesCreated()
	if event.GetType() != commonenums.EventType_EVENT_TYPE_NOTIFY_STATION_MESSAGES_CREATED || payload == nil {
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
