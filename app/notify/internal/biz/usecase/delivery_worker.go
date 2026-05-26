package usecase

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	deliveryWorkerInterval = 2 * time.Second
	deliveryWorkerLimit    = 50
)

type DeliveryWorker struct {
	log             *log.Helper
	deliveryUsecase *DeliveryUsecase
	ctx             context.Context
	cancel          context.CancelFunc
}

func NewDeliveryWorker(logger log.Logger, deliveryUsecase *DeliveryUsecase) *DeliveryWorker {
	return &DeliveryWorker{
		log:             log.NewHelper(logger),
		deliveryUsecase: deliveryUsecase,
	}
}

func (w *DeliveryWorker) Start(ctx context.Context) error {
	w.ctx, w.cancel = context.WithCancel(ctx)
	go func() {
		ticker := time.NewTicker(deliveryWorkerInterval)
		defer ticker.Stop()
		for {
			if err := w.deliveryUsecase.DispatchPending(w.ctx, deliveryWorkerLimit); err != nil {
				w.log.Errorf("dispatch delivery failed: %v", err)
			}
			select {
			case <-w.ctx.Done():
				return
			case <-ticker.C:
			}
		}
	}()
	return nil
}

func (w *DeliveryWorker) Stop(_ context.Context) error {
	if w.cancel == nil {
		return nil
	}
	w.cancel()
	return nil
}
