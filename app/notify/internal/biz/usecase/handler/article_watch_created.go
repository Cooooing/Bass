package handler

import (
	"context"
	"fmt"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type ArticleWatchCreatedHandler struct {
	articleActorHandler
}

func NewArticleWatchCreatedHandler(
	userClient repo.UserAccountRepo,
	contentClient repo.ContentClient,
	notifyUsecase *usecase.NotifyUsecase,
) *ArticleWatchCreatedHandler {
	return &ArticleWatchCreatedHandler{userClient: userClient, contentClient: contentClient, notifyUsecase: notifyUsecase}
}

func (*ArticleWatchCreatedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{
		NotificationType: notifyenum.NotificationTypeContentArticleWatch,
		Channel:          notifyenum.NotificationChannelStation,
		Language:         notifyenum.LanguageZhCN,
		Enabled:          true,
		StationTemplate: &model.NotificationStationTemplateDefinition{
			TitleTemplate:   notifytemplate.MustReadTemplate("station/article_watched.zh_CN.title.txt"),
			ContentTemplate: notifytemplate.MustReadTemplate("station/article_watched.zh_CN.content.txt"),
		},
	}}
}

func (h *ArticleWatchCreatedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetArticleWatched()
	if payload == nil || payload.GetArticleId() <= 0 {
		return nil
	}
	notificationContext, err := h.build(ctx, &articleActorBuildReq{EventID: event.GetEventId(), ArticleID: payload.GetArticleId(), SenderID: payload.GetSenderId()})
	if err != nil || notificationContext == nil {
		return err
	}
	businessKey := fmt.Sprintf("content:article-watched:%d:%d", payload.GetArticleId(), payload.GetSenderId())
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), NotificationType: req.NotificationType, ActorUserID: notificationContext.ActorUserID, BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: notificationContext.TemplateData, Recipients: notificationContext.Recipients})
}
