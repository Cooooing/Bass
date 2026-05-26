package handler

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/usecase"
)

type UserRegisterHandler struct{}

func NewUserRegisterHandler() *UserRegisterHandler {
	return &UserRegisterHandler{}
}

func (h *UserRegisterHandler) Build(_ context.Context, event *enums.Event) (*usecase.NotificationIntent, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetUserRegister()
	if payload == nil {
		return nil, nil
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(event.Type)
	if !ok {
		return nil, nil
	}

	intent := &usecase.NotificationIntent{
		EventID:   event.EventId,
		EventType: eventType,
		Vars:      payload,
	}
	switch payload.ContactType {
	case enums.UserRegisterContactType_USER_REGISTER_CONTACT_TYPE_EMAIL:
		if payload.Email == "" {
			return nil, nil
		}
		intent.Email = append(intent.Email, &usecase.EmailInput{Email: payload.Email})
	case enums.UserRegisterContactType_USER_REGISTER_CONTACT_TYPE_PHONE:
		if payload.Phone == "" {
			return nil, nil
		}
		intent.SMS = append(intent.SMS, &usecase.SMSInput{Phone: payload.Phone})
	default:
		return nil, nil
	}

	return intent, nil
}
