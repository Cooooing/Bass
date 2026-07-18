package repo

import (
	"common/pkg/client"
	"context"
	"fmt"
	"user/internal/biz/repo"
)

var _ repo.EventClient = (*NatsEventClient)(nil)

type NatsEventClient struct {
	natsClient *client.NatsClient
}

func NewNatsEventClient(natsClient *client.NatsClient) repo.EventClient {
	return &NatsEventClient{natsClient: natsClient}
}

func (p *NatsEventClient) Publish(ctx context.Context, msg *repo.EventClientMessage) error {
	if msg == nil {
		return fmt.Errorf("event message is nil")
	}
	err := p.natsClient.Publish(ctx, msg.Subject, &client.Message{
		Subject: msg.Subject,
		Data:    []byte(msg.Payload),
		Header:  msg.Headers,
	})
	if err != nil {
		return err
	}
	return nil
}
