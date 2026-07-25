package handler

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/usecase"
)

type UserVerificationCodeHandler struct{}

func NewUserVerificationCodeHandler() *UserVerificationCodeHandler {
	return &UserVerificationCodeHandler{}
}

func (h *UserVerificationCodeHandler) Build(
	ctx context.Context,
	event *enums.Event,
) (*usecase.NotificationContext, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	switch event.Type {
	case enums.EventType_EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE:
		payload := event.GetUserEmailVerificationCode()
		if payload == nil || payload.GetEmail() == "" || payload.GetCode() == "" {
			return nil, nil
		}
		return &usecase.NotificationContext{
			EventID: event.EventId,
			TemplateData: model.VerificationCodeTemplateData{
				Code:           payload.GetCode(),
				ExpiresSeconds: payload.GetExpiresSeconds(),
			},
			Recipients: []*usecase.NotificationRecipient{
				{Email: payload.GetEmail()},
			},
		}, nil
	case enums.EventType_EVENT_TYPE_USER_PHONE_VERIFICATION_CODE:
		payload := event.GetUserPhoneVerificationCode()
		if payload == nil || payload.GetPhone() == "" || payload.GetCode() == "" {
			return nil, nil
		}
		return &usecase.NotificationContext{
			EventID: event.EventId,
			TemplateData: model.VerificationCodeTemplateData{
				Code:           payload.GetCode(),
				ExpiresSeconds: payload.GetExpiresSeconds(),
			},
			Recipients: []*usecase.NotificationRecipient{
				{Phone: payload.GetPhone()},
			},
		}, nil
	default:
		return nil, nil
	}
}
