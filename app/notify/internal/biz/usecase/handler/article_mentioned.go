package handler

import (
	"context"
	"fmt"

	"common/pkg/enum"
	"common/pkg/util"

	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type ArticleMentionedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewArticleMentionedHandler(userClient repo.UserAccountRepo, contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *ArticleMentionedHandler {
	return &ArticleMentionedHandler{userClientHandler: userClientHandler{userClient: userClient}, contentClientHandler: contentClientHandler{contentClient: contentClient}, notifyUsecase: notifyUsecase}
}

func (*ArticleMentionedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: enum.EventTypeContentArticleMention, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/article_mentioned.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/article_mentioned.zh_CN.content.txt")}}}
}

func (h *ArticleMentionedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetArticlePublished()
	if payload == nil || payload.GetArticleId() <= 0 {
		return nil
	}
	article, err := h.contentClient.GetArticle(ctx, payload.GetArticleId())
	if err != nil || article == nil || article.AuthorID <= 0 {
		return err
	}
	mentioned, err := h.loadAccountsByNames(ctx, util.ParseAtUsernames(article.Content))
	if err != nil {
		return err
	}
	author, err := h.loadBasic(ctx, article.AuthorID)
	if err != nil {
		return err
	}
	recipients := make([]*model.NotificationRecipient, 0, len(mentioned))
	for userID := range mentioned {
		if userID != article.AuthorID {
			recipients = append(recipients, &model.NotificationRecipient{UserID: userID})
		}
	}
	businessKey := fmt.Sprintf("content:article-mention:%d", article.ID)
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: enum.EventTypeContentArticleMention, ActorUserID: new(article.AuthorID), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.ArticleMention{Article: h.articleTemplateData(article), Actor: h.templateUser(article.AuthorID, author)}, Recipients: recipients})
}
