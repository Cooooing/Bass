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

// ArticleAcceptedAnswerHandler 向被采纳回答的作者发送站内信。
type ArticleAcceptedAnswerHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewArticleAcceptedAnswerHandler(userClient repo.UserAccountRepo, contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *ArticleAcceptedAnswerHandler {
	return &ArticleAcceptedAnswerHandler{
		userClientHandler:    userClientHandler{userClient: userClient},
		contentClientHandler: contentClientHandler{contentClient: contentClient},
		notifyUsecase:        notifyUsecase,
	}
}

func (*ArticleAcceptedAnswerHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{
		EventType: commonenum.EventTypeContentArticleAcceptAnswer,
		Channel:   notifyenum.NotificationChannelStation,
		Language:  notifyenum.LanguageZhCN,
		Enabled:   true,
		StationTemplate: &model.NotificationStationTemplateDefinition{
			TitleTemplate:   notifytemplate.MustReadTemplate("station/article_answer_accepted.zh_CN.title.txt"),
			ContentTemplate: notifytemplate.MustReadTemplate("station/article_answer_accepted.zh_CN.content.txt"),
		},
	}}
}

func (h *ArticleAcceptedAnswerHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetArticleAcceptedAnswer()
	if payload == nil || payload.GetArticleId() <= 0 || payload.GetCommentId() <= 0 {
		return nil
	}
	comment, err := h.contentClient.GetComment(ctx, payload.GetCommentId())
	if err != nil {
		return err
	}
	if comment == nil || comment.ArticleID != payload.GetArticleId() || comment.UserID <= 0 || comment.UserID == payload.GetSenderId() {
		return nil
	}
	users, err := h.loadAccounts(ctx, payload.GetSenderId(), comment.UserID)
	if err != nil {
		return err
	}
	businessKey := fmt.Sprintf("content:article-answer-accepted:%d:%d", payload.GetArticleId(), payload.GetCommentId())
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: req.EventType, ActorUserID: new(payload.GetSenderId()), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.ArticleAnswerAccepted{Comment: h.commentTemplateData(comment), Author: h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()])}, Recipients: []*model.NotificationRecipient{{UserID: comment.UserID}}})
}
