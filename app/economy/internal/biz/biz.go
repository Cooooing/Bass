package biz

import (
	"economy/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 提供业务依赖
var BizProviderSet = wire.NewSet(
	usecase.NewEconomyUsecase,
)
