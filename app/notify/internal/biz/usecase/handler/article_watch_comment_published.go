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

type ArticleWatchCommentPublishedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewArticleWatchCommentPublishedHandler(
	userClient repo.UserAccountRepo,
	contentClient repo.ContentClient,
	notifyUsecase *usecase.NotifyUsecase,
) *ArticleWatchCommentPublishedHandler {
	return &ArticleWatchCommentPublishedHandler{
		userClientHandler:    userClientHandler{userClient: userClient},
		contentClientHandler: contentClientHandler{contentClient: contentClient},
		notifyUsecase:        notifyUsecase,
	}
}

func (*ArticleWatchCommentPublishedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{
		EventType: commonenum.EventTypeContentArticleWatchedCommentPublished,
		Channel:   notifyenum.NotificationChannelStation,
		Language:  notifyenum.LanguageZhCN,
		Enabled:   true,
		StationTemplate: &model.NotificationStationTemplateDefinition{
			TitleTemplate:   notifytemplate.MustReadTemplate("station/article_watch_comment_published.zh_CN.title.txt"),
			ContentTemplate: notifytemplate.MustReadTemplate("station/article_watch_comment_published.zh_CN.content.txt"),
		},
	}}
}

func (h *ArticleWatchCommentPublishedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetCommentPublished()
	if payload == nil || payload.GetCommentId() <= 0 {
		return nil
	}
	comment, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil || comment == nil || comment.ArticleID <= 0 || comment.UserID <= 0 {
		return err
	}
	article, err := h.contentClient.GetArticle(ctx, comment.ArticleID)
	if err != nil {
		return err
	}
	users, err := h.loadAccounts(ctx, comment.UserID)
	if err != nil {
		return err
	}
	watcherIDs, err := h.contentClient.ListArticleWatchers(ctx, comment.ArticleID)
	if err != nil {
		return err
	}
	recipients := make([]*model.NotificationRecipient, 0, len(watcherIDs))
	for _, userID := range watcherIDs {
		if userID > 0 && userID != comment.UserID {
			recipients = append(recipients, &model.NotificationRecipient{UserID: userID})
		}
	}
	businessKey := fmt.Sprintf("content:article-watch-comment:%d", comment.ID)
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: commonenum.EventTypeContentArticleWatchedCommentPublished, ActorUserID: new(comment.UserID), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.ArticleWatchActivity{Article: h.articleTemplateData(article), Actor: h.templateUser(comment.UserID, users[comment.UserID])}, Recipients: recipients})
}
