package repo

import (
	"context"
	"notify/internal/biz/model"
	"time"
)

type NotificationEmailDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationEmailDelivery) (*model.NotificationEmailDelivery, error)
	Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error)
	MarkSucceeded(ctx context.Context, id int64, providerMessageID *string, providerResponse *string, sentAt time.Time) error
	MarkFailed(ctx context.Context, id int64, providerResponse *string) error
	MarkUnknown(ctx context.Context, id int64, providerResponse *string) error
	MarkRateLimited(ctx context.Context, id int64) error
}

type NotificationTencentSMSDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (*model.NotificationTencentSMSDelivery, error)
	Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error)
	MarkSucceeded(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string, sentAt time.Time) error
	MarkFailed(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string) error
	MarkUnknown(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string) error
	MarkRateLimited(ctx context.Context, id int64) error
}

type NotificationLarkWebhookDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationLarkWebhookDelivery) (*model.NotificationLarkWebhookDelivery, error)
	Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error)
	MarkSucceeded(ctx context.Context, id int64, httpStatus *int, responseBody *string, sentAt time.Time) error
	MarkFailed(ctx context.Context, id int64, httpStatus *int, responseBody *string) error
	MarkUnknown(ctx context.Context, id int64, httpStatus *int, responseBody *string) error
}
