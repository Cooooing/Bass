package handler

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/usecase"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentActionHandler struct {
	BaseHandler
}

func NewCommentActionHandler(logger log.Logger, notifyService *usecase.NotifyUsecase) *CommentActionHandler {
	return &CommentActionHandler{BaseHandler: NewBaseHandler(logger, notifyService)}
}

func (h *CommentActionHandler) Handle(ctx context.Context, event *enums.Event) error {
	return h.notifyService.Deliver(ctx, &usecase.DeliveryRequest{
		EventId:     event.EventId,
		EventType:   event.Type,
		ReceiverIDs: event.ReceiverIds,
		Vars:        event.GetCommentLiked(),
	})
}
