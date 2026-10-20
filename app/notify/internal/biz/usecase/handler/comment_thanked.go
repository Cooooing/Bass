package handler

import (
	"context"
	"fmt"

	commonenum "common/pkg/enum"

	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	notifytemplate "notify/template"
)

type CommentThankedHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewCommentThankedHandler(userClient repo.UserAccountRepo, contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *CommentThankedHandler {
	return &CommentThankedHandler{userClientHandler: userClientHandler{userClient: userClient}, contentClientHandler: contentClientHandler{contentClient: contentClient}, notifyUsecase: notifyUsecase}
}

func (*CommentThankedHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: commonenum.EventTypeContentCommentThank, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/comment_thanked.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/comment_thanked.zh_CN.content.txt")}}}
}

func (h *CommentThankedHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetCommentThanked()
	if payload == nil || payload.GetCommentId() <= 0 || payload.GetSenderId() <= 0 {
		return nil
	}
	comment, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil {
		return err
	}
	if comment == nil || comment.UserID <= 0 || comment.UserID == payload.GetSenderId() {
		return nil
	}
	users, err := h.loadAccounts(ctx, payload.GetSenderId(), comment.UserID)
	if err != nil {
		return err
	}
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: req.EventType, ActorUserID: new(payload.SenderId), BusinessKey: new(fmt.Sprintf("content:comment-thanked:%d:%d", payload.GetCommentId(), payload.GetSenderId())), Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.CommentLiked{Comment: h.commentTemplateData(comment), Actor: h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()])}, Recipients: []*model.NotificationRecipient{{UserID: comment.UserID}}})
}
