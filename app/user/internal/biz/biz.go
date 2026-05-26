package biz

import (
	"common/pkg/util/jwt"
	"user/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	usecase.NewTokenUsecase,
	usecase.NewAccountValidationUsecase,

	usecase.NewAuthUsecase,
	usecase.NewRelationUsecase,
	usecase.NewTfaUsecase,
)
