package biz

import (
	"economy/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 鎻愪緵涓氬姟渚濊禆
var BizProviderSet = wire.NewSet(
	usecase.NewPointsUsecase,
)
