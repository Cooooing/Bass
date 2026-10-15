package task

import (
	commonscheduler "common/pkg/scheduler"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"log/slog"
	"time"

	schedulerenum "scheduler/internal/enum"
)

type Noop struct {
	logger *slog.Logger
}

func NewNoop(
	logger *slog.Logger,
) *Noop {
	return &Noop{
		logger: logger,
	}
}

func (n *Noop) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameNoop
}

func (n *Noop) Title() string {
	return "空任务"
}

func (n *Noop) Description() string {
	return "用于验证 scheduler 调度链路的空任务。"
}

func (n *Noop) DefaultScheduledTasks() []*DefaultScheduledTask {
	staleAfter := new(time.Duration)
	*staleAfter = 5 * time.Minute
	return []*DefaultScheduledTask{
		{
			TaskKey:       schedulerenum.TaskKeyNoopScheduledDefault,
			Title:         n.Title(),
			Description:   n.Description(),
			Enabled:       false,
			CronSpec:      "0/30 * * * * ?",
			Payload:       commonscheduler.MustMarshalPayload(&schedulerv1.NoopTaskPayload{Name: "scheduled-noop", SleepSeconds: 10}),
			Timeout:       30 * time.Second,
			StaleAfter:    staleAfter,
			MaxAttempts:   1,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest,
			AllowOverlap:  true,
		},
	}
}

func (n *Noop) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{
		{
			TaskKey:       schedulerenum.TaskKeyNoopDelayedDefault,
			Title:         n.Title(),
			Description:   n.Description(),
			Enabled:       false,
			Timeout:       30 * time.Second,
			MaxAttempts:   3,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll,
		},
	}
}

func (n *Noop) Execute(ctx context.Context, payload string) error {
	data := &schedulerv1.NoopTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(payload, data); err != nil {
		return err
	}
	if data.GetSleepSeconds() > 0 {
		select {
		case <-time.After(time.Duration(data.GetSleepSeconds()) * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	n.logger.InfoContext(ctx, "noop task executed", "name", data.GetName(), "sleep_seconds", data.GetSleepSeconds())
	return nil
}
