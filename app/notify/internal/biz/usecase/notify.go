package usecase

import (
	"context"
	"fmt"
	"log/slog"

	commonenum "common/pkg/enum"
	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	channelusecase "notify/internal/biz/usecase/channel"
	notifyenum "notify/internal/enum"

	"github.com/samber/lo"
)

type NotifyUsecase struct {
	log                *slog.Logger
	userAccountRepo    repo.UserAccountRepo
	templateUsecase    *TemplateUsecase
	stationUsecase     *channelusecase.StationUsecase
	emailUsecase       *channelusecase.EmailUsecase
	tencentSMSUsecase  *channelusecase.TencentSMSUsecase
	larkWebhookUsecase *channelusecase.LarkWebhookUsecase
}

func NewNotifyUsecase(
	logger *slog.Logger,
	userAccountRepo repo.UserAccountRepo,
	templateUsecase *TemplateUsecase,
	stationUsecase *channelusecase.StationUsecase,
	emailUsecase *channelusecase.EmailUsecase,
	tencentSMSUsecase *channelusecase.TencentSMSUsecase,
	larkWebhookUsecase *channelusecase.LarkWebhookUsecase,
) *NotifyUsecase {
	return &NotifyUsecase{
		log:                logger,
		userAccountRepo:    userAccountRepo,
		templateUsecase:    templateUsecase,
		stationUsecase:     stationUsecase,
		emailUsecase:       emailUsecase,
		tencentSMSUsecase:  tencentSMSUsecase,
		larkWebhookUsecase: larkWebhookUsecase,
	}
}

type NotifySendReq struct {
	EventID      string
	EventType    commonenum.EventType
	Language     notifyenum.Language
	Channels     []notifyenum.NotificationChannel
	TemplateData templatedata.NotificationTemplateData
	Recipients   []*model.NotificationRecipient
}

func (u *NotifyUsecase) Send(ctx context.Context, req *NotifySendReq) error {
	if req == nil || req.EventID == "" || len(req.Channels) == 0 {
		return nil
	}
	notificationContext := &model.NotificationContext{
		EventID:      req.EventID,
		EventType:    req.EventType,
		Language:     req.Language,
		TemplateData: req.TemplateData,
		Recipients:   req.Recipients,
	}
	rules, err := u.templateUsecase.ListEnabledRulesWithTemplates(ctx, &ListEnabledRulesWithTemplatesReq{
		EventType: req.EventType,
		Language:  req.Language,
		Channels:  req.Channels,
	})
	if err != nil || len(rules) == 0 {
		return err
	}
	needsContact := false
	for _, rule := range rules {
		if rule == nil || !rule.Enabled {
			continue
		}
		if rule.Channel == notifyenum.NotificationChannelEmail || rule.Channel == notifyenum.NotificationChannelTencentSMS {
			needsContact = true
			break
		}
	}
	accountsByUserID := map[int64]*model.UserAccount{}
	if needsContact && u.userAccountRepo != nil {
		userIDs := lo.Uniq(lo.FilterMap(notificationContext.Recipients, func(recipient *model.NotificationRecipient, _ int) (int64, bool) {
			return recipient.UserID, recipient != nil && recipient.UserID != 0 && (recipient.Email == "" || recipient.Phone == "")
		}))
		accountsByUserID, err = u.userAccountRepo.MapAccounts(ctx, userIDs)
	}
	if err != nil {
		return err
	}
	for _, rule := range rules {
		if rule == nil || !rule.Enabled {
			continue
		}
		status := notifyenum.NotificationChannelStatusSkipped
		switch rule.Channel {
		case notifyenum.NotificationChannelStation:
			status, err = u.stationUsecase.Process(ctx, &channelusecase.StationProcessReq{
				NotificationContext: notificationContext,
				Rule:                rule,
				AccountsByUserID:    accountsByUserID,
			})
		case notifyenum.NotificationChannelEmail:
			status, err = u.emailUsecase.Process(ctx, &channelusecase.EmailProcessReq{
				NotificationContext: notificationContext,
				Rule:                rule,
				AccountsByUserID:    accountsByUserID,
			})
		case notifyenum.NotificationChannelTencentSMS:
			status, err = u.tencentSMSUsecase.Process(ctx, &channelusecase.TencentSMSProcessReq{
				NotificationContext: notificationContext,
				Rule:                rule,
				AccountsByUserID:    accountsByUserID,
			})
		case notifyenum.NotificationChannelLarkWebhook:
			status, err = u.larkWebhookUsecase.Process(ctx, &channelusecase.LarkWebhookProcessReq{
				NotificationContext: notificationContext,
				Rule:                rule,
				AccountsByUserID:    accountsByUserID,
			})
		}
		if err != nil {
			return err
		}
		switch status {
		case notifyenum.NotificationChannelStatusProcessing, notifyenum.NotificationChannelStatusFailed, notifyenum.NotificationChannelStatusInternalError:
			return fmt.Errorf("notification channel not completed: event_id=%s channel=%s status=%s", notificationContext.EventID, rule.Channel, status)
		}
	}
	return nil
}
