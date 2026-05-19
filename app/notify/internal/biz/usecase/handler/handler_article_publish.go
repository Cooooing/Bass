package handler

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/usecase"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticlePublishHandler struct {
	BaseHandler
}

func NewArticlePublishHandler(logger log.Logger, notifyService *usecase.NotifyUsecase) *ArticlePublishHandler {
	return &ArticlePublishHandler{BaseHandler: NewBaseHandler(logger, notifyService)}
}

func (h *ArticlePublishHandler) Handle(ctx context.Context, event *enums.Event) error {
	vars := event.GetArticlePublished()
	if vars == nil {
		return nil
	}
	vars.SenderName = "test"

	return h.notifyService.Deliver(ctx, &usecase.DeliveryRequest{
		EventId:     event.EventId,
		EventType:   event.Type,
		ReceiverIDs: event.ReceiverIds,
		Vars:        vars,
	})
}
