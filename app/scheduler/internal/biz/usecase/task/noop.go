package task

import (
	"context"
	"encoding/json"
	"log/slog"
	schedulerenum "scheduler/internal/enum"
	"time"
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

func (n *Noop) Name() string {
	return "noop"
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
			Title:         n.Title(),
			Description:   n.Description(),
			Enabled:       false,
			CronSpec:      "0/30 * * * * ?",
			Payload:       `{"name":"scheduled-noop","sleep_seconds":10}`,
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
	data := struct {
		Name         string `json:"name"`
		SleepSeconds int32  `json:"sleep_seconds"`
	}{}
	if payload != "" && payload != "{}" {
		if err := json.Unmarshal([]byte(payload), &data); err != nil {
			return err
		}
	}
	if data.SleepSeconds > 0 {
		select {
		case <-time.After(time.Duration(data.SleepSeconds) * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	n.logger.InfoContext(ctx, "noop task executed", "name", data.Name, "sleep_seconds", data.SleepSeconds)
	return nil
}
