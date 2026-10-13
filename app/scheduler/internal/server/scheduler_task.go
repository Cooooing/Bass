package server

import (
	"context"

	"scheduler/internal/biz/usecase"
)

type SchedulerTaskServer struct {
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase
	delayedTaskUsecase   *usecase.DelayedTaskUsecase
}

func NewSchedulerTaskServer(
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase,
	delayedTaskUsecase *usecase.DelayedTaskUsecase,
) *SchedulerTaskServer {
	return &SchedulerTaskServer{
		scheduledTaskUsecase: scheduledTaskUsecase,
		delayedTaskUsecase:   delayedTaskUsecase,
	}
}

func (s *SchedulerTaskServer) Start(ctx context.Context) error {
	if err := s.scheduledTaskUsecase.SeedDefaultTasks(ctx); err != nil {
		return err
	}
	if err := s.delayedTaskUsecase.SeedDefaultTasks(ctx); err != nil {
		return err
	}
	if err := s.scheduledTaskUsecase.EnsureSchedule(ctx); err != nil {
		return err
	}
	if err := s.delayedTaskUsecase.EnsureSchedule(ctx); err != nil {
		return err
	}
	if err := s.scheduledTaskUsecase.StartConsuming(ctx); err != nil {
		return err
	}
	if err := s.delayedTaskUsecase.StartConsuming(ctx); err != nil {
		_ = s.scheduledTaskUsecase.StopConsuming(ctx)
		return err
	}
	return nil
}

func (s *SchedulerTaskServer) Stop(ctx context.Context) error {
	var err error
	if stopErr := s.delayedTaskUsecase.StopConsuming(ctx); stopErr != nil {
		err = stopErr
	}
	if stopErr := s.scheduledTaskUsecase.StopConsuming(ctx); stopErr != nil && err == nil {
		err = stopErr
	}
	return err
}
