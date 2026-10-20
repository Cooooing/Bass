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

type CommentLikedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewCommentLikedHandler(
	userClient repo.UserAccountRepo,
	contentClient repo.ContentClient,
	notifyUsecase *usecase.NotifyUsecase,
) *CommentLikedHandler {
	return &CommentLikedHandler{
		userClientHandler: userClientHandler{
			userClient: userClient,
		},
		contentClientHandler: contentClientHandler{
			contentClient: contentClient,
		},
		notifyUsecase: notifyUsecase,
	}
}

func (h *CommentLikedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: commonenum.EventTypeContentCommentLike, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/comment_liked.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/comment_liked.zh_CN.content.txt")}}}
}

func (h *CommentLikedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetCommentLiked()
	if payload == nil || payload.GetCommentId() == 0 || h.contentClient == nil {
		return nil
	}
	commentResp, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil {
		return err
	}
	comment := commentResp
	userIDs := []int64{payload.GetSenderId()}
	if comment != nil {
		userIDs = append(userIDs, comment.UserID, comment.ReplyUserID)
		if comment.Article != nil {
			userIDs = append(userIDs, comment.Article.AuthorID)
		}
	}
	users, err := h.loadAccounts(ctx, userIDs...)
	if err != nil {
		return err
	}
	if comment != nil {
		if user := users[comment.UserID]; user != nil {
			comment.UserName = user.Name
			comment.UserNickname = user.Nickname
		}
		if replyUser := users[comment.ReplyUserID]; replyUser != nil {
			comment.ReplyUserName = replyUser.Name
		}
		if comment.Article != nil {
			if author := users[comment.Article.AuthorID]; author != nil {
				comment.Article.AuthorName = author.Name
				comment.Article.AuthorNickname = author.Nickname
			}
		}
	}
	templateData := templatedata.CommentLiked{
		Comment: h.commentTemplateData(comment),
		Actor:   h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()]),
	}
	recipients := make([]*model.NotificationRecipient, 0, 1)
	if comment != nil && comment.UserID != 0 && comment.UserID != payload.GetSenderId() {
		recipients = append(recipients, &model.NotificationRecipient{
			UserID: comment.UserID,
		})
	}
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
		EventID:      event.GetEventId(),
		EventType:    req.EventType,
		ActorUserID:  new(payload.GetSenderId()),
		BusinessKey:  new(fmt.Sprintf("content:comment-liked:%d:%d", payload.GetCommentId(), payload.GetSenderId())),
		Language:     req.Language,
		Channels:     []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation},
		TemplateData: templateData,
		Recipients:   recipients,
	})
}
