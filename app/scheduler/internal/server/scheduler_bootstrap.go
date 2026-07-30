package server

import (
	"context"
	"scheduler/internal/biz/usecase"
)

type SchedulerBootstrapServer struct {
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase
	delayedTaskUsecase   *usecase.DelayedTaskUsecase
}

func NewSchedulerBootstrapServer(
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase,
	delayedTaskUsecase *usecase.DelayedTaskUsecase,
) *SchedulerBootstrapServer {
	return &SchedulerBootstrapServer{
		scheduledTaskUsecase: scheduledTaskUsecase,
		delayedTaskUsecase:   delayedTaskUsecase,
	}
}

func (s *SchedulerBootstrapServer) Start(ctx context.Context) error {
	if err := s.scheduledTaskUsecase.SeedDefaultTasks(ctx); err != nil {
		return err
	}
	if err := s.delayedTaskUsecase.SeedDefaultTasks(ctx); err != nil {
		return err
	}
	if err := s.scheduledTaskUsecase.EnsureSchedule(ctx); err != nil {
		return err
	}
	return s.delayedTaskUsecase.EnsureSchedule(ctx)
}
func (s *SchedulerBootstrapServer) Stop(ctx context.Context) error {
	return nil
}
