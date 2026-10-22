package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewPrivateChatService,
	NewGlobalChatService,
)

func ProvideServices(commonSystemService *CommonSystemService, privateChatService *PrivateChatService, globalChatService *GlobalChatService) []server.Service {
	return []server.Service{
		commonSystemService,
		privateChatService,
		globalChatService,
	}
}
