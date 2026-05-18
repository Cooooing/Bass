package handler

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleActionHandler struct {
	BaseHandler
}

func NewArticleActionHandler(logger log.Logger, notifyService *domain.NotifyService) *ArticleActionHandler {
	return &ArticleActionHandler{BaseHandler: NewBaseHandler(logger, notifyService)}
}

func (h *ArticleActionHandler) Handle(ctx context.Context, event *enums.Event) error {
	var payload any
	switch event.Type {
	case enums.EventType_EVENT_TYPE_ARTICLE_LIKED:
		payload = event.GetArticleLiked()
	case enums.EventType_EVENT_TYPE_ARTICLE_THANKED:
		payload = event.GetArticleThanked()
	case enums.EventType_EVENT_TYPE_ARTICLE_COLLECTED:
		payload = event.GetArticleCollected()
	case enums.EventType_EVENT_TYPE_ARTICLE_WATCHED:
		payload = event.GetArticleWatched()
	}

	return h.notifyService.Deliver(ctx, &domain.DeliveryRequest{
		EventId:     event.EventId,
		EventType:   event.Type,
		ReceiverIDs: event.ReceiverIds,
		Vars:        payload,
	})
}
