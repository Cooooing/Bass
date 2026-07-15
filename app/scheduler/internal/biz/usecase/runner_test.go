package usecase

import (
	"context"
	"io"
	"log/slog"
	taskimpl "scheduler/internal/biz/usecase/task"
	"scheduler/internal/config"
	schedulerenum "scheduler/internal/enum"
	"testing"
	"time"
)

func TestSchedulerRunnerStartLoadsEnabledTasksWithList(t *testing.T) {
	taskRepo := &fakeTaskRepo{task: testTask(true)}
	runner := newTestSchedulerRunner(t, taskRepo, &fakeExecutionRepo{created: true}, &fakeTask{})
	if err := runner.Start(context.Background()); err != nil {
		t.Fatalf("Start returned error: %v", err)
	}
	defer runner.Stop(context.Background())

	if taskRepo.listReq == nil || taskRepo.listReq.Enabled == nil || !*taskRepo.listReq.Enabled {
		t.Fatalf("expected Start to load enabled tasks through List, got %#v", taskRepo.listReq)
	}
}

func TestSchedulerRunnerScheduleFieldSkipsDatabase(t *testing.T) {
	task := testTask(true)
	executionRepo := &fakeExecutionRepo{created: true}
	taskLockRepo := &fakeTaskLockRepo{
		schedules: map[string]string{},
	}
	runner := newTestSchedulerRunnerWithTaskLock(t, &fakeTaskRepo{task: task}, executionRepo, &fakeTask{}, taskLockRepo)
	scheduledAt := time.Date(2026, 7, 4, 10, 0, 0, 0, time.UTC)
	taskLockRepo.schedules[executionPeriodKey(task.ID, scheduledAt)] = "token"

	if _, err := runner.taskUsecase.ScheduleExecution(context.Background(), task, scheduledAt, schedulerenum.TaskTriggerTypeSchedule); err != nil {
		t.Fatalf("ScheduleExecution returned error: %v", err)
	}

	if executionRepo.createCalled != 0 {
		t.Fatalf("expected no database create when schedule field exists, got %d", executionRepo.createCalled)
	}
	if executionRepo.getCalled != 0 {
		t.Fatalf("expected no database get when schedule field exists, got %d", executionRepo.getCalled)
	}
}

func TestSchedulerRunnerScheduleConflictSkipsExecution(t *testing.T) {
	task := testTask(true)
	executionRepo := &fakeExecutionRepo{scheduleConflict: true}
	runner := newTestSchedulerRunner(t, &fakeTaskRepo{task: task}, executionRepo, &fakeTask{})
	scheduledAt := time.Date(2026, 7, 4, 10, 0, 0, 0, time.UTC)

	if _, err := runner.taskUsecase.ScheduleExecution(context.Background(), task, scheduledAt, schedulerenum.TaskTriggerTypeSchedule); err != nil {
		t.Fatalf("ScheduleExecution returned error: %v", err)
	}

	if executionRepo.createCalled != 1 {
		t.Fatalf("expected one create attempt, got %d", executionRepo.createCalled)
	}
}

func newTestSchedulerRunner(t *testing.T, taskRepo *fakeTaskRepo, executionRepo *fakeExecutionRepo, task taskimpl.Task) *SchedulerRunner {
	t.Helper()
	return newTestSchedulerRunnerWithTaskLock(t, taskRepo, executionRepo, task, &fakeTaskLockRepo{})
}

func newTestSchedulerRunnerWithTaskLock(t *testing.T, taskRepo *fakeTaskRepo, executionRepo *fakeExecutionRepo, task taskimpl.Task, taskLockRepo *fakeTaskLockRepo) *SchedulerRunner {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	taskUsecase := NewTaskUsecase(
		logger,
		testConf(),
		testTx,
		taskRepo,
		&fakeTaskVersionRepo{},
		executionRepo,
		taskLockRepo,
		map[string]taskimpl.Task{"noop": task},
		&fakeTaskEventBus{},
		&fakeAlert{},
	)
	return NewSchedulerRunner(
		logger,
		&config.Bootstrap{Scheduler: testConf().Scheduler},
		taskRepo,
		taskUsecase,
		&fakeTaskEventBus{},
	)
}
