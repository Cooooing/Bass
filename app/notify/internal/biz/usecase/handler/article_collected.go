package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type ArticleCollectedHandler struct {
	articleActorHandler
}

func NewArticleCollectedHandler(userClient repo.UserClient, contentClient repo.ContentClient) *ArticleCollectedHandler {
	return &ArticleCollectedHandler{articleActorHandler: articleActorHandler{
		userClientHandler:    userClientHandler{userClient: userClient},
		contentClientHandler: contentClientHandler{contentClient: contentClient},
	}}
}

func (h *ArticleCollectedHandler) Build(ctx context.Context, event *enums.Event) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetArticleCollected()
	if payload == nil {
		return nil, nil
	}
	return h.build(ctx, event.EventId, payload.GetArticleId(), payload.GetSenderId())
}
