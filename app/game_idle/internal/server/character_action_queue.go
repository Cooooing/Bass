package server

import (
	"context"
	"game_idle/internal/biz/usecase"
	"log/slog"
)

// CharacterActionQueueServer 适配 Kratos 生命周期，启动行动队列调度循环。
type CharacterActionQueueServer struct {
	logger                 *slog.Logger
	actionSchedulerUsecase *usecase.CharacterActionSchedulerUsecase
}

func NewCharacterActionQueueServer(
	logger *slog.Logger,
	actionSchedulerUsecase *usecase.CharacterActionSchedulerUsecase,
) *CharacterActionQueueServer {
	return &CharacterActionQueueServer{
		logger:                 logger,
		actionSchedulerUsecase: actionSchedulerUsecase,
	}
}

func (s *CharacterActionQueueServer) Start(ctx context.Context) error {
	if err := s.actionSchedulerUsecase.Start(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle action queue started")
	return nil
}

func (s *CharacterActionQueueServer) Stop(ctx context.Context) error {
	if err := s.actionSchedulerUsecase.Stop(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle action queue stopped")
	return nil
}
