package biz

import (
	"game_town/internal/biz/agent"
	"game_town/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	agent.NewBladesRunner,
	usecase.NewGameUsecase,
)
