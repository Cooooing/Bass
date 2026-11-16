package repo

import (
	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	"context"
	"fmt"

	"content/internal/biz/repo"
)

var _ repo.EventClient = (*EventClient)(nil)

type EventClient struct {
	natsClient *nats.Client
}

func (p *EventClient) Ensure(ctx context.Context) error {
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamContent)
	return p.natsClient.EnsureStream(ctx, nats.StreamSpec{
		Name:              stream.Name.String(),
		Subjects:          stream.Subjects,
		Retention:         stream.Retention,
		AllowMsgSchedules: stream.AllowMsgSchedules,
		AllowMsgTTL:       stream.AllowMsgTTL,
		AllowRollup:       stream.AllowRollup,
	})
}

func NewEventClient(
	natsClient *nats.Client,
) repo.EventClient {
	return &EventClient{
		natsClient: natsClient,
	}
}

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
	if err != nil {
		return err
	}
	return nil
}
