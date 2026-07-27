package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationLarkWebhookDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationLarkWebhookDelivery) (*model.NotificationLarkWebhookDelivery, error)
	Get(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (*model.NotificationLarkWebhookDelivery, error)
	List(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) ([]*model.NotificationLarkWebhookDelivery, error)
	Map(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (map[int64]*model.NotificationLarkWebhookDelivery, error)
	Count(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (int, error)
	Page(ctx context.Context, query *NotificationLarkWebhookDeliveryQuery) (*NotificationLarkWebhookDeliveryPageResp, error)
	Claim(ctx context.Context, req *NotificationLarkWebhookDeliveryClaimReq) (bool, error)
	UpdateStatus(ctx context.Context, req *NotificationLarkWebhookDeliveryUpdateStatusReq) error
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

type NotificationLarkWebhookDeliveryUpdateStatusReq struct {
	ID         int64
	Status     notifyenum.NotificationChannelStatus
	HTTPStatus *int
	RespBody   *string
	SentAt     *time.Time
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
