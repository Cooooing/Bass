package biz

import (
	"badge/internal/biz/usecase"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	usecase.NewBadgeUsecase,
	usecase.NewBadgeRenderUsecase,
)
