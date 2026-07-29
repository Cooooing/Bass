package channel

import (
	"context"
	"fmt"
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
	if conf != nil && conf.Notify != nil && conf.Notify.NotificationRateLimit != nil {
		rateLimitEnabled = conf.Notify.NotificationRateLimit.Enable
		if conf.Notify.NotificationRateLimit.Window != nil && conf.Notify.NotificationRateLimit.Window.AsDuration() > 0 {
			rateLimitWindow = conf.Notify.NotificationRateLimit.Window.AsDuration()
		}
		if conf.Notify.NotificationRateLimit.MaxCount > 0 {
			rateLimitMaxCount = conf.Notify.NotificationRateLimit.MaxCount
		}
	}
	return &TencentSMSUsecase{
		tencentSMSDeliveryRepo:     tencentSMSDeliveryRepo,
		notificationRateLimitCache: notificationRateLimitCache,
		tencentSMSClient:           tencentSMSClient,
		smsEnabled:                 conf != nil && conf.Notify != nil && conf.Notify.Sms != nil && conf.Notify.Sms.Enable,
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
	notificationContext := req.NotificationContext
	rule := req.Rule
	accountsByUserID := req.AccountsByUserID
	if !u.smsEnabled {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if rule.TencentSMSTemplate == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if u.tencentSMSClient == nil {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("tencent sms client is nil")
	}
	status := notifyenum.NotificationChannelStatusSkipped
	seen := map[string]struct{}{}
	for _, recipient := range notificationContext.Recipients {
		if recipient == nil {
			continue
		}
		phone := strings.TrimSpace(recipient.Phone)
		if phone == "" && recipient.UserID != 0 {
			if account := accountsByUserID[recipient.UserID]; account != nil {
				phone = strings.TrimSpace(account.Phone)
			}
		}
		if phone == "" {
			continue
		}
		if _, ok := seen[phone]; ok {
			continue
		}
		seen[phone] = struct{}{}
		rendered, err := rule.TencentSMSTemplate.Render(ctx, notificationContext.TemplateData)
		if err != nil {
			continue
		}
		delivery := &model.NotificationTencentSMSDelivery{
			EventID:            notificationContext.EventID,
			EventType:          notificationContext.EventType,
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
		saveResp, err := u.tencentSMSDeliveryRepo.SaveOrGet(ctx, delivery)
		if err == nil {
			delivery = saveResp
		}
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		itemStatus := delivery.Status
		if itemStatus != notifyenum.NotificationChannelStatusSucceeded && itemStatus != notifyenum.NotificationChannelStatusUnknown && itemStatus != notifyenum.NotificationChannelStatusRateLimited {
			claimResp, err := u.tencentSMSDeliveryRepo.Claim(ctx, &repo.NotificationTencentSMSDeliveryClaimReq{
				ID:                delivery.ID,
				Now:               time.Now(),
				ProcessingTimeout: u.externalProcessingTimeout,
			})
			if err != nil {
				return notifyenum.NotificationChannelStatusInternalError, err
			}
			if claimResp {
				if u.rateLimitEnabled {
					allowResp, err := u.notificationRateLimitCache.Allow(ctx, &repo.NotificationRateLimitSpec{
						Channel:   notifyenum.NotificationChannelTencentSMS,
						Recipient: delivery.Phone,
						Window:    u.rateLimitWindow,
						MaxCount:  u.rateLimitMaxCount,
					})
					if err != nil {
						return notifyenum.NotificationChannelStatusInternalError, err
					}
					if !allowResp {
						if err := u.tencentSMSDeliveryRepo.UpdateStatus(ctx, &repo.NotificationTencentSMSDeliveryUpdateStatusReq{
							ID:     delivery.ID,
							Status: notifyenum.NotificationChannelStatusRateLimited,
						}); err != nil {
							return notifyenum.NotificationChannelStatusInternalError, err
						}
						itemStatus = notifyenum.NotificationChannelStatusRateLimited
					}
				}
				if itemStatus != notifyenum.NotificationChannelStatusRateLimited {
					result, err := u.tencentSMSClient.SendTencentSMS(ctx, &repo.TencentSMSRequest{
						IdempotencyKey:     fmt.Sprintf("%d", delivery.ID),
						Phone:              delivery.Phone,
						SMSSDKAppID:        delivery.SMSSDKAppID,
						SignName:           delivery.SignName,
						ProviderTemplateID: delivery.ProviderTemplateID,
						TemplateParams:     delivery.TemplateParams,
					})
					if err != nil {
						return notifyenum.NotificationChannelStatusInternalError, err
					}
					if result == nil {
						itemStatus = notifyenum.NotificationChannelStatusUnknown
						err = u.tencentSMSDeliveryRepo.UpdateStatus(ctx, &repo.NotificationTencentSMSDeliveryUpdateStatusReq{
							ID:     delivery.ID,
							Status: itemStatus,
						})
					} else {
						itemStatus = result.Status
						if itemStatus != notifyenum.NotificationChannelStatusSucceeded && itemStatus != notifyenum.NotificationChannelStatusFailed {
							itemStatus = notifyenum.NotificationChannelStatusUnknown
						}
						var sentAt *time.Time
						if itemStatus == notifyenum.NotificationChannelStatusSucceeded {
							sentAt = new(time.Now())
						}
						err = u.tencentSMSDeliveryRepo.UpdateStatus(ctx, &repo.NotificationTencentSMSDeliveryUpdateStatusReq{
							ID:                delivery.ID,
							Status:            itemStatus,
							ProviderRequestID: result.ProviderRequestID,
							ProviderCode:      result.ProviderCode,
							ProviderMessage:   result.ProviderMessage,
							SentAt:            sentAt,
						})
					}
					if err != nil {
						return itemStatus, err
					}
				}
			}
		}
		status = status.Merge(itemStatus)
		if status.Blocking() {
			return status, nil
		}
	}
	return status, nil
}
