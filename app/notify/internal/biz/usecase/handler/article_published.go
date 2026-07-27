package handler

import (
	"context"
	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
)

type ArticlePublishedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewArticlePublishedHandler(
	userClient repo.UserAccountRepo,
	contentClient repo.ContentClient,
	notifyUsecase *usecase.NotifyUsecase,
) *ArticlePublishedHandler {
	return &ArticlePublishedHandler{
		userClientHandler: userClientHandler{
			userClient: userClient,
		},
		contentClientHandler: contentClientHandler{
			contentClient: contentClient,
		},
		notifyUsecase: notifyUsecase,
	}
}

func (h *ArticlePublishedHandler) Templates() []*model.NotificationTemplateDefinition {
	return nil
}

func (h *ArticlePublishedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetArticlePublished()
	if payload == nil || payload.GetArticleId() == 0 || h.contentClient == nil {
		return nil
	}
	articleResp, err := h.contentClient.GetArticle(ctx, payload.GetArticleId())
	if err != nil {
		return err
	}
	article := articleResp
	if article != nil && article.AuthorID != 0 {
		users, err := h.loadAccounts(ctx, article.AuthorID)
		if err != nil {
			return err
		}
		if author := users[article.AuthorID]; author != nil {
			article.AuthorName = author.Name
			article.AuthorNickname = author.Nickname
		}
	}
	templateData := templatedata.ArticlePublished{
		Article: h.articleTemplateData(article),
	}
	if article == nil || article.AuthorID == 0 || h.userClient == nil {
		return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
			EventID:      event.GetEventId(),
			EventType:    req.EventType,
			Language:     req.Language,
			Channels:     []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation},
			TemplateData: templateData,
		})
	}
	followerResp, err := h.userClient.ListFollowerIDs(ctx, article.AuthorID)
	if err != nil {
		return err
	}
	recipients := make([]*model.NotificationRecipient, 0, len(followerResp))
	for _, followerID := range followerResp {
		if followerID == 0 || followerID == article.AuthorID {
			continue
		}
		recipients = append(recipients, &model.NotificationRecipient{
			UserID: followerID,
		})
	}
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
		EventID:      event.GetEventId(),
		EventType:    req.EventType,
		Language:     req.Language,
		Channels:     []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation},
		TemplateData: templateData,
		Recipients:   recipients,
	})
}
