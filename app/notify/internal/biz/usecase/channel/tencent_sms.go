package channel

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/config"
	notifyenum "notify/internal/enum"
)

type TencentSMSUsecase struct {
	tencentSMSDeliveryRepo     repo.NotificationTencentSMSDeliveryRepo
	notificationRateLimitCache repo.NotificationRateLimitCache
	tencentSMSClient           repo.TencentSMSClient
	smsEnabled                 bool
	rateLimitEnabled           bool
	rateLimitWindow            time.Duration
	rateLimitMaxCount          int64
	externalProcessingTimeout  time.Duration
}

func NewTencentSMSUsecase(
	conf *config.Bootstrap,
	tencentSMSDeliveryRepo repo.NotificationTencentSMSDeliveryRepo,
	notificationRateLimitCache repo.NotificationRateLimitCache,
	tencentSMSClient repo.TencentSMSClient,
) *TencentSMSUsecase {
	rateLimitEnabled := true
	rateLimitWindow := 5 * time.Minute
	rateLimitMaxCount := int64(5)
	if conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetNotificationRateLimit() != nil {
		rateLimitEnabled = conf.GetBusiness().GetNotificationRateLimit().GetEnable()
		if conf.GetBusiness().GetNotificationRateLimit().GetWindow() != nil && conf.GetBusiness().GetNotificationRateLimit().GetWindow().AsDuration() > 0 {
			rateLimitWindow = conf.GetBusiness().GetNotificationRateLimit().GetWindow().AsDuration()
		}
		if conf.GetBusiness().GetNotificationRateLimit().GetMaxCount() > 0 {
			rateLimitMaxCount = conf.GetBusiness().GetNotificationRateLimit().GetMaxCount()
		}
	}
	return &TencentSMSUsecase{
		tencentSMSDeliveryRepo:     tencentSMSDeliveryRepo,
		notificationRateLimitCache: notificationRateLimitCache,
		tencentSMSClient:           tencentSMSClient,
		smsEnabled:                 conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetSms() != nil && conf.GetBusiness().GetSms().GetEnable(),
		rateLimitEnabled:           rateLimitEnabled,
		rateLimitWindow:            rateLimitWindow,
		rateLimitMaxCount:          rateLimitMaxCount,
		externalProcessingTimeout:  10 * time.Minute,
	}
}

