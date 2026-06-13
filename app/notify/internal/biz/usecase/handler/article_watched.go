package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type ArticleWatchedHandler struct {
	articleActorHandler
}

func NewArticleWatchedHandler(userClient repo.UserClient, contentClient repo.ContentClient) *ArticleWatchedHandler {
	return &ArticleWatchedHandler{articleActorHandler: articleActorHandler{
		userClientHandler:    userClientHandler{userClient: userClient},
		contentClientHandler: contentClientHandler{contentClient: contentClient},
	}}
}

func (h *ArticleWatchedHandler) Build(ctx context.Context, event *enums.Event) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetArticleWatched()
	if payload == nil {
		return nil, nil
	}
	return h.build(ctx, event.EventId, payload.GetArticleId(), payload.GetSenderId())
}
