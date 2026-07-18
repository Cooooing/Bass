package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationEmailDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationEmailDelivery) (*model.NotificationEmailDelivery, error)
	Get(ctx context.Context, query *NotificationEmailDeliveryQuery) (*model.NotificationEmailDelivery, error)
	List(ctx context.Context, query *NotificationEmailDeliveryQuery) ([]*model.NotificationEmailDelivery, error)
	Map(ctx context.Context, query *NotificationEmailDeliveryQuery) (map[int64]*model.NotificationEmailDelivery, error)
	Count(ctx context.Context, query *NotificationEmailDeliveryQuery) (int, error)
	Page(ctx context.Context, query *NotificationEmailDeliveryQuery) (*NotificationEmailDeliveryPageResp, error)
	Claim(ctx context.Context, req *NotificationEmailDeliveryClaimReq) (bool, error)
	MarkSucceeded(ctx context.Context, req *NotificationEmailDeliveryMarkSucceededReq) error
	MarkFailed(ctx context.Context, req *NotificationEmailDeliveryMarkFailedReq) error
	MarkUnknown(ctx context.Context, req *NotificationEmailDeliveryMarkUnknownReq) error
	MarkRateLimited(ctx context.Context, id int64) error
}

type NotificationTencentSMSDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (*model.NotificationTencentSMSDelivery, error)
	Get(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (*model.NotificationTencentSMSDelivery, error)
	List(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) ([]*model.NotificationTencentSMSDelivery, error)
	Map(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (map[int64]*model.NotificationTencentSMSDelivery, error)
	Count(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (int, error)
	Page(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (*NotificationTencentSMSDeliveryPageResp, error)
	Claim(ctx context.Context, req *NotificationTencentSMSDeliveryClaimReq) (bool, error)
	MarkSucceeded(ctx context.Context, req *NotificationTencentSMSDeliveryMarkSucceededReq) error
	MarkFailed(ctx context.Context, req *NotificationTencentSMSDeliveryMarkFailedReq) error
	MarkUnknown(ctx context.Context, req *NotificationTencentSMSDeliveryMarkUnknownReq) error
	MarkRateLimited(ctx context.Context, id int64) error
}

type NotificationLarkWebhookDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationLarkWebhookDelivery) (*model.NotificationLarkWebhookDelivery, error)
	Get(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (*model.NotificationLarkWebhookDelivery, error)
	List(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) ([]*model.NotificationLarkWebhookDelivery, error)
	Map(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (map[int64]*model.NotificationLarkWebhookDelivery, error)
	Count(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (int, error)
	Page(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (*NotificationLarkWebhookDeliveryPageResp, error)
	Claim(ctx context.Context, req *NotificationLarkWebhookDeliveryClaimReq) (bool, error)
	MarkSucceeded(ctx context.Context, req *NotificationLarkWebhookDeliveryMarkSucceededReq) error
	MarkFailed(ctx context.Context, req *NotificationLarkWebhookDeliveryMarkFailedReq) error
	MarkUnknown(ctx context.Context, req *NotificationLarkWebhookDeliveryMarkUnknownReq) error
}

type NotificationEmailDeliveryPageResp struct {
	Rows []*model.NotificationEmailDelivery
	Page *base.PageResp
}

type NotificationEmailDeliveryClaimReq struct {
	ID                int64
	Now               time.Time
	ProcessingTimeout time.Duration
	RetryUnknown      bool
}

type NotificationEmailDeliveryMarkSucceededReq struct {
	ID                int64
	ProviderMessageID *string
	ProviderResp      *string
	SentAt            time.Time
}

type NotificationEmailDeliveryMarkFailedReq struct {
	ID           int64
	ProviderResp *string
}

type NotificationEmailDeliveryMarkUnknownReq struct {
	ID           int64
	ProviderResp *string
}

type NotificationEmailDeliveryQuery struct {
	Page       *base.PageRequest
	ID         *int64
	IDs        []int64
	EventID    *string
	EventIDs   []string
	EventType  *commonenum.EventType
	ReceiverID *int64
	ToEmail    *string
	Status     *notifyenum.NotificationChannelStatus
}

type NotificationTencentSMSDeliveryPageResp struct {
	Rows []*model.NotificationTencentSMSDelivery
	Page *base.PageResp
}

type NotificationTencentSMSDeliveryClaimReq struct {
	ID                int64
	Now               time.Time
	ProcessingTimeout time.Duration
	RetryUnknown      bool
}

type NotificationTencentSMSDeliveryMarkSucceededReq struct {
	ID                int64
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
	SentAt            time.Time
}

type NotificationTencentSMSDeliveryMarkFailedReq struct {
	ID                int64
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
}

type NotificationTencentSMSDeliveryMarkUnknownReq struct {
	ID                int64
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
}

type NotificationTencentSMSDeliveryQuery struct {
	Page       *base.PageRequest
	ID         *int64
	IDs        []int64
	EventID    *string
	EventIDs   []string
	EventType  *commonenum.EventType
	ReceiverID *int64
	Phone      *string
	Status     *notifyenum.NotificationChannelStatus
}

type NotificationLarkWebhookDeliveryPageResp struct {
	Rows []*model.NotificationLarkWebhookDelivery
	Page *base.PageResp
}

type NotificationLarkWebhookDeliveryClaimReq struct {
	ID                int64
	Now               time.Time
	ProcessingTimeout time.Duration
	RetryUnknown      bool
}

type NotificationLarkWebhookDeliveryMarkSucceededReq struct {
	ID         int64
	HTTPStatus *int
	RespBody   *string
	SentAt     time.Time
}

type NotificationLarkWebhookDeliveryMarkFailedReq struct {
	ID         int64
	HTTPStatus *int
	RespBody   *string
}

type NotificationLarkWebhookDeliveryMarkUnknownReq struct {
	ID         int64
	HTTPStatus *int
	RespBody   *string
}

type NotificationLarkWebhookDeliveryQuery struct {
	Page      *base.PageRequest
	ID        *int64
	IDs       []int64
	EventID   *string
	EventIDs  []string
	EventType *commonenum.EventType
	WebhookID *string
	Status    *notifyenum.NotificationChannelStatus
}
