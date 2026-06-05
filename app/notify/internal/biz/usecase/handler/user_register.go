package handler

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type UserRegisterHandler struct {
	userClientHandler
}

func NewUserRegisterHandler(userClient repo.UserClient) *UserRegisterHandler {
	return &UserRegisterHandler{userClientHandler: userClientHandler{userClient: userClient}}
}

func (h *UserRegisterHandler) Build(ctx context.Context, event *enums.Event) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetUserRegister()
	if payload == nil || payload.GetUserId() == 0 {
		return nil, nil
	}
	user, err := h.loadBasic(ctx, payload.GetUserId())
	if err != nil {
		return nil, err
	}
	templateData := model.UserRegisterTemplateData{
		User: h.templateUser(payload.GetUserId(), user),
	}
	return &usecase.NotificationContext{
		EventID:      event.EventId,
		TemplateData: templateData,
		Recipients: []*usecase.NotificationRecipient{
			{UserID: payload.GetUserId()},
		},
	}, nil
}
