package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type ArticleLikedHandler struct {
	articleActorHandler
}

func NewArticleLikedHandler(userClient repo.UserClient, contentClient repo.ContentClient) *ArticleLikedHandler {
	return &ArticleLikedHandler{articleActorHandler: articleActorHandler{
		userClientHandler:    userClientHandler{userClient: userClient},
		contentClientHandler: contentClientHandler{contentClient: contentClient},
	}}
}

func (h *ArticleLikedHandler) Build(ctx context.Context, event *enums.Event) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetArticleLiked()
	if payload == nil {
		return nil, nil
	}
	return h.build(ctx, event.EventId, payload.GetArticleId(), payload.GetSenderId())
}
