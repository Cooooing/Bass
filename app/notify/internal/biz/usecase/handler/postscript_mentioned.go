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

type PostscriptMentionedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewPostscriptMentionedHandler(userClient repo.UserAccountRepo, contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *PostscriptMentionedHandler {
	return &PostscriptMentionedHandler{userClientHandler: userClientHandler{userClient: userClient}, contentClientHandler: contentClientHandler{contentClient: contentClient}, notifyUsecase: notifyUsecase}
}

func (*PostscriptMentionedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: enum.EventTypeContentPostscriptMention, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/postscript_mentioned.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/postscript_mentioned.zh_CN.content.txt")}}}
}

func (h *PostscriptMentionedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetArticlePostscriptAdded()
	if payload == nil || payload.GetArticleId() <= 0 || payload.GetPostscriptId() <= 0 || payload.GetSenderId() <= 0 {
		return nil
	}
	postscript, err := h.contentClient.GetPostscript(ctx, payload.GetArticleId(), payload.GetPostscriptId())
	if err != nil || postscript == nil {
		return err
	}
	article, err := h.contentClient.GetArticle(ctx, postscript.ArticleID)
	if err != nil || article == nil {
		return err
	}
	mentioned, err := h.loadAccountsByNames(ctx, util.ParseAtUsernames(postscript.Content))
	if err != nil {
		return err
	}
	users, err := h.loadAccounts(ctx, payload.GetSenderId(), article.AuthorID)
	if err != nil {
		return err
	}
	if author := users[article.AuthorID]; author != nil {
		article.AuthorName = author.Name
		article.AuthorNickname = author.Nickname
	}
	recipients := make([]*model.NotificationRecipient, 0, len(mentioned))
	for userID := range mentioned {
		if userID != payload.GetSenderId() {
			recipients = append(recipients, &model.NotificationRecipient{UserID: userID})
		}
	}
	businessKey := fmt.Sprintf("content:postscript-mention:%d", postscript.ID)
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: enum.EventTypeContentPostscriptMention, ActorUserID: new(payload.SenderId), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.PostscriptMention{Article: h.articleTemplateData(article), Actor: h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()])}, Recipients: recipients})
}
