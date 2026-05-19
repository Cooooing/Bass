package biz

import (
	"common/pkg/util"
	"common/pkg/util/jwt"
	"user/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	usecase.NewTokenUsecase,
	util.NewEventPool,

	usecase.NewAuthUsecase,
	usecase.NewUserUsecase,
	usecase.NewUserRelationUsecase,
	usecase.NewTwoFactorAuthUsecase,
)
