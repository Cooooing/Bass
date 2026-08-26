package server

import (
	"common/pkg/client/timewheel"
	"context"
	"log/slog"
)

// TimeWheelServer 是挂机游戏运行时任务时间轮服务。
type TimeWheelServer struct {
	logger *slog.Logger
	wheel  *timewheel.TimeWheel
}

// NewTimeWheelServer 创建时间轮生命周期适配服务。
func NewTimeWheelServer(
	logger *slog.Logger,
	wheel *timewheel.TimeWheel,
) *TimeWheelServer {
	return &TimeWheelServer{
		logger: logger,
		wheel:  wheel,
	}
}

// Start 启动时间轮。
func (s *TimeWheelServer) Start(ctx context.Context) error {
	if err := s.wheel.Start(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle time wheel started")
	return nil
}

// Stop 停止时间轮。
func (s *TimeWheelServer) Stop(ctx context.Context) error {
	if err := s.wheel.Stop(ctx); err != nil {
		return err
	}
	s.logger.Info("game idle time wheel stopped")
	return nil
}
