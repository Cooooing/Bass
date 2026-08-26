package server

import (
	"context"
	"game_idle/internal/biz/usecase"
	"log/slog"
)

// ActionQueueServer 适配 Kratos 生命周期，启动行动队列调度循环。
type ActionQueueServer struct {
	logger             *slog.Logger
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionQueueServer(
	logger *slog.Logger,
	actionQueueUsecase *usecase.ActionQueueUsecase,
) *ActionQueueServer {
	return &ActionQueueServer{
		logger:             logger,
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (s *ActionQueueServer) Start(ctx context.Context) error {
	if err := s.actionQueueUsecase.Start(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle action queue started")
	return nil
}

func (s *ActionQueueServer) Stop(ctx context.Context) error {
	if err := s.actionQueueUsecase.Stop(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle action queue stopped")
	return nil
}
