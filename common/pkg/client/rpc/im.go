package rpc

import (
	"common/pkg/client/localrpc"
	imv1 "common/proto/gen/im/v1"

	"google.golang.org/grpc"
)

type IMClient struct {
	Group   imv1.IMChatGroupServiceClient
	Message imv1.IMChatMessageServiceClient
	Session imv1.IMChatSessionServiceClient
}

func NewIMClient(
	conn grpc.ClientConnInterface,
) *IMClient {
	return &IMClient{
		Group:   imv1.NewIMChatGroupServiceClient(conn),
		Message: imv1.NewIMChatMessageServiceClient(conn),
		Session: imv1.NewIMChatSessionServiceClient(conn),
	}
}

func NewLocalIMClient[T any](services []T) *IMClient {
	conn := localrpc.NewConn()
	MountIMServices(conn, services)
	return NewIMClient(conn)
}

func MountIMServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&imv1.IMChatGroupService_ServiceDesc, service)
		conn.RegisterMatching(&imv1.IMChatMessageService_ServiceDesc, service)
		conn.RegisterMatching(&imv1.IMChatSessionService_ServiceDesc, service)
	}
}
