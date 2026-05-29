package handler

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type UserFollowHandler struct {
	userClientHandler
}

func NewUserFollowHandler(userClient repo.UserClient) *UserFollowHandler {
	return &UserFollowHandler{userClientHandler: userClientHandler{userClient: userClient}}
}

func (h *UserFollowHandler) Build(ctx context.Context, event *enums.Event) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetUserFollow()
	if payload == nil || payload.GetFollowedId() == 0 {
		return nil, nil
	}
	users, err := h.loadAccounts(ctx, payload.GetSenderId(), payload.GetFollowedId())
	if err != nil {
		return nil, err
	}
	templateData := model.UserFollowTemplateData{
		Follower: h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()]),
		Followed: h.templateUser(payload.GetFollowedId(), users[payload.GetFollowedId()]),
	}
	return &usecase.NotificationContext{
		EventID:      event.EventId,
		TemplateData: templateData,
		Recipients: []*usecase.NotificationRecipient{
			{UserID: payload.GetFollowedId()},
		},
	}, nil
}
