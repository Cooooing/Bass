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

type EmailUsecase struct {
	emailDeliveryRepo          repo.NotificationEmailDeliveryRepo
	notificationRateLimitCache repo.NotificationRateLimitCache
	emailClient                repo.EmailClient
	rateLimitEnabled           bool
	rateLimitWindow            time.Duration
	rateLimitMaxCount          int64
	externalProcessingTimeout  time.Duration
}

func NewEmailUsecase(
	conf *config.Bootstrap,
	emailDeliveryRepo repo.NotificationEmailDeliveryRepo,
	notificationRateLimitCache repo.NotificationRateLimitCache,
	emailClient repo.EmailClient,
) *EmailUsecase {
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
	return &EmailUsecase{
		emailDeliveryRepo:          emailDeliveryRepo,
		notificationRateLimitCache: notificationRateLimitCache,
		emailClient:                emailClient,
		rateLimitEnabled:           rateLimitEnabled,
		rateLimitWindow:            rateLimitWindow,
		rateLimitMaxCount:          rateLimitMaxCount,
		externalProcessingTimeout:  10 * time.Minute,
	}
}

type EmailProcessReq struct {
	NotificationContext *model.NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

func (u *EmailUsecase) Process(ctx context.Context, req *EmailProcessReq) (notifyenum.NotificationChannelStatus, error) {
	if req == nil || req.NotificationContext == nil || req.Rule == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	notificationContext := req.NotificationContext
	rule := req.Rule
	accountsByUserID := req.AccountsByUserID
	if rule.EmailTemplate == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if u.emailClient == nil {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("email client is nil")
	}
	status := notifyenum.NotificationChannelStatusSkipped
	seen := map[string]struct{}{}
	for _, recipient := range notificationContext.Recipients {
		if recipient == nil {
			continue
		}
		toEmail := strings.TrimSpace(recipient.Email)
		if toEmail == "" && recipient.UserID != 0 {
			if account := accountsByUserID[recipient.UserID]; account != nil {
				toEmail = strings.TrimSpace(account.Email)
			}
		}
		if toEmail == "" {
			continue
		}
		if _, ok := seen[toEmail]; ok {
			continue
		}
		seen[toEmail] = struct{}{}
		rendered, err := rule.EmailTemplate.Render(ctx, notificationContext.TemplateData)
		if err != nil {
			continue
		}
		delivery := &model.NotificationEmailDelivery{
			EventID:     notificationContext.EventID,
			EventType:   notificationContext.EventType,
			ToEmail:     toEmail,
			Subject:     rendered.Subject,
			Body:        rendered.Body,
			ContentType: rendered.ContentType,
			Status:      notifyenum.NotificationChannelStatusProcessing,
		}
		if recipient.UserID != 0 {
			delivery.ReceiverID = new(recipient.UserID)
		}
		saveResp, err := u.emailDeliveryRepo.SaveOrGet(ctx, delivery)
		if err == nil {
			delivery = saveResp
		}
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		itemStatus := delivery.Status
		if itemStatus != notifyenum.NotificationChannelStatusSucceeded && itemStatus != notifyenum.NotificationChannelStatusUnknown && itemStatus != notifyenum.NotificationChannelStatusRateLimited {
			claimResp, err := u.emailDeliveryRepo.Claim(ctx, &repo.NotificationEmailDeliveryClaimReq{
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
						Channel:   notifyenum.NotificationChannelEmail,
						Recipient: delivery.ToEmail,
						Window:    u.rateLimitWindow,
						MaxCount:  u.rateLimitMaxCount,
					})
					if err != nil {
						return notifyenum.NotificationChannelStatusInternalError, err
					}
					if !allowResp {
						if err := u.emailDeliveryRepo.UpdateStatus(ctx, &repo.NotificationEmailDeliveryUpdateStatusReq{
							ID:     delivery.ID,
							Status: notifyenum.NotificationChannelStatusRateLimited,
						}); err != nil {
							return notifyenum.NotificationChannelStatusInternalError, err
						}
						itemStatus = notifyenum.NotificationChannelStatusRateLimited
					}
				}
				if itemStatus != notifyenum.NotificationChannelStatusRateLimited {
					result, err := u.emailClient.SendEmail(ctx, &repo.EmailRequest{
						IdempotencyKey: fmt.Sprintf("%d", delivery.ID),
						ToEmail:        delivery.ToEmail,
						Subject:        delivery.Subject,
						Body:           delivery.Body,
						ContentType:    delivery.ContentType,
					})
					if err != nil {
						return notifyenum.NotificationChannelStatusInternalError, err
					}
					if result == nil {
						itemStatus = notifyenum.NotificationChannelStatusUnknown
						err = u.emailDeliveryRepo.UpdateStatus(ctx, &repo.NotificationEmailDeliveryUpdateStatusReq{
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
							now := time.Now()
							sentAt = &now
						}
						err = u.emailDeliveryRepo.UpdateStatus(ctx, &repo.NotificationEmailDeliveryUpdateStatusReq{
							ID:                delivery.ID,
							Status:            itemStatus,
							ProviderMessageID: result.ProviderMessageID,
							ProviderResp:      result.ProviderResp,
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
