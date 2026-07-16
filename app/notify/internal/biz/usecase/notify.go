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
	"notify/internal/config"
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
	conf *config.Bootstrap,
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
	if conf != nil && conf.Notify != nil && conf.Notify.NotificationRateLimit != nil {
		rateLimitEnabled = conf.Notify.NotificationRateLimit.Enable
		if conf.Notify.NotificationRateLimit.Window != nil && conf.Notify.NotificationRateLimit.Window.AsDuration() > 0 {
			rateLimitWindow = conf.Notify.NotificationRateLimit.Window.AsDuration()
		}
		if conf.Notify.NotificationRateLimit.MaxCount > 0 {
			rateLimitMaxCount = conf.Notify.NotificationRateLimit.MaxCount
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
		smsEnabled:                 conf != nil && conf.Notify != nil && conf.Notify.Sms != nil && conf.Notify.Sms.Enable,
		rateLimitEnabled:           rateLimitEnabled,
		rateLimitWindow:            rateLimitWindow,
		rateLimitMaxCount:          rateLimitMaxCount,
		externalProcessingTimeout:  10 * time.Minute,
	}
}

type NotifyListEnabledRulesReq struct {
	EventType commonenum.EventType
	Language  notifyenum.Language
}

type NotifyListEnabledRulesResponse struct {
	Rules []*model.NotificationRule
}

func (u *NotifyUsecase) ListEnabledRules(ctx context.Context, req *NotifyListEnabledRulesReq) (*NotifyListEnabledRulesResponse, error) {
	if req == nil {
		req = &NotifyListEnabledRulesReq{}
	}
	enabled := true
	rulesResponse, err := u.notificationRuleRepo.List(ctx, &repo.NotificationRuleListReq{Query: &repo.NotificationRuleQuery{
		EventType: &req.EventType,
		Language:  &req.Language,
		Enabled:   &enabled,
	}})
	if err != nil {
		return nil, err
	}
	return &NotifyListEnabledRulesResponse{Rules: rulesResponse.Rows}, nil
}

type NotifyProcessReq struct {
	NotificationContext *NotificationContext
	Rules               []*model.NotificationRule
}

func (u *NotifyUsecase) Process(ctx context.Context, req *NotifyProcessReq) error {
	if req == nil || req.NotificationContext == nil || req.NotificationContext.EventID == "" {
		return nil
	}
	notificationContext := req.NotificationContext
	rules := req.Rules
	loadAccountsResp, err := u.loadAccounts(ctx, &loadAccountsReq{NotificationContext: notificationContext, Rules: rules})
	if err != nil {
		return err
	}
	accountsByUserID := loadAccountsResp.AccountsByUserID
	for _, rule := range rules {
		processRuleResp, err := u.processRule(ctx, &processRuleReq{NotificationContext: notificationContext, Rule: rule, AccountsByUserID: accountsByUserID})
		if err != nil {
			return err
		}
		switch processRuleResp.Status {
		case notifyenum.NotificationChannelStatusProcessing, notifyenum.NotificationChannelStatusFailed, notifyenum.NotificationChannelStatusInternalError:
			return fmt.Errorf("notification channel not completed: event_id=%s channel=%s status=%s", notificationContext.EventID, rule.Channel, processRuleResp.Status)
		}
	}
	return nil
}

