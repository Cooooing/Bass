package handler

import (
	"context"
	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
)

type CommentPublishedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewCommentPublishedHandler(
	userClient repo.UserAccountRepo,
	contentClient repo.ContentClient,
	notifyUsecase *usecase.NotifyUsecase,
) *CommentPublishedHandler {
	return &CommentPublishedHandler{
		userClientHandler: userClientHandler{
			userClient: userClient,
		},
		contentClientHandler: contentClientHandler{
			contentClient: contentClient,
		},
		notifyUsecase: notifyUsecase,
	}
}

func (h *CommentPublishedHandler) Templates() []*model.NotificationTemplateDefinition {
	return nil
}

func (h *CommentPublishedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetCommentPublished()
	if payload == nil || payload.GetCommentId() == 0 || h.contentClient == nil {
		return nil
	}
	commentResp, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil {
		return err
	}
	comment := commentResp
	if comment != nil {
		userIDs := []int64{comment.UserID, comment.ReplyUserID}
		if comment.Article != nil {
			userIDs = append(userIDs, comment.Article.AuthorID)
		}
		users, err := h.loadAccounts(ctx, userIDs...)
		if err != nil {
			return err
		}
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
	templateData := templatedata.CommentPublished{
		Comment: h.commentTemplateData(comment),
	}
	recipients := make([]*model.NotificationRecipient, 0, 1)
	if comment != nil && comment.ReplyUserID != 0 && comment.ReplyUserID != comment.UserID {
		recipients = append(recipients, &model.NotificationRecipient{
			UserID: comment.ReplyUserID,
		})
	} else if comment != nil && comment.Article != nil && comment.Article.AuthorID != 0 && comment.Article.AuthorID != comment.UserID {
		recipients = append(recipients, &model.NotificationRecipient{
			UserID: comment.Article.AuthorID,
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
