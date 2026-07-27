package repo

import (
	"context"
	notifyenum "notify/internal/enum"
)

type EmailClient interface {
	SendEmail(ctx context.Context, req *EmailRequest) (*EmailSendResult, error)
}

type EmailRequest struct {
	IdempotencyKey string
	ToEmail        string
	Subject        string
	Body           string
	ContentType    string
}

type EmailSendResult struct {
	Status            notifyenum.NotificationChannelStatus
	ProviderMessageID *string
	ProviderResp      *string
}
