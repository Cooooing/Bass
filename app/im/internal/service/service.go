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
	NewEmojiService,
)

func ProvideServices(commonSystemService *CommonSystemService, privateChatService *PrivateChatService, globalChatService *GlobalChatService, tccGlobalChatTopicService *TCCGlobalChatTopicService, emojiService *EmojiService) []server.Service {
	return []server.Service{
		commonSystemService,
		privateChatService,
		globalChatService,
		tccGlobalChatTopicService,
		emojiService,
	}
}
