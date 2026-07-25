package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type ArticlePublishedHandler struct {
	userClientHandler
	contentClientHandler
}

func NewArticlePublishedHandler(
	userClient repo.UserClient,
	contentClient repo.ContentClient,
) *ArticlePublishedHandler {
	return &ArticlePublishedHandler{
		userClientHandler: userClientHandler{
			userClient: userClient,
		},
		contentClientHandler: contentClientHandler{
			contentClient: contentClient,
		},
	}
}

func (h *ArticlePublishedHandler) Build(
	ctx context.Context,
	event *enums.Event,
) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetArticlePublished()
	if payload == nil || payload.GetArticleId() == 0 || h.contentClient == nil {
		return nil, nil
	}
	articleResp, err := h.contentClient.GetArticle(ctx, payload.GetArticleId())
	if err != nil {
		return nil, err
	}
	article := articleResp
	templateData := model.ArticlePublishedTemplateData{
		Article: h.articleTemplateData(article),
	}
	if article == nil || article.AuthorID == 0 || h.userClient == nil {
		return &usecase.NotificationContext{
			EventID:      event.EventId,
			TemplateData: templateData,
		}, nil
	}
	followerResp, err := h.userClient.ListFollowerIDs(ctx, article.AuthorID)
	if err != nil {
		return nil, err
	}
	recipients := make([]*usecase.NotificationRecipient, 0, len(followerResp))
	for _, followerID := range followerResp {
		if followerID == 0 || followerID == article.AuthorID {
			continue
		}
		recipients = append(recipients, &usecase.NotificationRecipient{
			UserID: followerID,
		})
	}
	return &usecase.NotificationContext{
		EventID:      event.EventId,
		TemplateData: templateData,
		Recipients:   recipients,
	}, nil
}
