package usecase

import (
	"context"
	"log/slog"
	"time"
)

type DelayedTaskRunner struct {
	logger  *slog.Logger
	usecase *DelayedTaskUsecase
	cancel  context.CancelFunc
}

func NewDelayedTaskRunner(logger *slog.Logger, usecase *DelayedTaskUsecase) *DelayedTaskRunner {
	return &DelayedTaskRunner{logger: logger, usecase: usecase}
}

func (r *DelayedTaskRunner) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	r.cancel = cancel
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-runCtx.Done():
				return
			case <-ticker.C:
				if err := r.usecase.RunDue(runCtx, 100); err != nil && runCtx.Err() == nil {
					r.logger.ErrorContext(runCtx, "scan delayed tasks failed", "err", err)
				}
			}
		}
	}()
	return nil
}

func (r *DelayedTaskRunner) Stop(ctx context.Context) error {
	if r.cancel != nil {
		r.cancel()
	}
	return nil
}
