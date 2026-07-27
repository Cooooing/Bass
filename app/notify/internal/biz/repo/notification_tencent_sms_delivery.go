package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationTencentSMSDeliveryRepo interface {
	SaveOrGet(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (*model.NotificationTencentSMSDelivery, error)
	Get(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (*model.NotificationTencentSMSDelivery, error)
	List(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) ([]*model.NotificationTencentSMSDelivery, error)
	Map(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (map[int64]*model.NotificationTencentSMSDelivery, error)
	Count(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (int, error)
	Page(ctx context.Context, query *NotificationTencentSMSDeliveryQuery) (*NotificationTencentSMSDeliveryPageResp, error)
	Claim(ctx context.Context, req *NotificationTencentSMSDeliveryClaimReq) (bool, error)
	UpdateStatus(ctx context.Context, req *NotificationTencentSMSDeliveryUpdateStatusReq) error
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

type NotificationTencentSMSDeliveryUpdateStatusReq struct {
	ID                int64
	Status            notifyenum.NotificationChannelStatus
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
	SentAt            *time.Time
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
