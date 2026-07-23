package biz

import (
	"game_town/internal/biz/usecase"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	usecase.NewEventUsecase,
	usecase.NewAgentConfigUsecase,
	usecase.NewPlayerUsecase,
	usecase.NewWorldUsecase,
	usecase.NewWorldMemberUsecase,
	usecase.NewNpcUsecase,
	usecase.NewLocationUsecase,
	usecase.NewFactionUsecase,
	usecase.NewWorldAgentRunner,
)
