package repo

import (
	"context"
	"fmt"

	"common/pkg/client"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"
	bizrepo "realtime_hub/internal/biz/repo"

	"google.golang.org/protobuf/encoding/protojson"
)

type ControlPublisher struct {
	nats client.Publisher
}

func NewControlPublisher(nats client.Publisher) bizrepo.ControlPublisher {
	return &ControlPublisher{nats: nats}
}

func (p *ControlPublisher) PublishCloseConnection(ctx context.Context, realm, nodeID, connectionID, reason string) error {
	payload, err := protojson.Marshal(&realtimehubv1.HubControl{
		Realm:  realm,
		NodeId: nodeID,
		Control: &realtimehubv1.Control{
			Payload: &realtimehubv1.Control_CloseConnection{
				CloseConnection: &realtimehubv1.CloseConnection{ConnectionId: connectionID, Reason: reason},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("encode realtime control: %w", err)
	}
	return p.nats.Publish(ctx, "realtime.control."+realm+"."+nodeID, &client.Message{Data: payload})
}
