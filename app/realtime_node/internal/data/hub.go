package data

import (
	"context"

	"common/pkg/client/rpc"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"
	"realtime_node/internal/biz/repo"
)

type HubClient struct {
	client realtimehubv1.RealtimeHubServiceClient
}

func NewHubClient(client *rpc.RealtimeHubClient) repo.Hub {
	return &HubClient{client: client.Hub}
}

func (c *HubClient) Connect(ctx context.Context) (realtimehubv1.RealtimeHubService_ConnectClient, error) {
	return c.client.Connect(ctx)
}

func (c *HubClient) ConsumeTicket(ctx context.Context, realm, ticket string) (int64, error) {
	response, err := c.client.ConsumeTicket(ctx, &realtimehubv1.ConsumeRealtimeTicket_Req{Realm: realm, Ticket: ticket})
	if err != nil {
		return 0, err
	}
	return response.GetUserId(), nil
}
