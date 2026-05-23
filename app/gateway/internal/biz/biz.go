package biz

import (
	"common/pkg/util/jwt"
	"gateway/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	usecase.NewIpUsecase,
)
