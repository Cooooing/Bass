package handler

import (
	"context"
	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
)

type UserFollowHandler struct {
	userClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewUserFollowHandler(
	userClient repo.UserAccountRepo,
	notifyUsecase *usecase.NotifyUsecase,
) *UserFollowHandler {
	return &UserFollowHandler{
		userClientHandler: userClientHandler{
			userClient: userClient,
		},
		notifyUsecase: notifyUsecase,
	}
}

func (h *UserFollowHandler) Templates() []*model.NotificationTemplateDefinition {
	return nil
}

func (h *UserFollowHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetUserFollow()
	if payload == nil || payload.GetFollowedId() == 0 {
		return nil
	}
	users, err := h.loadAccounts(ctx, payload.GetSenderId(), payload.GetFollowedId())
	if err != nil {
		return err
	}
	templateData := templatedata.UserFollow{
		Follower: h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()]),
		Followed: h.templateUser(payload.GetFollowedId(), users[payload.GetFollowedId()]),
	}
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
		EventID:      event.GetEventId(),
		EventType:    req.EventType,
		Language:     req.Language,
		Channels:     []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation},
		TemplateData: templateData,
		Recipients: []*model.NotificationRecipient{
			{UserID: payload.GetFollowedId()},
		},
	})
}
