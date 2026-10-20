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

type ArticleWatchPollEndedHandler struct {
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewArticleWatchPollEndedHandler(contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *ArticleWatchPollEndedHandler {
	return &ArticleWatchPollEndedHandler{contentClientHandler: contentClientHandler{contentClient: contentClient}, notifyUsecase: notifyUsecase}
}
func (*ArticleWatchPollEndedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: commonenum.EventTypeContentArticleWatchedPollEnded, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/article_watch_poll_ended.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/article_watch_poll_ended.zh_CN.content.txt")}}}
}
func (h *ArticleWatchPollEndedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetArticlePollEnded()
	if payload == nil || payload.GetArticleId() <= 0 {
		return nil
	}
	article, err := h.contentClient.GetArticle(ctx, payload.GetArticleId())
	if err != nil {
		return err
	}
	participantIDs, err := h.contentClient.ListPollParticipants(ctx, payload.GetArticleId())
	if err != nil {
		return err
	}
	watcherIDs, err := h.contentClient.ListArticleWatchers(ctx, payload.GetArticleId())
	if err != nil {
		return err
	}
	seen := map[int64]struct{}{}
	recipients := make([]*model.NotificationRecipient, 0, len(participantIDs)+len(watcherIDs)+1)
	for _, userID := range append(append([]int64{article.AuthorID}, participantIDs...), watcherIDs...) {
		if userID > 0 {
			if _, ok := seen[userID]; !ok {
				recipients = append(recipients, &model.NotificationRecipient{UserID: userID})
				seen[userID] = struct{}{}
			}
		}
	}
	businessKey := fmt.Sprintf("content:article-watch-poll-ended:%d", payload.GetArticleId())
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: commonenum.EventTypeContentArticleWatchedPollEnded, ActorUserID: new(payload.SenderId), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.ArticleWatchActivity{Article: h.articleTemplateData(article)}, Recipients: recipients})
}
