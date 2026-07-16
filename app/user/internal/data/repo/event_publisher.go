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

func (p *NatsEventClient) Publish(ctx context.Context, req *repo.EventClientPublishReq) (*repo.EventClientPublishResponse, error) {
	if req == nil || req.Message == nil {
		return nil, fmt.Errorf("event message is nil")
	}
	msg := req.Message
	err := p.natsClient.Publish(ctx, msg.Subject, &client.Message{
		Subject: msg.Subject,
		Data:    []byte(msg.Payload),
		Header:  msg.Headers,
	})
	if err != nil {
		return nil, err
	}
	return &repo.EventClientPublishResponse{}, nil
}
