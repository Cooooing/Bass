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
	if conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetNotificationRateLimit() != nil {
		rateLimitEnabled = conf.GetBusiness().GetNotificationRateLimit().GetEnable()
		if conf.GetBusiness().GetNotificationRateLimit().GetWindow() != nil && conf.GetBusiness().GetNotificationRateLimit().GetWindow().AsDuration() > 0 {
			rateLimitWindow = conf.GetBusiness().GetNotificationRateLimit().GetWindow().AsDuration()
		}
		if conf.GetBusiness().GetNotificationRateLimit().GetMaxCount() > 0 {
			rateLimitMaxCount = conf.GetBusiness().GetNotificationRateLimit().GetMaxCount()
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
	if req.Rule.EmailTemplate == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if u.emailClient == nil {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("email client is nil")
	}
	status := notifyenum.NotificationChannelStatusSkipped
	seen := map[string]struct{}{}
	for _, recipient := range req.NotificationContext.Recipients {
		toEmail := u.resolveRecipientEmail(recipient, req.AccountsByUserID)
		if toEmail == "" {
			continue
		}
		if _, ok := seen[toEmail]; ok {
			continue
		}
		seen[toEmail] = struct{}{}
		itemStatus, err := u.processRecipient(ctx, req, recipient, toEmail)
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

func (*EmailUsecase) resolveRecipientEmail(recipient *model.NotificationRecipient, accountsByUserID map[int64]*model.UserAccount) string {
	if recipient == nil {
		return ""
	}
	toEmail := strings.TrimSpace(recipient.Email)
	if toEmail == "" && recipient.UserID != 0 {
		if account := accountsByUserID[recipient.UserID]; account != nil {
			toEmail = strings.TrimSpace(account.Email)
		}
	}
	return toEmail
}

func (u *EmailUsecase) processRecipient(ctx context.Context, req *EmailProcessReq, recipient *model.NotificationRecipient, toEmail string) (notifyenum.NotificationChannelStatus, error) {
	rendered, renderedOK := u.renderDelivery(ctx, req)
	if !renderedOK {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	delivery := &model.NotificationEmailDelivery{
		EventID:          req.NotificationContext.EventID,
		NotificationType: req.NotificationContext.NotificationType,
		ToEmail:          toEmail,
		Subject:          rendered.Subject,
		Body:             rendered.Body,
		ContentType:      rendered.ContentType,
		Status:           notifyenum.NotificationChannelStatusProcessing,
	}
	if recipient.UserID != 0 {
		delivery.ReceiverID = new(recipient.UserID)
	}
	var err error
	delivery, err = u.emailDeliveryRepo.SaveOrGet(ctx, delivery)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	return u.claimAndSend(ctx, delivery)
}

func (*EmailUsecase) renderDelivery(ctx context.Context, req *EmailProcessReq) (*model.NotificationEmailRendered, bool) {
	rendered, err := req.Rule.EmailTemplate.Render(ctx, req.NotificationContext.TemplateData)
	return rendered, err == nil
}

func (u *EmailUsecase) claimAndSend(ctx context.Context, delivery *model.NotificationEmailDelivery) (notifyenum.NotificationChannelStatus, error) {
	status := delivery.Status
	if status == notifyenum.NotificationChannelStatusSucceeded || status == notifyenum.NotificationChannelStatusUnknown || status == notifyenum.NotificationChannelStatusRateLimited {
		return status, nil
	}
	claimed, err := u.emailDeliveryRepo.Claim(ctx, &repo.NotificationEmailDeliveryClaimReq{ID: delivery.ID, Now: time.Now(), ProcessingTimeout: u.externalProcessingTimeout})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if !claimed {
		return status, nil
	}
	if u.rateLimitEnabled {
		allowed, err := u.notificationRateLimitCache.Allow(ctx, &repo.NotificationRateLimitSpec{Channel: notifyenum.NotificationChannelEmail, Recipient: delivery.ToEmail, Window: u.rateLimitWindow, MaxCount: u.rateLimitMaxCount})
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		if !allowed {
			if err := u.emailDeliveryRepo.UpdateStatus(ctx, &repo.NotificationEmailDeliveryUpdateStatusReq{ID: delivery.ID, Status: notifyenum.NotificationChannelStatusRateLimited}); err != nil {
				return notifyenum.NotificationChannelStatusInternalError, err
			}
			return notifyenum.NotificationChannelStatusRateLimited, nil
		}
	}
	return u.sendEmailDelivery(ctx, delivery)
}

func (u *EmailUsecase) sendEmailDelivery(ctx context.Context, delivery *model.NotificationEmailDelivery) (notifyenum.NotificationChannelStatus, error) {
	result, err := u.emailClient.SendEmail(ctx, &repo.EmailRequest{IdempotencyKey: strconv.FormatInt(delivery.ID, 10), ToEmail: delivery.ToEmail, Subject: delivery.Subject, Body: delivery.Body, ContentType: delivery.ContentType})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if result == nil {
		status := notifyenum.NotificationChannelStatusUnknown
		return status, u.emailDeliveryRepo.UpdateStatus(ctx, &repo.NotificationEmailDeliveryUpdateStatusReq{ID: delivery.ID, Status: status})
	}
	status := result.Status
	if status != notifyenum.NotificationChannelStatusSucceeded && status != notifyenum.NotificationChannelStatusFailed {
		status = notifyenum.NotificationChannelStatusUnknown
	}
	var sentAt *time.Time
	if status == notifyenum.NotificationChannelStatusSucceeded {
		sentAt = new(time.Now())
	}
	err = u.emailDeliveryRepo.UpdateStatus(ctx, &repo.NotificationEmailDeliveryUpdateStatusReq{ID: delivery.ID, Status: status, ProviderMessageID: result.ProviderMessageID, ProviderResp: result.ProviderResp, SentAt: sentAt})
	return status, err
}
