package repo

import (
	"context"

	realtimehubv1 "common/proto/gen/realtime_hub/v1"
)

// Hub 提供 node 需要的 Hub 控制面能力。
type Hub interface {
	Connect(ctx context.Context) (realtimehubv1.RealtimeHubService_ConnectClient, error)
	ConsumeTicket(ctx context.Context, realm, ticket string) (int64, error)
}
