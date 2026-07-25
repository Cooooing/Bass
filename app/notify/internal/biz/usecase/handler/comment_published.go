package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type CommentPublishedHandler struct {
	contentClientHandler
}

func NewCommentPublishedHandler(
	contentClient repo.ContentClient,
) *CommentPublishedHandler {
	return &CommentPublishedHandler{
		contentClientHandler: contentClientHandler{
			contentClient: contentClient,
		},
	}
}

func (h *CommentPublishedHandler) Build(
	ctx context.Context,
	event *enums.Event,
) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetCommentPublished()
	if payload == nil || payload.GetCommentId() == 0 || h.contentClient == nil {
		return nil, nil
	}
	commentResp, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil {
		return nil, err
	}
	comment := commentResp
	templateData := model.CommentPublishedTemplateData{
		Comment: h.commentTemplateData(comment),
	}
	recipients := make([]*usecase.NotificationRecipient, 0, 1)
	if comment != nil && comment.ReplyUserID != 0 && comment.ReplyUserID != comment.UserID {
		recipients = append(recipients, &usecase.NotificationRecipient{
			UserID: comment.ReplyUserID,
		})
	} else if comment != nil && comment.Article != nil && comment.Article.AuthorID != 0 && comment.Article.AuthorID != comment.UserID {
		recipients = append(recipients, &usecase.NotificationRecipient{
			UserID: comment.Article.AuthorID,
		})
	}
	return &usecase.NotificationContext{
		EventID:      event.EventId,
		TemplateData: templateData,
		Recipients:   recipients,
	}, nil
}
