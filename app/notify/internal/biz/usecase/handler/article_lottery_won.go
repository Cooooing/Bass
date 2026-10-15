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

// ArticleLotteryWonHandler 向抽奖中奖者发送站内信。
type ArticleLotteryWonHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

func NewArticleLotteryWonHandler(userClient repo.UserAccountRepo, contentClient repo.ContentClient, notifyUsecase *usecase.NotifyUsecase) *ArticleLotteryWonHandler {
	return &ArticleLotteryWonHandler{userClientHandler: userClientHandler{userClient: userClient}, contentClientHandler: contentClientHandler{contentClient: contentClient}, notifyUsecase: notifyUsecase}
}

func (*ArticleLotteryWonHandler) Templates() []*model.NotificationTemplateDefinition {
	return []*model.NotificationTemplateDefinition{{EventType: commonenum.EventTypeContentArticleLotteryWon, Channel: notifyenum.NotificationChannelStation, Language: notifyenum.LanguageZhCN, Enabled: true, StationTemplate: &model.NotificationStationTemplateDefinition{TitleTemplate: notifytemplate.MustReadTemplate("station/article_lottery_won.zh_CN.title.txt"), ContentTemplate: notifytemplate.MustReadTemplate("station/article_lottery_won.zh_CN.content.txt")}}}
}

func (h *ArticleLotteryWonHandler) Handle(ctx context.Context, req *usecase.EventHandleReq) error {
	event := req.Event
	if event == nil || event.GetEventId() == "" || h.contentClient == nil {
		return nil
	}
	payload := event.GetArticleLotteryWon()
	if payload == nil || payload.GetArticleId() <= 0 || payload.GetWinnerId() <= 0 || payload.GetWinnerId() == payload.GetSenderId() {
		return nil
	}
	article, err := h.contentClient.GetArticle(ctx, payload.GetArticleId())
	if err != nil {
		return err
	}
	users, err := h.loadAccounts(ctx, payload.GetSenderId())
	if err != nil {
		return err
	}
	businessKey := fmt.Sprintf("content:article-lottery-won:%d:%d", payload.GetArticleId(), payload.GetWinnerId())
	return h.notifyUsecase.Send(ctx, &usecase.NotifySendReq{EventID: event.GetEventId(), EventType: req.EventType, ActorUserID: new(payload.SenderId), BusinessKey: &businessKey, Language: req.Language, Channels: []notifyenum.NotificationChannel{notifyenum.NotificationChannelStation}, TemplateData: templatedata.ArticleLotteryWon{Article: h.articleTemplateData(article), Author: h.templateUser(payload.GetSenderId(), users[payload.GetSenderId()]), Prize: fmt.Sprintf("奖项 #%d", payload.GetPrizeId())}, Recipients: []*model.NotificationRecipient{{UserID: payload.GetWinnerId()}}})
}
