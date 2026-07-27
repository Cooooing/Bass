package channel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	notifyenum "notify/internal/enum"
)

type LarkWebhookUsecase struct {
	larkWebhookDeliveryRepo   repo.NotificationLarkWebhookDeliveryRepo
	larkWebhookClient         repo.LarkWebhookClient
	externalProcessingTimeout time.Duration
}

func NewLarkWebhookUsecase(
	larkWebhookDeliveryRepo repo.NotificationLarkWebhookDeliveryRepo,
	larkWebhookClient repo.LarkWebhookClient,
) *LarkWebhookUsecase {
	return &LarkWebhookUsecase{
		larkWebhookDeliveryRepo:   larkWebhookDeliveryRepo,
		larkWebhookClient:         larkWebhookClient,
		externalProcessingTimeout: 10 * time.Minute,
	}
}

type LarkWebhookProcessReq struct {
	NotificationContext *model.NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

func (u *LarkWebhookUsecase) Process(ctx context.Context, req *LarkWebhookProcessReq) (notifyenum.NotificationChannelStatus, error) {
	if req == nil || req.NotificationContext == nil || req.Rule == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	notificationContext := req.NotificationContext
	rule := req.Rule
	if rule.LarkWebhookTemplate == nil || rule.LarkWebhookTemplate.WebhookID == "" || rule.LarkWebhookTemplate.Token == "" {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	if u.larkWebhookClient == nil {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("lark webhook client is nil")
	}
	rendered, err := rule.LarkWebhookTemplate.Render(ctx, notificationContext.TemplateData)
	if err != nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	content := json.RawMessage(bytes.TrimSpace([]byte(rendered.Content)))
	if len(content) == 0 || content[0] != '{' {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("lark webhook content must be json object")
	}
	if !json.Valid(content) {
		return notifyenum.NotificationChannelStatusInternalError, fmt.Errorf("lark webhook content must be valid json")
	}
	msgType := strings.TrimSpace(rendered.MsgType)
	if msgType == "" {
		msgType = "text"
	}
	requestBodyBytes, err := json.Marshal(struct {
		MsgType string          `json:"msg_type"`
		Content json.RawMessage `json:"content"`
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
		WebhookID:   rendered.WebhookID,
		RequestBody: string(requestBodyBytes),
		Status:      notifyenum.NotificationChannelStatusProcessing,
	}
	saveResp, err := u.larkWebhookDeliveryRepo.SaveOrGet(ctx, delivery)
	if err == nil {
		delivery = saveResp
	}
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if delivery.Status == notifyenum.NotificationChannelStatusSucceeded || delivery.Status == notifyenum.NotificationChannelStatusUnknown {
		return delivery.Status, nil
	}
	claimResp, err := u.larkWebhookDeliveryRepo.Claim(ctx, &repo.NotificationLarkWebhookDeliveryClaimReq{
		ID:                delivery.ID,
		Now:               time.Now(),
		ProcessingTimeout: u.externalProcessingTimeout,
	})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if !claimResp {
		return delivery.Status, nil
	}
	result, err := u.larkWebhookClient.SendLarkWebhook(ctx, &repo.LarkWebhookRequest{
		IdempotencyKey: fmt.Sprintf("%d", delivery.ID),
		Token:          rendered.Token,
		Secret:         rendered.Secret,
		RequestBody:    delivery.RequestBody,
	})
	if err != nil {
		return notifyenum.NotificationChannelStatusInternalError, err
	}
	if result == nil {
		err := u.larkWebhookDeliveryRepo.UpdateStatus(ctx, &repo.NotificationLarkWebhookDeliveryUpdateStatusReq{
			ID:     delivery.ID,
			Status: notifyenum.NotificationChannelStatusUnknown,
		})
		return notifyenum.NotificationChannelStatusUnknown, err
	}
	status := result.Status
	if status != notifyenum.NotificationChannelStatusSucceeded && status != notifyenum.NotificationChannelStatusFailed {
		status = notifyenum.NotificationChannelStatusUnknown
	}
	var sentAt *time.Time
	if status == notifyenum.NotificationChannelStatusSucceeded {
		now := time.Now()
		sentAt = &now
	}
	err = u.larkWebhookDeliveryRepo.UpdateStatus(ctx, &repo.NotificationLarkWebhookDeliveryUpdateStatusReq{
		ID:         delivery.ID,
		Status:     status,
		HTTPStatus: result.HTTPStatus,
		RespBody:   result.RespBody,
		SentAt:     sentAt,
	})
	return status, err
}
