package handler

import (
	"context"
	"fmt"

	commonenum "common/pkg/enum"

	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type ArticleWatchPostscriptAddedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewArticleWatchPostscriptAddedHandler(userClient repo.UserAccountRepo, contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *ArticleWatchPostscriptAddedHandler {
	return &ArticleWatchPostscriptAddedHandler{userClientHandler: userClientHandler{userClient: userClient}, contentClientHandler: contentClientHandler{contentClient: contentClient}, notifyUsecase: notifyUsecase}
}

func (*ArticleWatchPostscriptAddedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: commonenum.EventTypeContentArticleWatchedPostscriptAdded, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/article_watch_postscript_added.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/article_watch_postscript_added.zh_CN.content.txt")}}}
}

func (h *ArticleWatchPostscriptAddedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetArticlePostscriptAdded()
	if payload == nil || payload.GetArticleId() <= 0 || payload.GetSenderId() <= 0 {
		return nil
	}
	article, err := h.contentClient.GetArticle(ctx, payload.GetArticleId())
	if err != nil {
		return err
	}
	users, err := h.loadAccounts(ctx, payload.GetSenderId())
	if err != nil {
		return err
	}
	watcherIDs, err := h.contentClient.ListArticleWatchers(ctx, payload.GetArticleId())
	if err != nil {
		return err
	}
	recipients := make([]*model.NotificationRecipient, 0, len(watcherIDs))
	for _, userID := range watcherIDs {
		if userID > 0 && userID != payload.GetSenderId() {
			recipients = append(recipients, &model.NotificationRecipient{UserID: userID})
		}
	}
	businessKey := fmt.Sprintf("content:article-watch-postscript:%d", payload.GetPostscriptId())
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: commonenum.EventTypeContentArticleWatchedPostscriptAdded, ActorUserID: new(payload.SenderId), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.ArticleWatchActivity{Article: h.articleTemplateData(article), Actor: h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()])}, Recipients: recipients})
}
