package repo

import (
	"common/pkg/client"
	"context"
	"fmt"
	"user/internal/biz/repo"
)

var _ repo.EventClient = (*NatsEventClient)(nil)

// NatsEventClient 通过 NATS 发布领域事件。
type NatsEventClient struct {
	natsClient *client.NatsClient
}

// NewNatsEventClient 创建基于 NATS 的事件客户端。
func NewNatsEventClient(natsClient *client.NatsClient) repo.EventClient {
	return &NatsEventClient{
		natsClient: natsClient,
	}
}

// Publish 通过 NATS 发布事件消息。
func (p *NatsEventClient) Publish(ctx context.Context, msg *repo.EventClientMessage) error {
	if msg == nil {
		return fmt.Errorf("event message is nil")
	}
	return p.natsClient.Publish(ctx, msg.Subject, &client.Message{
		Subject: msg.Subject,
		Data:    []byte(msg.Payload),
		Header:  msg.Headers,
	})
}
