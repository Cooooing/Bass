package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,
	ProvideGrpcServices,
	ProvideHttpServices,

	NewChatGroupService,
	NewChatSessionService,
	NewChatMessageService,
)

func ProvideGrpcServices(
	commonSystemService *CommonSystemService,
	chatGroupService *ChatGroupService,
	chatSessionService *ChatSessionService,
	chatMessageService *ChatMessageService,
) []server.GrpcService {
	return []server.GrpcService{
		commonSystemService,
		chatGroupService,
		chatSessionService,
		chatMessageService,
	}
}

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
	chatGroupService *ChatGroupService,
	chatSessionService *ChatSessionService,
	chatMessageService *ChatMessageService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
		chatGroupService,
		chatSessionService,
		chatMessageService,
	}
}
