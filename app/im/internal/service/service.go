package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	ProvideGrpcServices,
	ProvideHttpServices,

	NewChatGroupService,
	NewChatSessionService,
	NewChatMessageService,
)

func ProvideGrpcServices(
	systemService *SystemService,
	chatGroupService *ChatGroupService,
	chatSessionService *ChatSessionService,
	chatMessageService *ChatMessageService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		chatGroupService,
		chatSessionService,
		chatMessageService,
	}
}

func ProvideHttpServices(
	systemService *SystemService,
	chatGroupService *ChatGroupService,
	chatSessionService *ChatSessionService,
	chatMessageService *ChatMessageService,
) []server.HttpService {
	return []server.HttpService{
		systemService,
		chatGroupService,
		chatSessionService,
		chatMessageService,
	}
}
