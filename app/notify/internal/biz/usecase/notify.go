package usecase

import (
	"bytes"
	commonenum "common/pkg/enum"
	"context"
	"encoding/json"
	"fmt"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	notifyenum "notify/internal/enum"
	"strings"
	"text/template"
	"time"

	"log/slog"
)

type NotificationRecipient struct {
	UserID int64
	Email  string
	Phone  string
}

type NotificationContext struct {
	EventID      string
	EventType    commonenum.EventType
	Language     notifyenum.Language
	TemplateData any
	Recipients   []*NotificationRecipient
}

type NotifyUsecase struct {
	log                        *slog.Logger
	userClient                 repo.UserClient
	notificationRuleRepo       repo.NotificationRuleRepo
	stationMessageRepo         repo.NotificationStationMessageRepo
	emailDeliveryRepo          repo.NotificationEmailDeliveryRepo
	tencentSMSDeliveryRepo     repo.NotificationTencentSMSDeliveryRepo
	larkWebhookDeliveryRepo    repo.NotificationLarkWebhookDeliveryRepo
	notificationRateLimitCache repo.NotificationRateLimitCache
	emailClient                bizchannel.EmailClient
	tencentSMSClient           bizchannel.TencentSMSClient
	larkWebhookClient          bizchannel.LarkWebhookClient
	smsEnabled                 bool
	rateLimitEnabled           bool
	rateLimitWindow            time.Duration
	rateLimitMaxCount          int64
	externalProcessingTimeout  time.Duration
}

func NewNotifyUsecase(
	logger *slog.Logger,
	conf *conf.Bootstrap,
	userClient repo.UserClient,
	notificationRuleRepo repo.NotificationRuleRepo,
	stationMessageRepo repo.NotificationStationMessageRepo,
	emailDeliveryRepo repo.NotificationEmailDeliveryRepo,
	tencentSMSDeliveryRepo repo.NotificationTencentSMSDeliveryRepo,
	larkWebhookDeliveryRepo repo.NotificationLarkWebhookDeliveryRepo,
	notificationRateLimitCache repo.NotificationRateLimitCache,
	emailClient bizchannel.EmailClient,
	tencentSMSClient bizchannel.TencentSMSClient,
	larkWebhookClient bizchannel.LarkWebhookClient,
) *NotifyUsecase {
	rateLimitEnabled := true
	rateLimitWindow := 5 * time.Minute
	rateLimitMaxCount := int64(5)
	if conf != nil && conf.Server != nil && conf.Server.NotificationRateLimit != nil {
		rateLimitEnabled = conf.Server.NotificationRateLimit.Enable
		if conf.Server.NotificationRateLimit.Window != nil && conf.Server.NotificationRateLimit.Window.AsDuration() > 0 {
			rateLimitWindow = conf.Server.NotificationRateLimit.Window.AsDuration()
		}
		if conf.Server.NotificationRateLimit.MaxCount > 0 {
			rateLimitMaxCount = conf.Server.NotificationRateLimit.MaxCount
		}
	}
	return &NotifyUsecase{
		log:                        logger,
		userClient:                 userClient,
		notificationRuleRepo:       notificationRuleRepo,
		stationMessageRepo:         stationMessageRepo,
		emailDeliveryRepo:          emailDeliveryRepo,
		tencentSMSDeliveryRepo:     tencentSMSDeliveryRepo,
		larkWebhookDeliveryRepo:    larkWebhookDeliveryRepo,
		notificationRateLimitCache: notificationRateLimitCache,
		emailClient:                emailClient,
		tencentSMSClient:           tencentSMSClient,
		larkWebhookClient:          larkWebhookClient,
		smsEnabled:                 conf != nil && conf.Server != nil && conf.Server.Sms != nil && conf.Server.Sms.Enable,
		rateLimitEnabled:           rateLimitEnabled,
		rateLimitWindow:            rateLimitWindow,
		rateLimitMaxCount:          rateLimitMaxCount,
		externalProcessingTimeout:  10 * time.Minute,
	}
}

