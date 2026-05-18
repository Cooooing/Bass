package handler

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type DefaultHandler struct {
	BaseHandler
}

func NewDefaultHandler(logger log.Logger, notifyService *domain.NotifyService) *DefaultHandler {
	return &DefaultHandler{BaseHandler: NewBaseHandler(logger, notifyService)}
}

func (h *DefaultHandler) Handle(ctx context.Context, event *enums.Event) error {
	if len(event.ReceiverIds) == 0 {
		return nil
	}

	return h.notifyService.Deliver(ctx, &domain.DeliveryRequest{
		EventId:     event.EventId,
		EventType:   event.Type,
		ReceiverIDs: event.ReceiverIds,
		Vars:        struct{}{},
	})
}
