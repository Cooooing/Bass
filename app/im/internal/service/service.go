package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	ProvideGrpcServices,
	ProvideHttpServices,

	NewChatSessionService,
	NewChatMessageService,
)

func ProvideGrpcServices(
	systemService *SystemService,
	chatSessionService *ChatSessionService,
	chatMessageService *ChatMessageService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		chatSessionService,
		chatMessageService,
	}
}

func ProvideHttpServices(
	systemService *SystemService,
	chatSessionService *ChatSessionService,
	chatMessageService *ChatMessageService,
) []server.HttpService {
	return []server.HttpService{
		systemService,
		chatSessionService,
		chatMessageService,
	}
}
