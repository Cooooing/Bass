package biz

import (
	"common/pkg/util"
	"common/pkg/util/jwt"
	"user/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	usecase.NewTokenUsecase,
	util.NewEventPool,

	usecase.NewAuthUsecase,
	usecase.NewAccountUsecase,
	usecase.NewRelationUsecase,
	usecase.NewTfaUsecase,
)
