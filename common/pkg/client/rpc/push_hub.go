package rpc

import (
	"common/pkg/client/localrpc"
	pushhubv1 "common/proto/gen/push_hub/v1"

	"google.golang.org/grpc"
)

type PushHubClient struct {
	Event pushhubv1.PushHubEventServiceClient
	Node  pushhubv1.PushHubNodeServiceClient
}

func NewPushHubClient(
	conn grpc.ClientConnInterface,
) *PushHubClient {
	return &PushHubClient{
		Event: pushhubv1.NewPushHubEventServiceClient(conn),
		Node:  pushhubv1.NewPushHubNodeServiceClient(conn),
	}
}

func NewLocalPushHubClient[T any](services []T) *PushHubClient {
	conn := localrpc.NewConn()
	MountPushHubServices(conn, services)
	return NewPushHubClient(conn)
}

func MountPushHubServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&pushhubv1.PushHubEventService_ServiceDesc, service)
		conn.RegisterMatching(&pushhubv1.PushHubNodeService_ServiceDesc, service)
	}
}
