package util

import (
	"common/pkg/constant"
	"log/slog"
	"runtime/debug"

	"github.com/panjf2000/ants/v2"
)

type EventPool struct {
	pool *ants.Pool
	size int
}

func NewEventPool(logger *slog.Logger) (*EventPool, func(), error) {
	size := 16
	pool, err := ants.NewPool(
		size,
		ants.WithNonblocking(false),
		ants.WithPanicHandler(func(err interface{}) {
			logger.Error("event worker panic recovered", constant.LogFieldKind, constant.LogKindSystem, "panic", err, "stack", string(debug.Stack()))
		}),
	)
	e := &EventPool{
		pool: pool,
		size: size,
	}
	cleanup := func() {
		e.pool.Release()
	}
	return e, cleanup, err
}

func (p *EventPool) Submit(task func()) error {
	return p.pool.Submit(task)
}
