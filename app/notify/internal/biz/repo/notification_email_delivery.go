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
	UpdateStatus(ctx context.Context, req *NotificationEmailDeliveryUpdateStatusReq) error
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

type NotificationEmailDeliveryUpdateStatusReq struct {
	ID                int64
	Status            notifyenum.NotificationChannelStatus
	ProviderMessageID *string
	ProviderResp      *string
	SentAt            *time.Time
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
