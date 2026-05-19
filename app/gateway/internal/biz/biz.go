package biz

import (
	"common/pkg/util/jwt"
	"gateway/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	usecase.NewIpUsecase,
)