type processRuleReq struct {
	NotificationContext *NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

type processRuleResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) processRule(ctx context.Context, req *processRuleReq) (*processRuleResponse, error) {
	rule := req.Rule
	if rule == nil || !rule.Enabled {
		return &processRuleResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
	switch rule.Channel {
	case notifyenum.NotificationChannelStation:
		resp, err := u.processStation(ctx, &processStationReq{NotificationContext: req.NotificationContext, Rule: rule})
		if err != nil {
			return nil, err
		}
		return &processRuleResponse{Status: resp.Status}, nil
	case notifyenum.NotificationChannelEmail:
		resp, err := u.processEmail(ctx, &processEmailReq{NotificationContext: req.NotificationContext, Rule: rule, AccountsByUserID: req.AccountsByUserID})
		if err != nil {
			return nil, err
		}
		return &processRuleResponse{Status: resp.Status}, nil
	case notifyenum.NotificationChannelTencentSMS:
		resp, err := u.processTencentSMS(ctx, &processTencentSMSReq{NotificationContext: req.NotificationContext, Rule: rule, AccountsByUserID: req.AccountsByUserID})
		if err != nil {
			return nil, err
		}
		return &processRuleResponse{Status: resp.Status}, nil
	case notifyenum.NotificationChannelLarkWebhook:
		resp, err := u.processLarkWebhook(ctx, &processLarkWebhookReq{NotificationContext: req.NotificationContext, Rule: rule})
		if err != nil {
			return nil, err
		}
		return &processRuleResponse{Status: resp.Status}, nil
	default:
		return &processRuleResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
}

type processStationReq struct {
	NotificationContext *NotificationContext
	Rule                *model.NotificationRule
}

type processStationResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) processStation(ctx context.Context, req *processStationReq) (*processStationResponse, error) {
	notificationContext := req.NotificationContext
	rule := req.Rule
	if rule.StationTemplate == nil {
		return &processStationResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
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
		_, err := u.stationMessageRepo.Save(ctx, &repo.NotificationStationMessageSaveReq{Message: &model.NotificationStationMessage{
			EventID:    notificationContext.EventID,
			EventType:  notificationContext.EventType,
			ReceiverID: recipient.UserID,
			Title:      title,
			Content:    content,
			Status:     notifyenum.NotificationChannelStatusSucceeded,
		}})
		if err != nil {
			return &processStationResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
		}
		written = true
	}
	if !written {
		return &processStationResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
	return &processStationResponse{Status: notifyenum.NotificationChannelStatusSucceeded}, nil
}

type processEmailReq struct {
	NotificationContext *NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

type processEmailResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) processEmail(ctx context.Context, req *processEmailReq) (*processEmailResponse, error) {
	notificationContext := req.NotificationContext
	rule := req.Rule
	accountsByUserID := req.AccountsByUserID
	if rule.EmailTemplate == nil {
		return &processEmailResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
	if u.emailClient == nil {
		return &processEmailResponse{Status: notifyenum.NotificationChannelStatusInternalError}, fmt.Errorf("email client is nil")
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
		sendEmailResp, err := u.sendEmail(ctx, &sendEmailReq{Delivery: delivery})
		itemStatus := sendEmailResp.Status
		if err != nil {
			return &processEmailResponse{Status: itemStatus}, err
		}
		status = status.Merge(itemStatus)
		if status.Blocking() {
			return &processEmailResponse{Status: status}, nil
		}
	}
	return &processEmailResponse{Status: status}, nil
}

type sendEmailReq struct {
	Delivery *model.NotificationEmailDelivery
}

type sendEmailResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) sendEmail(ctx context.Context, req *sendEmailReq) (*sendEmailResponse, error) {
	delivery := req.Delivery
	saveResp, err := u.emailDeliveryRepo.SaveOrGet(ctx, &repo.NotificationEmailDeliverySaveOrGetReq{Delivery: delivery})
	if err == nil {
		delivery = saveResp.Delivery
	}
	if err != nil {
		return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if delivery.Status == notifyenum.NotificationChannelStatusSucceeded {
		return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusSucceeded}, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusUnknown {
		return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusUnknown}, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusRateLimited {
		return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusRateLimited}, nil
	}
	claimResp, err := u.emailDeliveryRepo.Claim(ctx, &repo.NotificationEmailDeliveryClaimReq{ID: delivery.ID, Now: time.Now(), ProcessingTimeout: u.externalProcessingTimeout})
	if err != nil {
		return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if !claimResp.Claimed {
		return &sendEmailResponse{Status: delivery.Status}, nil
	}
	if u.rateLimitEnabled {
		allowResp, err := u.notificationRateLimitCache.Allow(ctx, &repo.NotificationRateLimitAllowReq{Spec: &repo.NotificationRateLimitSpec{
			Channel:   notifyenum.NotificationChannelEmail,
			Recipient: delivery.ToEmail,
			Window:    u.rateLimitWindow,
			MaxCount:  u.rateLimitMaxCount,
		}})
		if err != nil {
			return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
		}
		if !allowResp.Allowed {
			if _, err := u.emailDeliveryRepo.MarkRateLimited(ctx, &repo.NotificationEmailDeliveryMarkRateLimitedReq{ID: delivery.ID}); err != nil {
				return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
			}
			return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusRateLimited}, nil
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
		return &sendEmailResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	finishResp, err := u.finishEmail(ctx, &finishEmailReq{DeliveryID: delivery.ID, Result: result})
	if err != nil {
		return &sendEmailResponse{Status: finishResp.Status}, err
	}
	return &sendEmailResponse{Status: finishResp.Status}, nil
}

type finishEmailReq struct {
	DeliveryID int64
	Result     *bizchannel.SendResult
}

type finishEmailResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) finishEmail(ctx context.Context, req *finishEmailReq) (*finishEmailResponse, error) {
	deliveryID := req.DeliveryID
	result := req.Result
	if result == nil {
		_, err := u.emailDeliveryRepo.MarkUnknown(ctx, &repo.NotificationEmailDeliveryMarkUnknownReq{ID: deliveryID})
		return &finishEmailResponse{Status: notifyenum.NotificationChannelStatusUnknown}, err
	}
	switch result.Status {
	case notifyenum.NotificationChannelStatusSucceeded:
		_, err := u.emailDeliveryRepo.MarkSucceeded(ctx, &repo.NotificationEmailDeliveryMarkSucceededReq{ID: deliveryID, ProviderMessageID: result.ProviderMessageID, ProviderResponse: result.ProviderResponse, SentAt: time.Now()})
		return &finishEmailResponse{Status: notifyenum.NotificationChannelStatusSucceeded}, err
	case notifyenum.NotificationChannelStatusFailed:
		_, err := u.emailDeliveryRepo.MarkFailed(ctx, &repo.NotificationEmailDeliveryMarkFailedReq{ID: deliveryID, ProviderResponse: result.ProviderResponse})
		return &finishEmailResponse{Status: notifyenum.NotificationChannelStatusFailed}, err
	default:
		_, err := u.emailDeliveryRepo.MarkUnknown(ctx, &repo.NotificationEmailDeliveryMarkUnknownReq{ID: deliveryID, ProviderResponse: result.ProviderResponse})
		return &finishEmailResponse{Status: notifyenum.NotificationChannelStatusUnknown}, err
	}
}

type processTencentSMSReq struct {
	NotificationContext *NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

type processTencentSMSResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) processTencentSMS(ctx context.Context, req *processTencentSMSReq) (*processTencentSMSResponse, error) {
	notificationContext := req.NotificationContext
	rule := req.Rule
	accountsByUserID := req.AccountsByUserID
	if !u.smsEnabled {
		return &processTencentSMSResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
	if rule.TencentSMSTemplate == nil {
		return &processTencentSMSResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
	if u.tencentSMSClient == nil {
		return &processTencentSMSResponse{Status: notifyenum.NotificationChannelStatusInternalError}, fmt.Errorf("tencent sms client is nil")
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
		sendTencentSMSResp, err := u.sendTencentSMS(ctx, &sendTencentSMSReq{Delivery: delivery})
		itemStatus := sendTencentSMSResp.Status
		if err != nil {
			return &processTencentSMSResponse{Status: itemStatus}, err
		}
		status = status.Merge(itemStatus)
		if status.Blocking() {
			return &processTencentSMSResponse{Status: status}, nil
		}
	}
	return &processTencentSMSResponse{Status: status}, nil
}

type sendTencentSMSReq struct {
	Delivery *model.NotificationTencentSMSDelivery
}

type sendTencentSMSResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) sendTencentSMS(ctx context.Context, req *sendTencentSMSReq) (*sendTencentSMSResponse, error) {
	delivery := req.Delivery
	saveResp, err := u.tencentSMSDeliveryRepo.SaveOrGet(ctx, &repo.NotificationTencentSMSDeliverySaveOrGetReq{Delivery: delivery})
	if err == nil {
		delivery = saveResp.Delivery
	}
	if err != nil {
		return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if delivery.Status == notifyenum.NotificationChannelStatusSucceeded {
		return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusSucceeded}, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusUnknown {
		return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusUnknown}, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusRateLimited {
		return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusRateLimited}, nil
	}
	claimResp, err := u.tencentSMSDeliveryRepo.Claim(ctx, &repo.NotificationTencentSMSDeliveryClaimReq{ID: delivery.ID, Now: time.Now(), ProcessingTimeout: u.externalProcessingTimeout})
	if err != nil {
		return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if !claimResp.Claimed {
		return &sendTencentSMSResponse{Status: delivery.Status}, nil
	}
	if u.rateLimitEnabled {
		allowResp, err := u.notificationRateLimitCache.Allow(ctx, &repo.NotificationRateLimitAllowReq{Spec: &repo.NotificationRateLimitSpec{
			Channel:   notifyenum.NotificationChannelTencentSMS,
			Recipient: delivery.Phone,
			Window:    u.rateLimitWindow,
			MaxCount:  u.rateLimitMaxCount,
		}})
		if err != nil {
			return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
		}
		if !allowResp.Allowed {
			if _, err := u.tencentSMSDeliveryRepo.MarkRateLimited(ctx, &repo.NotificationTencentSMSDeliveryMarkRateLimitedReq{ID: delivery.ID}); err != nil {
				return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
			}
			return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusRateLimited}, nil
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
		return &sendTencentSMSResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	finishResp, err := u.finishTencentSMS(ctx, &finishTencentSMSReq{DeliveryID: delivery.ID, Result: result})
	if err != nil {
		return &sendTencentSMSResponse{Status: finishResp.Status}, err
	}
	return &sendTencentSMSResponse{Status: finishResp.Status}, nil
}

type finishTencentSMSReq struct {
	DeliveryID int64
	Result     *bizchannel.SendResult
}

type finishTencentSMSResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) finishTencentSMS(ctx context.Context, req *finishTencentSMSReq) (*finishTencentSMSResponse, error) {
	deliveryID := req.DeliveryID
	result := req.Result
	if result == nil {
		_, err := u.tencentSMSDeliveryRepo.MarkUnknown(ctx, &repo.NotificationTencentSMSDeliveryMarkUnknownReq{ID: deliveryID})
		return &finishTencentSMSResponse{Status: notifyenum.NotificationChannelStatusUnknown}, err
	}
	switch result.Status {
	case notifyenum.NotificationChannelStatusSucceeded:
		_, err := u.tencentSMSDeliveryRepo.MarkSucceeded(ctx, &repo.NotificationTencentSMSDeliveryMarkSucceededReq{ID: deliveryID, ProviderRequestID: result.ProviderRequestID, ProviderCode: result.ProviderCode, ProviderMessage: result.ProviderMessage, SentAt: time.Now()})
		return &finishTencentSMSResponse{Status: notifyenum.NotificationChannelStatusSucceeded}, err
	case notifyenum.NotificationChannelStatusFailed:
		_, err := u.tencentSMSDeliveryRepo.MarkFailed(ctx, &repo.NotificationTencentSMSDeliveryMarkFailedReq{ID: deliveryID, ProviderRequestID: result.ProviderRequestID, ProviderCode: result.ProviderCode, ProviderMessage: result.ProviderMessage})
		return &finishTencentSMSResponse{Status: notifyenum.NotificationChannelStatusFailed}, err
	default:
		_, err := u.tencentSMSDeliveryRepo.MarkUnknown(ctx, &repo.NotificationTencentSMSDeliveryMarkUnknownReq{ID: deliveryID, ProviderRequestID: result.ProviderRequestID, ProviderCode: result.ProviderCode, ProviderMessage: result.ProviderMessage})
		return &finishTencentSMSResponse{Status: notifyenum.NotificationChannelStatusUnknown}, err
	}
}

type processLarkWebhookReq struct {
	NotificationContext *NotificationContext
	Rule                *model.NotificationRule
}

type processLarkWebhookResponse struct {
	Status notifyenum.NotificationChannelStatus
}

func (u *NotifyUsecase) processLarkWebhook(ctx context.Context, req *processLarkWebhookReq) (*processLarkWebhookResponse, error) {
	notificationContext := req.NotificationContext
	rule := req.Rule
	if rule.LarkWebhookTemplate == nil || rule.LarkWebhookTemplate.WebhookID == "" || rule.LarkWebhookTemplate.Token == "" {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
	if u.larkWebhookClient == nil {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusInternalError}, fmt.Errorf("lark webhook client is nil")
	}
	renderedContent, ok := u.renderTemplate(rule.LarkWebhookTemplate.ContentTemplate, notificationContext.TemplateData)
	if !ok {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusSkipped}, nil
	}
	var content map[string]any
	if err := json.Unmarshal([]byte(renderedContent), &content); err != nil {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if content == nil {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusInternalError}, fmt.Errorf("lark webhook content must be json object")
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
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	delivery := &model.NotificationLarkWebhookDelivery{
		EventID:     notificationContext.EventID,
		EventType:   notificationContext.EventType,
		WebhookID:   rule.LarkWebhookTemplate.WebhookID,
		RequestBody: string(requestBodyBytes),
		Status:      notifyenum.NotificationChannelStatusProcessing,
	}
	saveResp, err := u.larkWebhookDeliveryRepo.SaveOrGet(ctx, &repo.NotificationLarkWebhookDeliverySaveOrGetReq{Delivery: delivery})
	if err == nil {
		delivery = saveResp.Delivery
	}
	if err != nil {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if delivery.Status == notifyenum.NotificationChannelStatusSucceeded {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusSucceeded}, nil
	}
	if delivery.Status == notifyenum.NotificationChannelStatusUnknown {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusUnknown}, nil
	}
	claimResp, err := u.larkWebhookDeliveryRepo.Claim(ctx, &repo.NotificationLarkWebhookDeliveryClaimReq{ID: delivery.ID, Now: time.Now(), ProcessingTimeout: u.externalProcessingTimeout})
	if err != nil {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if !claimResp.Claimed {
		return &processLarkWebhookResponse{Status: delivery.Status}, nil
	}
	result, err := u.larkWebhookClient.SendLarkWebhook(ctx, &bizchannel.LarkWebhookRequest{
		IdempotencyKey: fmt.Sprintf("%d", delivery.ID),
		Token:          rule.LarkWebhookTemplate.Token,
		Secret:         rule.LarkWebhookTemplate.Secret,
		RequestBody:    delivery.RequestBody,
	})
	if err != nil {
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusInternalError}, err
	}
	if result == nil {
		_, err := u.larkWebhookDeliveryRepo.MarkUnknown(ctx, &repo.NotificationLarkWebhookDeliveryMarkUnknownReq{ID: delivery.ID})
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusUnknown}, err
	}
	switch result.Status {
	case notifyenum.NotificationChannelStatusSucceeded:
		_, err = u.larkWebhookDeliveryRepo.MarkSucceeded(ctx, &repo.NotificationLarkWebhookDeliveryMarkSucceededReq{ID: delivery.ID, HTTPStatus: result.HTTPStatus, ResponseBody: result.ResponseBody, SentAt: time.Now()})
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusSucceeded}, err
	case notifyenum.NotificationChannelStatusFailed:
		_, err = u.larkWebhookDeliveryRepo.MarkFailed(ctx, &repo.NotificationLarkWebhookDeliveryMarkFailedReq{ID: delivery.ID, HTTPStatus: result.HTTPStatus, ResponseBody: result.ResponseBody})
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusFailed}, err
	default:
		_, err = u.larkWebhookDeliveryRepo.MarkUnknown(ctx, &repo.NotificationLarkWebhookDeliveryMarkUnknownReq{ID: delivery.ID, HTTPStatus: result.HTTPStatus, ResponseBody: result.ResponseBody})
		return &processLarkWebhookResponse{Status: notifyenum.NotificationChannelStatusUnknown}, err
	}
}

type loadAccountsReq struct {
	NotificationContext *NotificationContext
	Rules               []*model.NotificationRule
}

type loadAccountsResponse struct {
	AccountsByUserID map[int64]*model.UserAccount
}

func (u *NotifyUsecase) loadAccounts(ctx context.Context, req *loadAccountsReq) (*loadAccountsResponse, error) {
	notificationContext := req.NotificationContext
	rules := req.Rules
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
		return &loadAccountsResponse{AccountsByUserID: map[int64]*model.UserAccount{}}, nil
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
		return &loadAccountsResponse{AccountsByUserID: map[int64]*model.UserAccount{}}, nil
	}
	resp, err := u.userClient.MapAccounts(ctx, &repo.UserMapAccountsReq{UserIDs: userIDs})
	if err != nil {
		return nil, err
	}
	return &loadAccountsResponse{AccountsByUserID: resp.Rows}, nil
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
