package handler

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type FollowHandler struct {
	BaseHandler
}

func NewFollowHandler(logger log.Logger, notifyService *domain.NotifyService) *FollowHandler {
	return &FollowHandler{BaseHandler: NewBaseHandler(logger, notifyService)}
}

func (h *FollowHandler) Handle(ctx context.Context, event *enums.Event) error {
	return h.notifyService.Deliver(ctx, &domain.DeliveryRequest{
		EventId:     event.EventId,
		EventType:   event.Type,
		ReceiverIDs: event.ReceiverIds,
		Vars:        event.GetUserFollowCreated(),
	})
}
