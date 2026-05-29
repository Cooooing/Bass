package channel

import (
	"context"
	notifyenum "notify/internal/enum"
)

type SendResult struct {
	Status            notifyenum.NotificationChannelStatus
	ProviderMessageID *string
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
	ProviderResponse  *string
	HTTPStatus        *int
	ResponseBody      *string
}

type EmailClient interface {
	SendEmail(ctx context.Context, req *EmailRequest) (*SendResult, error)
}

type EmailRequest struct {
	IdempotencyKey string
	ToEmail        string
	Subject        string
	Body           string
	ContentType    string
}

type TencentSMSClient interface {
	SendTencentSMS(ctx context.Context, req *TencentSMSRequest) (*SendResult, error)
}

type TencentSMSRequest struct {
	IdempotencyKey     string
	Phone              string
	SMSSDKAppID        string
	SignName           string
	ProviderTemplateID string
	TemplateParams     []string
}

type LarkWebhookClient interface {
	SendLarkWebhook(ctx context.Context, req *LarkWebhookRequest) (*SendResult, error)
}

type LarkWebhookRequest struct {
	IdempotencyKey string
	Token          string
	RequestBody    string
}
