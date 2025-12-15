package biz

import (
	"runtime/debug"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/ants/v2"
)

type EventPool struct {
	pool *ants.Pool
	size int
}

func NewEventPool(log *log.Helper) (*EventPool, func(), error) {
	size := 16
	pool, err := ants.NewPool(
		size,
		ants.WithNonblocking(false),
		ants.WithPanicHandler(func(err interface{}) {
			// 错误兜底逻辑
			log.Errorf("[ants] worker panic recovered: %v\n%s", err, debug.Stack())
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
