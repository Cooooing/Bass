package biz

import (
	"common/pkg/util"
	"im/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	util.NewEventPool,

	usecase.NewChatGroupUsecase,
	usecase.NewChatSessionUsecase,
	usecase.NewChatMessageUsecase,
)
