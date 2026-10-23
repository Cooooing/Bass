package handler

import (
	"context"
	"fmt"

	"common/pkg/enum"
	"common/pkg/util"

	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type GlobalChatMentionedHandler struct {
	userClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewGlobalChatMentionedHandler(userClient repo.UserAccountRepo, notifyUsecase *usecase.NotifyUsecase) *GlobalChatMentionedHandler {
	return &GlobalChatMentionedHandler{userClientHandler: userClientHandler{userClient: userClient}, notifyUsecase: notifyUsecase}
}

func (*GlobalChatMentionedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: enum.EventTypeIMGlobalChatMention, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/global_chat_mentioned.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/global_chat_mentioned.zh_CN.content.txt")}}}
}

func (h *GlobalChatMentionedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" {
		return nil
	}
	payload := event.GetImGlobalChatMessageCreated()
	if payload == nil || payload.GetMessageId() <= 0 || payload.GetSenderId() <= 0 {
		return nil
	}
	mentioned, err := h.loadAccountsByNames(ctx, util.ParseAtUsernames(payload.GetContent()))
	if err != nil {
		return err
	}
	actor, err := h.loadBasic(ctx, payload.GetSenderId())
	if err != nil {
		return err
	}
	recipients := make([]*model.NotificationRecipient, 0, len(mentioned))
	for userID := range mentioned {
		if userID != payload.GetSenderId() {
			recipients = append(recipients, &model.NotificationRecipient{UserID: userID})
		}
	}
	businessKey := fmt.Sprintf("im:global-chat-mention:%d", payload.GetMessageId())
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: enum.EventTypeIMGlobalChatMention, ActorUserID: new(payload.SenderId), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.GlobalChatMention{Actor: h.templateUser(payload.GetSenderId(), actor)}, Recipients: recipients})
}
