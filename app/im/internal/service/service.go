package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	ProvideServices,

	NewChatSessionService,
	NewChatMessageService,
)

func ProvideServices(
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
