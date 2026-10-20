package server

import (
	"context"
	"fmt"

	"common/pkg/client"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"
	"realtime_hub/internal/biz/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

type ControlConsumer struct {
	nats     client.Subscriber
	realtime *usecase.RealtimeUsecase
	cancel   context.CancelFunc
}

func NewControlConsumer(nats client.Subscriber, realtime *usecase.RealtimeUsecase) *ControlConsumer {
	return &ControlConsumer{nats: nats, realtime: realtime}
}

func (c *ControlConsumer) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	_, err := c.nats.Subscribe(runCtx, "realtime.control.*.*", c.handle)
	return err
}

func (c *ControlConsumer) Stop(context.Context) error {
	if c.cancel != nil {
		c.cancel()
	}
	return nil
}

func (c *ControlConsumer) handle(ctx context.Context, message *client.Message) error {
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
