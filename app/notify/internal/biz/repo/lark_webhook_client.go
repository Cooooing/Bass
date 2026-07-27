package repo

import (
	"context"
	notifyenum "notify/internal/enum"
)

type LarkWebhookClient interface {
	SendLarkWebhook(ctx context.Context, req *LarkWebhookRequest) (*LarkWebhookSendResult, error)
}

type LarkWebhookRequest struct {
	IdempotencyKey string
	Token          string
	Secret         string
	RequestBody    string
}

type LarkWebhookSendResult struct {
	Status     notifyenum.NotificationChannelStatus
	HTTPStatus *int
	RespBody   *string
}
