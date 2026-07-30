package repo

import (
	"common/pkg/client"
	"content/internal/biz/repo"
	"context"
	"fmt"
)

var _ repo.EventClient = (*NatsEventClient)(nil)

type NatsEventClient struct {
	natsClient *client.NatsClient
}

func NewNatsEventClient(
	natsClient *client.NatsClient,
) repo.EventClient {
	return &NatsEventClient{
		natsClient: natsClient,
	}
}

func (p *NatsEventClient) Publish(ctx context.Context, msg *repo.EventClientMessage) error {
	if msg == nil {
		return fmt.Errorf("event message is nil")
	}
	if err := p.natsClient.Publish(ctx, msg.Subject, &client.Message{
		Subject: msg.Subject,
		Data:    []byte(msg.Payload),
		Header:  msg.Headers,
	}); err != nil {
		return err
	}
	return nil
}
