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
	NewTCCGlobalChatTopicService,
	NewGlobalChatBarrageService,
	NewTCCGlobalChatBarrageService,
	NewEmojiService,
)

func ProvideServices(commonSystemService *CommonSystemService, privateChatService *PrivateChatService, globalChatService *GlobalChatService, tccGlobalChatTopicService *TCCGlobalChatTopicService, globalChatBarrageService *GlobalChatBarrageService, tccGlobalChatBarrageService *TCCGlobalChatBarrageService, emojiService *EmojiService) []server.Service {
	return []server.Service{
		commonSystemService,
		privateChatService,
		globalChatService,
		tccGlobalChatTopicService,
		globalChatBarrageService,
		tccGlobalChatBarrageService,
		emojiService,
	}
}
