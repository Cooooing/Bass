package server

import (
	"context"
	"scheduler/internal/biz/usecase"
)

type ScheduledTaskConsumerServer struct {
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase
}

func NewScheduledTaskConsumerServer(
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase,
) *ScheduledTaskConsumerServer {
	return &ScheduledTaskConsumerServer{
		scheduledTaskUsecase: scheduledTaskUsecase,
	}
}
func (s *ScheduledTaskConsumerServer) Start(ctx context.Context) error {
	return s.scheduledTaskUsecase.StartConsuming(ctx)
}
func (s *ScheduledTaskConsumerServer) Stop(ctx context.Context) error {
	return s.scheduledTaskUsecase.StopConsuming(ctx)
}
