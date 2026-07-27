package repo

import (
	"context"
	notifyenum "notify/internal/enum"
)

type TencentSMSClient interface {
	SendTencentSMS(ctx context.Context, req *TencentSMSRequest) (*TencentSMSSendResult, error)
}

type TencentSMSRequest struct {
	IdempotencyKey     string
	Phone              string
	SMSSDKAppID        string
	SignName           string
	ProviderTemplateID string
	TemplateParams     []string
}

type TencentSMSSendResult struct {
	Status            notifyenum.NotificationChannelStatus
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
	ProviderResp      *string
}
