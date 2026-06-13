package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	"context"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationEmailDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationEmailDelivery) (*model.NotificationEmailDelivery, error)
	Get(ctx context.Context, req *NotificationEmailDeliveryQuery) (*model.NotificationEmailDelivery, error)
	List(ctx context.Context, req *NotificationEmailDeliveryQuery) ([]*model.NotificationEmailDelivery, error)
	Map(ctx context.Context, req *NotificationEmailDeliveryQuery) (map[int64]*model.NotificationEmailDelivery, error)
	Count(ctx context.Context, req *NotificationEmailDeliveryQuery) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *NotificationEmailDeliveryQuery) ([]*model.NotificationEmailDelivery, *common.PageReply, error)
	Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error)
	MarkSucceeded(ctx context.Context, id int64, providerMessageID *string, providerResponse *string, sentAt time.Time) error
	MarkFailed(ctx context.Context, id int64, providerResponse *string) error
	MarkUnknown(ctx context.Context, id int64, providerResponse *string) error
	MarkRateLimited(ctx context.Context, id int64) error
}

type NotificationTencentSMSDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (*model.NotificationTencentSMSDelivery, error)
	Get(ctx context.Context, req *NotificationTencentSMSDeliveryQuery) (*model.NotificationTencentSMSDelivery, error)
	List(ctx context.Context, req *NotificationTencentSMSDeliveryQuery) ([]*model.NotificationTencentSMSDelivery, error)
	Map(ctx context.Context, req *NotificationTencentSMSDeliveryQuery) (map[int64]*model.NotificationTencentSMSDelivery, error)
	Count(ctx context.Context, req *NotificationTencentSMSDeliveryQuery) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *NotificationTencentSMSDeliveryQuery) ([]*model.NotificationTencentSMSDelivery, *common.PageReply, error)
	Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error)
	MarkSucceeded(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string, sentAt time.Time) error
	MarkFailed(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string) error
	MarkUnknown(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string) error
	MarkRateLimited(ctx context.Context, id int64) error
}

type NotificationLarkWebhookDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationLarkWebhookDelivery) (*model.NotificationLarkWebhookDelivery, error)
	Get(ctx context.Context, req *NotificationLarkWebhookDeliveryQuery) (*model.NotificationLarkWebhookDelivery, error)
	List(ctx context.Context, req *NotificationLarkWebhookDeliveryQuery) ([]*model.NotificationLarkWebhookDelivery, error)
	Map(ctx context.Context, req *NotificationLarkWebhookDeliveryQuery) (map[int64]*model.NotificationLarkWebhookDelivery, error)
	Count(ctx context.Context, req *NotificationLarkWebhookDeliveryQuery) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *NotificationLarkWebhookDeliveryQuery) ([]*model.NotificationLarkWebhookDelivery, *common.PageReply, error)
	Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error)
	MarkSucceeded(ctx context.Context, id int64, httpStatus *int, responseBody *string, sentAt time.Time) error
	MarkFailed(ctx context.Context, id int64, httpStatus *int, responseBody *string) error
	MarkUnknown(ctx context.Context, id int64, httpStatus *int, responseBody *string) error
}

type NotificationEmailDeliveryQuery struct {
	ID         *int64
	IDs        []int64
	EventID    *string
	EventIDs   []string
	EventType  *commonenum.EventType
	ReceiverID *int64
	ToEmail    *string
	Status     *notifyenum.NotificationChannelStatus
}

type NotificationTencentSMSDeliveryQuery struct {
	ID         *int64
	IDs        []int64
	EventID    *string
	EventIDs   []string
	EventType  *commonenum.EventType
	ReceiverID *int64
	Phone      *string
	Status     *notifyenum.NotificationChannelStatus
}

type NotificationLarkWebhookDeliveryQuery struct {
	ID        *int64
	IDs       []int64
	EventID   *string
	EventIDs  []string
	EventType *commonenum.EventType
	WebhookID *string
	Status    *notifyenum.NotificationChannelStatus
}
