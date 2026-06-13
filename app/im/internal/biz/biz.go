package biz

import (
	"common/pkg/util"
	"common/pkg/auth"
	"im/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	auth.NewTokenCache,
	util.NewEventPool,

	usecase.NewChatGroupUsecase,
	usecase.NewChatSessionUsecase,
	usecase.NewChatMessageUsecase,
)
