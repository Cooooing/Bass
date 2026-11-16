package handler

import (
	"context"

	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type UserVerificationCodeHandler struct {
	notifyUsecase *usecase.NotifyUsecase
}

func NewUserVerificationCodeHandler(
	notifyUsecase *usecase.NotifyUsecase,
) *UserVerificationCodeHandler {
	return &UserVerificationCodeHandler{
		notifyUsecase: notifyUsecase,
	}
}

func (h *UserVerificationCodeHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{
		{
			NotificationType: notifyenum.NotificationTypeUserEmailVerificationCode,
			Channel:          notifyenum.NotificationChannelEmail,
			Language:         notifyenum.LanguageZhCN,
			Enabled:          true,
			EmailTemplate: &model.NotificationEmailTemplateDefinition{
				SubjectTemplate: notifytemplate.MustReadTemplate("email/user_email_verification_code.zh_CN.subject.txt"),
				BodyTemplate:    notifytemplate.MustReadTemplate("email/user_email_verification_code.zh_CN.body.html"),
				ContentType:     "text/html",
			},
		},
		{
			NotificationType: notifyenum.NotificationTypeUserPhoneVerificationCode,
			Channel:          notifyenum.NotificationChannelTencentSMS,
			Language:         notifyenum.LanguageZhCN,
			Enabled:          true,
			TencentSMSTemplate: &model.NotificationTencentSMSTemplateDefinition{
				ParamTemplates: []string{"{{.Code}}"},
			},
		},
	}
}

func (h *UserVerificationCodeHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	switch req.NotificationType {
	case notifyenum.NotificationTypeUserEmailVerificationCode:
		payload := event.GetUserEmailVerificationCode()
		if payload == nil || payload.GetEmail() == "" || payload.GetCode() == "" {
			return nil
		}
		return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
			EventID:          event.GetEventId(),
			NotificationType: req.NotificationType,
			Language:         req.Language,
			Channels:         []notifyenum.NotificationChannel{notifyenum.NotificationChannelEmail},
			TemplateData: templatedata.VerificationCode{
				Code:           payload.GetCode(),
				ExpiresSeconds: payload.GetExpiresSeconds(),
			},
			Recipients: []*model.NotificationRecipient{
				{Email: payload.GetEmail()},
			},
		})
	case notifyenum.NotificationTypeUserPhoneVerificationCode:
		payload := event.GetUserPhoneVerificationCode()
		if payload == nil || payload.GetPhone() == "" || payload.GetCode() == "" {
			return nil
		}
		return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
			EventID:          event.GetEventId(),
			NotificationType: req.NotificationType,
			Language:         req.Language,
			Channels:         []notifyenum.NotificationChannel{notifyenum.NotificationChannelTencentSMS},
			TemplateData: templatedata.VerificationCode{
				Code:           payload.GetCode(),
				ExpiresSeconds: payload.GetExpiresSeconds(),
			},
			Recipients: []*model.NotificationRecipient{
				{Phone: payload.GetPhone()},
			},
		})
	default:
		return nil
	}
}
