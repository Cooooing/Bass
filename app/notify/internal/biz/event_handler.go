package biz

import (
	"common/pkg/constant"
	"context"
	"time"

	"github.com/panjf2000/ants/v2"
)

type EventHandler struct {
	*BaseDomain
	pool   *ants.Pool
	ctx    context.Context
	cancel context.CancelFunc
}

func NewEventHandler(base *BaseDomain) (*EventHandler, func(), error) {
	pool, err := ants.NewPool(16, ants.WithNonblocking(false))
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	h := &EventHandler{
		BaseDomain: base,
		pool:       pool,
		ctx:        ctx,
		cancel:     cancel,
	}
	return h, h.CleanUp, nil
}

func (h *EventHandler) Handle() {
	msgs, ch, err := h.rabbitmq.Consume(constant.QueueNotify.String())
	if err != nil {
		h.log.Error("consume error: %v", err)
		return
	}
	defer ch.Close()
	for {
		select {
		case <-h.ctx.Done():
			h.log.Info("Handle exited due to context cancel")
			return
		case msg, ok := <-msgs:
			if !ok {
				h.log.Info("Channel closed")
				return
			}

			// 提交给线程池
			m := msg
			h.pool.Submit(func() {
				h.log.Infof("receive message: %s", string(m.Body))

				// 模拟业务处理
				time.Sleep(3 * time.Second)

				// ack 消息
				if err := m.Ack(false); err != nil {
					h.log.Errorf("ack failed: %v", err)
					// 可选择重试或 nack
					// m.Nack(false, true)
				}
			})
		}
	}
}

func (h *EventHandler) CleanUp() {
	h.cancel()
	h.pool.Release()
}
