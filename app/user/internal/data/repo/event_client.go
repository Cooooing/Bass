package repo

import (
	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	"context"
	"fmt"

	"user/internal/biz/repo"
)

var _ repo.EventClient = (*EventClient)(nil)

// EventClient 通过 NATS 发布领域事件。
type EventClient struct {
	natsClient *nats.Client
}

func (p *EventClient) Ensure(ctx context.Context) error {
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamUser)
	return p.natsClient.EnsureStream(ctx, nats.StreamSpec{
		Name:              stream.Name.String(),
		Subjects:          stream.Subjects,
		Retention:         stream.Retention,
		AllowMsgSchedules: stream.AllowMsgSchedules,
		AllowMsgTTL:       stream.AllowMsgTTL,
		AllowRollup:       stream.AllowRollup,
	})
}

// NewEventClient 创建基于 NATS 的事件客户端。
func NewEventClient(natsClient *nats.Client) repo.EventClient {
	return &EventClient{
		natsClient: natsClient,
	}
}

// Publish 通过 JetStream 发布 outbox 事件并等待持久化确认。
func (p *EventClient) Publish(ctx context.Context, msg *repo.EventClientMessage) error {
	if msg == nil || msg.EventID == "" {
		return fmt.Errorf("event message and event id are required")
	}
	_, err := p.natsClient.JSPublish(ctx, msg.Subject, &nats.Message{
		ID:      msg.EventID,
		Subject: msg.Subject,
		Data:    []byte(msg.Payload),
		Header:  msg.Headers,
	})
	return err
}
