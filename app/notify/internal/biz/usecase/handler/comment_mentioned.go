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

type CommentMentionedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewCommentMentionedHandler(userClient repo.UserAccountRepo, contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *CommentMentionedHandler {
	return &CommentMentionedHandler{userClientHandler: userClientHandler{userClient: userClient}, contentClientHandler: contentClientHandler{contentClient: contentClient}, notifyUsecase: notifyUsecase}
}

func (*CommentMentionedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: enum.EventTypeContentCommentMention, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/comment_mentioned.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/comment_mentioned.zh_CN.content.txt")}}}
}

func (h *CommentMentionedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetCommentPublished()
	if payload == nil || payload.GetCommentId() <= 0 {
		return nil
	}
	comment, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil || comment == nil || comment.UserID <= 0 {
		return err
	}
	mentioned, err := h.loadAccountsByNames(ctx, util.ParseAtUsernames(comment.Content))
	if err != nil {
		return err
	}
	articleAuthorID := int64(0)
	if comment.Article != nil {
		articleAuthorID = comment.Article.AuthorID
	}
	users, err := h.loadAccounts(ctx, comment.UserID, articleAuthorID)
	if err != nil {
		return err
	}
	if comment.Article != nil {
		if author := users[articleAuthorID]; author != nil {
			comment.Article.AuthorName = author.Name
			comment.Article.AuthorNickname = author.Nickname
		}
	}
	if actor := users[comment.UserID]; actor != nil {
		comment.UserName = actor.Name
		comment.UserNickname = actor.Nickname
	}
	recipients := make([]*model.NotificationRecipient, 0, len(mentioned))
	for userID := range mentioned {
		if userID != comment.UserID {
			recipients = append(recipients, &model.NotificationRecipient{UserID: userID})
		}
	}
	businessKey := fmt.Sprintf("content:comment-mention:%d", comment.ID)
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: enum.EventTypeContentCommentMention, ActorUserID: new(comment.UserID), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.CommentMention{Comment: h.commentTemplateData(comment), Actor: h.templateUser(comment.UserID, users[comment.UserID])}, Recipients: recipients})
}
