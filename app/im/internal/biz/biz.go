package biz

import (
	"common/pkg/util"
	"common/pkg/util/jwt"
	"im/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	util.NewEventPool,

	usecase.NewChatGroupUsecase,
	usecase.NewChatSessionUsecase,
	usecase.NewChatMessageUsecase,
)
