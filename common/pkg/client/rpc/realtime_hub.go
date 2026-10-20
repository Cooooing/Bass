package rpc

import (
	"common/pkg/client/localrpc"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"

	"google.golang.org/grpc"
)

type RealtimeHubClient struct {
	Hub realtimehubv1.RealtimeHubServiceClient
}

func NewRealtimeHubClient(conn grpc.ClientConnInterface) *RealtimeHubClient {
	return &RealtimeHubClient{Hub: realtimehubv1.NewRealtimeHubServiceClient(conn)}
}

func MountRealtimeHubServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&realtimehubv1.RealtimeHubService_ServiceDesc, service)
	}
}
