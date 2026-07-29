package server

import (
	"context"
	"scheduler/internal/biz/usecase"
)

type DelayedTaskConsumerServer struct {
	delayedTaskUsecase *usecase.DelayedTaskUsecase
}

func NewDelayedTaskConsumerServer(
	delayedTaskUsecase *usecase.DelayedTaskUsecase,
) *DelayedTaskConsumerServer {
	return &DelayedTaskConsumerServer{
		delayedTaskUsecase: delayedTaskUsecase,
	}
}
func (s *DelayedTaskConsumerServer) Start(ctx context.Context) error {
	return s.delayedTaskUsecase.StartConsuming(ctx)
}
func (s *DelayedTaskConsumerServer) Stop(ctx context.Context) error {
	return s.delayedTaskUsecase.StopConsuming(ctx)
}
