package biz

import (
	"realtime_hub/internal/biz/usecase"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(usecase.NewRealtimeUsecase)
