package rpc

import (
	pushhubv1 "common/proto/gen/push_hub/v1"

	"google.golang.org/grpc"
)

type PushHubClient struct {
	Event pushhubv1.PushHubEventServiceClient
	Node  pushhubv1.PushHubNodeServiceClient
}

func NewPushHubClient(
	conn *grpc.ClientConn,
) *PushHubClient {
	return &PushHubClient{
		Event: pushhubv1.NewPushHubEventServiceClient(conn),
		Node:  pushhubv1.NewPushHubNodeServiceClient(conn),
	}
}
