package server

import (
	"context"
	"game_idle/internal/biz/usecase"
	"log/slog"
)

// MetadataCacheServer 适配 Kratos 生命周期，启动游戏元数据本地缓存刷新。
type MetadataCacheServer struct {
	logger               *slog.Logger
	metadataCacheUsecase *usecase.MetadataCacheUsecase
}

func NewMetadataCacheServer(
	logger *slog.Logger,
	metadataCacheUsecase *usecase.MetadataCacheUsecase,
) *MetadataCacheServer {
	return &MetadataCacheServer{
		logger:               logger,
		metadataCacheUsecase: metadataCacheUsecase,
	}
}

func (s *MetadataCacheServer) Start(ctx context.Context) error {
	if err := s.metadataCacheUsecase.Start(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle metadata cache refresh started")
	return nil
}

func (s *MetadataCacheServer) Stop(ctx context.Context) error {
	if err := s.metadataCacheUsecase.Stop(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle metadata cache refresh stopped")
	return nil
}
