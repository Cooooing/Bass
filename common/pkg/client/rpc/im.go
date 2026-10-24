package rpc

import (
	"common/pkg/client/localrpc"
	imv1 "common/proto/gen/im/v1"

	"google.golang.org/grpc"
)

type IMClient struct {
	PrivateChat       imv1.IMPrivateChatServiceClient
	GlobalChat        imv1.IMGlobalChatServiceClient
	GlobalChatBarrage imv1.IMGlobalChatBarrageServiceClient
	Emoji             imv1.IMEmojiServiceClient
}

func NewIMClient(
	conn grpc.ClientConnInterface,
) *IMClient {
	return &IMClient{
		PrivateChat:       imv1.NewIMPrivateChatServiceClient(conn),
		GlobalChat:        imv1.NewIMGlobalChatServiceClient(conn),
		GlobalChatBarrage: imv1.NewIMGlobalChatBarrageServiceClient(conn),
		Emoji:             imv1.NewIMEmojiServiceClient(conn),
	}
}

func MountIMServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&imv1.IMPrivateChatService_ServiceDesc, service)
		conn.RegisterMatching(&imv1.IMGlobalChatService_ServiceDesc, service)
		conn.RegisterMatching(&imv1.IMGlobalChatBarrageService_ServiceDesc, service)
		conn.RegisterMatching(&imv1.IMEmojiService_ServiceDesc, service)
	}
}
