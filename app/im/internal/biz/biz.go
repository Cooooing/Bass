package biz

import (
	"im/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	usecase.NewPrivateChatUsecase,
	usecase.NewGlobalChatUsecase,
	usecase.NewEmojiUsecase,
)
