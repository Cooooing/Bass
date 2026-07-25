package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type ArticleThankedHandler struct {
	articleActorHandler
}

func NewArticleThankedHandler(
	userClient repo.UserClient,
	contentClient repo.ContentClient,
) *ArticleThankedHandler {
	return &ArticleThankedHandler{
		articleActorHandler: articleActorHandler{
			userClientHandler: userClientHandler{
				userClient: userClient,
			},
			contentClientHandler: contentClientHandler{
				contentClient: contentClient,
			},
		},
	}
}

func (h *ArticleThankedHandler) Build(ctx context.Context, event *enums.Event) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetArticleThanked()
	if payload == nil {
		return nil, nil
	}
	return h.build(ctx, event.EventId, payload.GetArticleId(), payload.GetSenderId())
}