func (u *NotifyUsecase) ListEnabledRules(ctx context.Context, eventType commonenum.EventType, language notifyenum.Language) ([]*model.NotificationRule, error) {
	enabled := true
	return u.notificationRuleRepo.List(ctx, &repo.NotificationRuleQuery{
		EventType: &eventType,
		Language:  &language,
		Enabled:   &enabled,
	})
}

func (u *NotifyUsecase) Process(ctx context.Context, notificationContext *NotificationContext, rules []*model.NotificationRule) error {
	if notificationContext == nil || notificationContext.EventID == "" {
		return nil
	}
	accountsByUserID, err := u.loadAccounts(ctx, notificationContext, rules)
	if err != nil {
		return err
	}
	for _, rule := range rules {
		status, err := u.processRule(ctx, notificationContext, rule, accountsByUserID)
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

func (u *NotifyUsecase) processRule(ctx context.Context, notificationContext *NotificationContext, rule *model.NotificationRule, accountsByUserID map[int64]*model.UserAccount) (notifyenum.NotificationChannelStatus, error) {
	if rule == nil || !rule.Enabled {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	switch rule.Channel {
	case notifyenum.NotificationChannelStation:
		return u.processStation(ctx, notificationContext, rule)
	case notifyenum.NotificationChannelEmail:
		return u.processEmail(ctx, notificationContext, rule, accountsByUserID)
	case notifyenum.NotificationChannelTencentSMS:
		return u.processTencentSMS(ctx, notificationContext, rule, accountsByUserID)
	case notifyenum.NotificationChannelLarkWebhook:
		return u.processLarkWebhook(ctx, notificationContext, rule)
	default:
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
}

func (u *NotifyUsecase) processStation(ctx context.Context, notificationContext *NotificationContext, rule *model.NotificationRule) (notifyenum.NotificationChannelStatus, error) {
	if rule.StationTemplate == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	seen := make(map[int64]struct{}, len(notificationContext.Recipients))
	written := false
	for _, recipient := range notificationContext.Recipients {
		if recipient == nil || recipient.UserID == 0 {
			continue
		}
		if _, ok := seen[recipient.UserID]; ok {
			continue
		}
		seen[recipient.UserID] = struct{}{}
		title, ok := u.renderTemplate(rule.StationTemplate.TitleTemplate, notificationContext.TemplateData)
		if !ok {
			continue
		}
		content, ok := u.renderTemplate(rule.StationTemplate.ContentTemplate, notificationContext.TemplateData)
		if !ok {
			continue
		}
		_, err := u.stationMessageRepo.Save(ctx, &model.NotificationStationMessage{
			EventID:    notificationContext.EventID,
			EventType:  notificationContext.EventType,
			ReceiverID: recipient.UserID,
			Title:      title,
			Content:    content,
			Status:     notifyenum.NotificationChannelStatusSucceeded,
		})
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		written = true
	}
	if !written {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	return notifyenum.NotificationChannelStatusSucceeded, nil
}

func (u *NotifyUsecase) processEmail(ctx context.Context, notificationContext *NotificationContext, rule *model.NotificationRule, accountsByUserID map[int64]*model.UserAccount) (notifyenum.NotificationChannelStatus, error) {
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
		subject, ok := u.renderTemplate(rule.EmailTemplate.SubjectTemplate, notificationContext.TemplateData)
		if !ok {
			continue
		}
		body, ok := u.renderTemplate(rule.EmailTemplate.BodyTemplate, notificationContext.TemplateData)
		if !ok {
			continue
		}
		delivery := &model.NotificationEmailDelivery{
			EventID:     notificationContext.EventID,
			EventType:   notificationContext.EventType,
			ToEmail:     toEmail,
			Subject:     subject,
			Body:        body,
			ContentType: rule.EmailTemplate.ContentType,
			Status:      notifyenum.NotificationChannelStatusProcessing,
		}
		if recipient.UserID != 0 {
			delivery.ReceiverID = new(recipient.UserID)
		}
		itemStatus, err := u.sendEmail(ctx, delivery)
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

func (u *NotifyUsecase) sendEmail(ctx context.Context, delivery *model.NotificationEmailDelivery) (notifyenum.NotificationChannelStatus, error) {
	delivery, err := u.emailDeliveryRepo.SaveOrGet(ctx, delivery)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if delivery.Status == notifyenum.NotificationChannelStatusSucceeded {
		return notifyenum.NotificationChannelStatusSucceeded, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusUnknown {
		return notifyenum.NotificationChannelStatusUnknown, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusRateLimited {
		return notifyenum.NotificationChannelStatusRateLimited, nil
	}
	claimed, err := u.emailDeliveryRepo.Claim(ctx, delivery.ID, time.Now(), u.externalProcessingTimeout, false)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if !claimed {
		return delivery.Status, nil
	}
	if u.rateLimitEnabled {
		allowed, err := u.notificationRateLimitCache.Allow(ctx, &repo.NotificationRateLimitSpec{
			Channel:   notifyenum.NotificationChannelEmail,
			Recipient: delivery.ToEmail,
			Window:    u.rateLimitWindow,
			MaxCount:  u.rateLimitMaxCount,
		})
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		if !allowed {
			if err := u.emailDeliveryRepo.MarkRateLimited(ctx, delivery.ID); err != nil {
				return notifyenum.NotificationChannelStatusInternalError, err
			}
			return notifyenum.NotificationChannelStatusRateLimited, nil
		}
	}
	result, err := u.emailClient.SendEmail(ctx, &bizchannel.EmailRequest{
		IdempotencyKey: fmt.Sprintf("%d", delivery.ID),
		ToEmail:        delivery.ToEmail,
		Subject:        delivery.Subject,
		Body:           delivery.Body,
		ContentType:    delivery.ContentType,
	})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	return u.finishEmail(ctx, delivery.ID, result)
}

func (u *NotifyUsecase) finishEmail(ctx context.Context, deliveryID int64, result *bizchannel.SendResult) (notifyenum.NotificationChannelStatus, error) {
	if result == nil {
		return notifyenum.NotificationChannelStatusUnknown, u.emailDeliveryRepo.MarkUnknown(ctx, deliveryID, nil)
	}
	switch result.Status {
	case notifyenum.NotificationChannelStatusSucceeded:
		err := u.emailDeliveryRepo.MarkSucceeded(ctx, deliveryID, result.ProviderMessageID, result.ProviderResponse, time.Now())
		return notifyenum.NotificationChannelStatusSucceeded, err
	case notifyenum.NotificationChannelStatusFailed:
		err := u.emailDeliveryRepo.MarkFailed(ctx, deliveryID, result.ProviderResponse)
		return notifyenum.NotificationChannelStatusFailed, err
	default:
		err := u.emailDeliveryRepo.MarkUnknown(ctx, deliveryID, result.ProviderResponse)
		return notifyenum.NotificationChannelStatusUnknown, err
	}
}

func (u *NotifyUsecase) processTencentSMS(ctx context.Context, notificationContext *NotificationContext, rule *model.NotificationRule, accountsByUserID map[int64]*model.UserAccount) (notifyenum.NotificationChannelStatus, error) {
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
		params := make([]string, 0, len(rule.TencentSMSTemplate.ParamTemplates))
		renderFailed := false
		for _, paramTemplate := range rule.TencentSMSTemplate.ParamTemplates {
			param, ok := u.renderTemplate(paramTemplate, notificationContext.TemplateData)
			if !ok {
				renderFailed = true
				break
			}
			params = append(params, param)
		}
		if renderFailed {
			continue
		}
		delivery := &model.NotificationTencentSMSDelivery{
			EventID:            notificationContext.EventID,
			EventType:          notificationContext.EventType,
			Phone:              phone,
			SMSSDKAppID:        rule.TencentSMSTemplate.SMSSDKAppID,
			SignName:           rule.TencentSMSTemplate.SignName,
			ProviderTemplateID: rule.TencentSMSTemplate.ProviderTemplateID,
			TemplateParams:     params,
			Status:             notifyenum.NotificationChannelStatusProcessing,
		}
		if recipient.UserID != 0 {
			delivery.ReceiverID = new(recipient.UserID)
		}
		itemStatus, err := u.sendTencentSMS(ctx, delivery)
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

func (u *NotifyUsecase) sendTencentSMS(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (notifyenum.NotificationChannelStatus, error) {
	delivery, err := u.tencentSMSDeliveryRepo.SaveOrGet(ctx, delivery)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if delivery.Status == notifyenum.NotificationChannelStatusSucceeded {
		return notifyenum.NotificationChannelStatusSucceeded, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusUnknown {
		return notifyenum.NotificationChannelStatusUnknown, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusRateLimited {
		return notifyenum.NotificationChannelStatusRateLimited, nil
	}
	claimed, err := u.tencentSMSDeliveryRepo.Claim(ctx, delivery.ID, time.Now(), u.externalProcessingTimeout, false)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if !claimed {
		return delivery.Status, nil
	}
	if u.rateLimitEnabled {
		allowed, err := u.notificationRateLimitCache.Allow(ctx, &repo.NotificationRateLimitSpec{
			Channel:   notifyenum.NotificationChannelTencentSMS,
			Recipient: delivery.Phone,
			Window:    u.rateLimitWindow,
			MaxCount:  u.rateLimitMaxCount,
		})
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		if !allowed {
			if err := u.tencentSMSDeliveryRepo.MarkRateLimited(ctx, delivery.ID); err != nil {
				return notifyenum.NotificationChannelStatusInternalError, err
			}
			return notifyenum.NotificationChannelStatusRateLimited, nil
		}
	}
	result, err := u.tencentSMSClient.SendTencentSMS(ctx, &bizchannel.TencentSMSRequest{
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
	return u.finishTencentSMS(ctx, delivery.ID, result)
}

func (u *NotifyUsecase) finishTencentSMS(ctx context.Context, deliveryID int64, result *bizchannel.SendResult) (notifyenum.NotificationChannelStatus, error) {
	if result == nil {
		return notifyenum.NotificationChannelStatusUnknown, u.tencentSMSDeliveryRepo.MarkUnknown(ctx, deliveryID, nil, nil, nil)
	}
	switch result.Status {
	case notifyenum.NotificationChannelStatusSucceeded:
		err := u.tencentSMSDeliveryRepo.MarkSucceeded(ctx, deliveryID, result.ProviderRequestID, result.ProviderCode, result.ProviderMessage, time.Now())
		return notifyenum.NotificationChannelStatusSucceeded, err
	case notifyenum.NotificationChannelStatusFailed:
		err := u.tencentSMSDeliveryRepo.MarkFailed(ctx, deliveryID, result.ProviderRequestID, result.ProviderCode, result.ProviderMessage)
		return notifyenum.NotificationChannelStatusFailed, err
	default:
		err := u.tencentSMSDeliveryRepo.MarkUnknown(ctx, deliveryID, result.ProviderRequestID, result.ProviderCode, result.ProviderMessage)
		return notifyenum.NotificationChannelStatusUnknown, err
	}
}

func (u *NotifyUsecase) processLarkWebhook(ctx context.Context, notificationContext *NotificationContext, rule *model.NotificationRule) (notifyenum.NotificationChannelStatus, error) {
	if rule.LarkWebhookTemplate == nil || rule.LarkWebhookTemplate.WebhookID == "" || rule.LarkWebhookTemplate.Token == "" {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if u.larkWebhookClient == nil {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("lark webhook client is nil")
	}
	renderedContent, ok := u.renderTemplate(rule.LarkWebhookTemplate.ContentTemplate, notificationContext.TemplateData)
	if !ok {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	var content map[string]any
	if err := json.Unmarshal([]byte(renderedContent), &content); err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if content == nil {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("lark webhook content must be json object")
	}
	msgType := strings.TrimSpace(rule.LarkWebhookTemplate.MsgType)
	if msgType == "" {
		msgType = "text"
	}
	requestBodyBytes, err := json.Marshal(struct {
		MsgType string         `json:"msg_type"`
		Content map[string]any `json:"content"`
	}{
		MsgType: msgType,
		Content: content,
	})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	delivery := &model.NotificationLarkWebhookDelivery{
		EventID:     notificationContext.EventID,
		EventType:   notificationContext.EventType,
		WebhookID:   rule.LarkWebhookTemplate.WebhookID,
		RequestBody: string(requestBodyBytes),
		Status:      notifyenum.NotificationChannelStatusProcessing,
	}
	delivery, err = u.larkWebhookDeliveryRepo.SaveOrGet(ctx, delivery)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if delivery.Status == notifyenum.NotificationChannelStatusSucceeded {
		return notifyenum.NotificationChannelStatusSucceeded, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusUnknown {
		return notifyenum.NotificationChannelStatusUnknown, nil
	}
	claimed, err := u.larkWebhookDeliveryRepo.Claim(ctx, delivery.ID, time.Now(), u.externalProcessingTimeout, false)
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if !claimed {
		return delivery.Status, nil
	}
	result, err := u.larkWebhookClient.SendLarkWebhook(ctx, &bizchannel.LarkWebhookRequest{
		IdempotencyKey: fmt.Sprintf("%d", delivery.ID),
		Token:          rule.LarkWebhookTemplate.Token,
		Secret:         rule.LarkWebhookTemplate.Secret,
		RequestBody:    delivery.RequestBody,
	})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if result == nil {
		return notifyenum.NotificationChannelStatusUnknown, u.larkWebhookDeliveryRepo.MarkUnknown(ctx, delivery.ID, nil, nil)
	}
	switch result.Status {
	case notifyenum.NotificationChannelStatusSucceeded:
		err = u.larkWebhookDeliveryRepo.MarkSucceeded(ctx, delivery.ID, result.HTTPStatus, result.ResponseBody, time.Now())
		return notifyenum.NotificationChannelStatusSucceeded, err
	case notifyenum.NotificationChannelStatusFailed:
		err = u.larkWebhookDeliveryRepo.MarkFailed(ctx, delivery.ID, result.HTTPStatus, result.ResponseBody)
		return notifyenum.NotificationChannelStatusFailed, err
	default:
		err = u.larkWebhookDeliveryRepo.MarkUnknown(ctx, delivery.ID, result.HTTPStatus, result.ResponseBody)
		return notifyenum.NotificationChannelStatusUnknown, err
	}
}

func (u *NotifyUsecase) loadAccounts(ctx context.Context, notificationContext *NotificationContext, rules []*model.NotificationRule) (map[int64]*model.UserAccount, error) {
	needsContact := false
	for _, rule := range rules {
		if rule == nil || !rule.Enabled {
			continue
		}
		if rule.Channel == notifyenum.NotificationChannelEmail || (rule.Channel == notifyenum.NotificationChannelTencentSMS && u.smsEnabled) {
			needsContact = true
			break
		}
	}
	if !needsContact || u.userClient == nil {
		return map[int64]*model.UserAccount{}, nil
	}
	userIDs := make([]int64, 0)
	seen := map[int64]struct{}{}
	for _, recipient := range notificationContext.Recipients {
		if recipient == nil || recipient.UserID == 0 {
			continue
		}
		if recipient.Email != "" && recipient.Phone != "" {
			continue
		}
		if _, ok := seen[recipient.UserID]; ok {
			continue
		}
		seen[recipient.UserID] = struct{}{}
		userIDs = append(userIDs, recipient.UserID)
	}
	if len(userIDs) == 0 {
		return map[int64]*model.UserAccount{}, nil
	}
	return u.userClient.MapAccounts(ctx, userIDs)
}

func (u *NotifyUsecase) renderTemplate(tplStr string, data any) (string, bool) {
	tpl, err := template.New("").Option("missingkey=error").Parse(tplStr)
	if err != nil {
		return "", false
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", false
	}
	return buf.String(), true
}
