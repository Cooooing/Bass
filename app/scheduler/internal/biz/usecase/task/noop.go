package task

import (
	"context"
	"log/slog"
)

type Noop struct {
	logger *slog.Logger
}

func NewNoop(logger *slog.Logger) *Noop {
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
	return "用于验证 scheduler 调度链路的空任务，不执行任何业务操作。"
}

func (n *Noop) Execute(ctx context.Context, _ string) error {
	n.logger.InfoContext(ctx, "noop task executed")
	return nil
}
