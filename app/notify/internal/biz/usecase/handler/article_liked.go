package handler

import (
	"context"
	"fmt"

	commonenum "common/pkg/enum"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type ArticleLikedHandler struct {
	articleActorHandler
}

func NewArticleLikedHandler(
	userClient repo.UserAccountRepo,
	contentClient repo.ContentClient,
	notifyUsecase *usecase.NotifyUsecase,
) *ArticleLikedHandler {
	return &ArticleLikedHandler{
		articleActorHandler: articleActorHandler{
			userClientHandler: userClientHandler{
				userClient: userClient,
			},
			contentClientHandler: contentClientHandler{
				contentClient: contentClient,
			},
			notifyUsecase: notifyUsecase,
		},
	}
}

func (h *ArticleLikedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: commonenum.EventTypeContentArticleLike, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/article_liked.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/article_liked.zh_CN.content.txt")}}}
}

func (h *ArticleLikedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetArticleLiked()
	if payload == nil || payload.GetArticleId() == 0 {
		return nil
	}
	notificationContext, err := h.build(ctx, &articleActorBuildReq{
		EventID:   event.GetEventId(),
		ArticleID: payload.GetArticleId(),
		SenderID:  payload.GetSenderId(),
	})
	if err != nil || notificationContext == nil {
		return err
	}
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
		EventID:      event.GetEventId(),
		EventType:    req.EventType,
		ActorUserID:  notificationContext.ActorUserID,
		BusinessKey:  new(fmt.Sprintf("content:article-liked:%d:%d", payload.GetArticleId(), payload.GetSenderId())),
		Language:     req.Language,
		Channels:     []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation},
		TemplateData: notificationContext.TemplateData,
		Recipients:   notificationContext.Recipients,
	})
}
