package handler

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type UserRegisterHandler struct {
	userClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewUserRegisterHandler(
	userClient repo.UserAccountRepo,
	notifyUsecase *usecase.NotifyUsecase,
) *UserRegisterHandler {
	return &UserRegisterHandler{
		userClientHandler: userClientHandler{
			userClient: userClient,
		},
		notifyUsecase: notifyUsecase,
	}
}

func (h *UserRegisterHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{
		{
			EventType: commonenum.EventTypeUserRegister,
			Channel:   notifyenum.NotificationChannelStation,
			Language:  notifyenum.LanguageZhCN,
			Enabled:   true,
			StationTemplate: &model.NotificationStationTemplateDefinition{
				TitleTemplate:   notifytemplate.MustReadTemplate("station/user_register.zh_CN.title.txt"),
				ContentTemplate: notifytemplate.MustReadTemplate("station/user_register.zh_CN.content.txt"),
			},
		},
	}
}

func (h *UserRegisterHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetUserRegister()
	if payload == nil || payload.GetUserId() == 0 {
		return nil
	}
	user, err := h.loadBasic(ctx, payload.GetUserId())
	if err != nil {
		return err
	}
	templateData := templatedata.UserRegister{
		User: h.templateUser(payload.GetUserId(), user),
	}
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{
		EventID:      event.GetEventId(),
		EventType:    req.EventType,
		Language:     req.Language,
		Channels:     []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation},
		TemplateData: templateData,
		Recipients: []*model.NotificationRecipient{
			{UserID: payload.GetUserId()},
		},
	})
}
