package biz

import (
	"push_hub/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	usecase.NewPushEventUsecase,
	usecase.NewNodeUsecase,
)
