package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewChatGroupService,
	NewChatSessionService,
	NewChatMessageService,
)

func ProvideServices(commonSystemService *CommonSystemService, chatGroupService *ChatGroupService, chatSessionService *ChatSessionService, chatMessageService *ChatMessageService) []server.Service {
	return []server.Service{
		commonSystemService,
		chatGroupService,
		chatSessionService,
		chatMessageService,
	}
}
