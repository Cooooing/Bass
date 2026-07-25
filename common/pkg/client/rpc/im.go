package rpc

import (
	imv1 "common/proto/gen/im/v1"

	"google.golang.org/grpc"
)

type IMClient struct {
	Group   imv1.IMChatGroupServiceClient
	Message imv1.IMChatMessageServiceClient
	Session imv1.IMChatSessionServiceClient
}

func NewIMClient(
	conn *grpc.ClientConn,
) *IMClient {
	return &IMClient{
		Group:   imv1.NewIMChatGroupServiceClient(conn),
		Message: imv1.NewIMChatMessageServiceClient(conn),
		Session: imv1.NewIMChatSessionServiceClient(conn),
	}
}
