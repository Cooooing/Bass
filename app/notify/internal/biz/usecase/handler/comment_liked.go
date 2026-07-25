package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type CommentLikedHandler struct {
	userClientHandler
	contentClientHandler
}

func NewCommentLikedHandler(
	userClient repo.UserClient,
	contentClient repo.ContentClient,
) *CommentLikedHandler {
	return &CommentLikedHandler{
		userClientHandler: userClientHandler{
			userClient: userClient,
		},
		contentClientHandler: contentClientHandler{
			contentClient: contentClient,
		},
	}
}

func (h *CommentLikedHandler) Build(
	ctx context.Context,
	event *enums.Event,
) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetCommentLiked()
	if payload == nil || payload.GetCommentId() == 0 || h.contentClient == nil {
		return nil, nil
	}
	commentResp, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil {
		return nil, err
	}
	comment := commentResp
	users, err := h.loadAccounts(ctx, payload.GetSenderId())
	if err != nil {
		return nil, err
	}
	templateData := model.CommentLikedTemplateData{
		Comment: h.commentTemplateData(comment),
		Actor:   h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()]),
	}
	recipients := make([]*usecase.NotificationRecipient, 0, 1)
	if comment != nil && comment.UserID != 0 && comment.UserID != payload.GetSenderId() {
		recipients = append(recipients, &usecase.NotificationRecipient{
			UserID: comment.UserID,
		})
	}
	return &usecase.NotificationContext{
		EventID:      event.EventId,
		TemplateData: templateData,
		Recipients:   recipients,
	}, nil
}