type TencentSMSProcessReq struct {
	NotificationContext *model.NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

func (u *TencentSMSUsecase) Process(ctx context.Context, req *TencentSMSProcessReq) (notifyenum.NotificationChannelStatus, error) {
	if req == nil || req.NotificationContext == nil || req.Rule == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if !u.smsEnabled {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if req.Rule.TencentSMSTemplate == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if u.tencentSMSClient == nil {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("tencent sms client is nil")
	}
	status := notifyenum.NotificationChannelStatusSkipped
	seen := map[string]struct{}{}
	for _, recipient := range req.NotificationContext.Recipients {
		phone := u.resolveRecipientPhone(recipient, req.AccountsByUserID)
		if phone == "" {
			continue
		}
		if _, ok := seen[phone]; ok {
			continue
		}
		seen[phone] = struct{}{}
		itemStatus, err := u.processRecipient(ctx, req, recipient, phone)
		if err != nil {
			return itemStatus, err
		}
		status = status.Merge(itemStatus)
		if status.Blocking() {
			return status, nil
		}
	}
	return status, nil
}

func (*TencentSMSUsecase) resolveRecipientPhone(recipient *model.NotificationRecipient, accountsByUserID map[int64]*model.UserAccount) string {
	if recipient == nil {
		return ""
	}
	phone := strings.TrimSpace(recipient.Phone)
	if phone == "" && recipient.UserID != 0 {
		if account := accountsByUserID[recipient.UserID]; account != nil {
			phone = strings.TrimSpace(account.Phone)
		}
	}
	return phone
}

func (u *TencentSMSUsecase) processRecipient(ctx context.Context, req *TencentSMSProcessReq, recipient *model.NotificationRecipient, phone string) (notifyenum.NotificationChannelStatus, error) {
	rendered, renderedOK := u.renderDelivery(ctx, req)
	if !renderedOK {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	delivery := &model.NotificationTencentSMSDelivery{
		EventID:            req.NotificationContext.EventID,
		NotificationType:   req.NotificationContext.NotificationType,
		Phone:              phone,
		SMSSDKAppID:        rendered.SMSSDKAppID,
		SignName:           rendered.SignName,
		ProviderTemplateID: rendered.ProviderTemplateID,
		TemplateParams:     rendered.TemplateParams,
		Status:             notifyenum.NotificationChannelStatusProcessing,
	}
	if recipient.UserID != 0 {
		delivery.ReceiverID = new(recipient.UserID)
	}
	var err error
	delivery, err = u.tencentSMSDeliveryRepo.SaveOrGet(ctx, delivery)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	return u.claimAndSend(ctx, delivery)
}

func (*TencentSMSUsecase) renderDelivery(ctx context.Context, req *TencentSMSProcessReq) (*model.NotificationTencentSMSRendered, bool) {
	rendered, err := req.Rule.TencentSMSTemplate.Render(ctx, req.NotificationContext.TemplateData)
	return rendered, err == nil
}

func (u *TencentSMSUsecase) claimAndSend(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (notifyenum.NotificationChannelStatus, error) {
	status := delivery.Status
	if status == notifyenum.NotificationChannelStatusSucceeded || status == notifyenum.NotificationChannelStatusUnknown || status == notifyenum.NotificationChannelStatusRateLimited {
		return status, nil
	}
	claimed, err := u.tencentSMSDeliveryRepo.Claim(ctx, &repo.NotificationTencentSMSDeliveryClaimReq{ID: delivery.ID, Now: time.Now(), ProcessingTimeout: u.externalProcessingTimeout})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if !claimed {
		return status, nil
	}
	if u.rateLimitEnabled {
		allowed, err := u.notificationRateLimitCache.Allow(ctx, &repo.NotificationRateLimitSpec{Channel: notifyenum.NotificationChannelTencentSMS, Recipient: delivery.Phone, Window: u.rateLimitWindow, MaxCount: u.rateLimitMaxCount})
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		if !allowed {
			if err := u.tencentSMSDeliveryRepo.UpdateStatus(ctx, &repo.NotificationTencentSMSDeliveryUpdateStatusReq{ID: delivery.ID, Status: notifyenum.NotificationChannelStatusRateLimited}); err != nil {
				return notifyenum.NotificationChannelStatusInternalError, err
			}
			return notifyenum.NotificationChannelStatusRateLimited, nil
		}
	}
	return u.sendTencentSMSDelivery(ctx, delivery)
}

func (u *TencentSMSUsecase) sendTencentSMSDelivery(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (notifyenum.NotificationChannelStatus, error) {
	result, err := u.tencentSMSClient.SendTencentSMS(ctx, &repo.TencentSMSRequest{IdempotencyKey: strconv.FormatInt(delivery.ID, 10), Phone: delivery.Phone, SMSSDKAppID: delivery.SMSSDKAppID, SignName: delivery.SignName, ProviderTemplateID: delivery.ProviderTemplateID, TemplateParams: delivery.TemplateParams})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if result == nil {
		status := notifyenum.NotificationChannelStatusUnknown
		return status, u.tencentSMSDeliveryRepo.UpdateStatus(ctx, &repo.NotificationTencentSMSDeliveryUpdateStatusReq{ID: delivery.ID, Status: status})
	}
	status := result.Status
	if status != notifyenum.NotificationChannelStatusSucceeded && status != notifyenum.NotificationChannelStatusFailed {
		status = notifyenum.NotificationChannelStatusUnknown
	}
	var sentAt *time.Time
	if status == notifyenum.NotificationChannelStatusSucceeded {
		sentAt = new(time.Now())
	}
	err = u.tencentSMSDeliveryRepo.UpdateStatus(ctx, &repo.NotificationTencentSMSDeliveryUpdateStatusReq{ID: delivery.ID, Status: status, ProviderRequestID: result.ProviderRequestID, ProviderCode: result.ProviderCode, ProviderMessage: result.ProviderMessage, SentAt: sentAt})
	return status, err
}
