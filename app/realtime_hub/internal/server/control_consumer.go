package server

import (
	"context"
	"fmt"

	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"

	"realtime_hub/internal/biz/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

type ControlConsumer struct {
	nats         *nats.Client
	realtime     *usecase.RealtimeUsecase
	subscription *nats.Subscription
	cancel       context.CancelFunc
}

func (c *ControlConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	subscription, err := c.nats.Subscribe(runCtx, commonenum.RealtimeControlSubject, c.handle)
	if err != nil {
		cancel()
		return err
	}
	c.subscription = subscription
	return nil
}

func (c *ControlConsumer) Stop(context.Context) error {
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

func (c *ControlConsumer) handle(ctx context.Context, message *nats.Message) error {
	control := &realtimehubv1.HubControl{}
	if err := protojson.Unmarshal(message.Data, control); err != nil {
		return fmt.Errorf("decode realtime control: %w", err)
	}
	closeConnection := control.GetControl().GetCloseConnection()
	if control.GetRealm() == "" || control.GetNodeId() == "" || closeConnection == nil || closeConnection.GetConnectionId() == "" {
		return fmt.Errorf("realtime control is invalid")
	}
	_ = c.realtime.DeliverCloseConnection(ctx, control.GetRealm(), control.GetNodeId(), closeConnection.GetConnectionId(), closeConnection.GetReason())
	return nil
}
