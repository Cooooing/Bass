package biz

import (
	"game_idle_bff/internal/biz/usecase"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	usecase.NewAuthUsecase,
	usecase.NewCharacterUsecase,
	usecase.NewBackpackUsecase,
	usecase.NewActionQueueUsecase,
)
