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
	SaveOrGet(ctx context.Context, req *NotificationEmailDeliverySaveOrGetReq) (*NotificationEmailDeliverySaveOrGetResponse, error)
	Get(ctx context.Context, req *NotificationEmailDeliveryGetReq) (*NotificationEmailDeliveryGetResponse, error)
	List(ctx context.Context, req *NotificationEmailDeliveryListReq) (*NotificationEmailDeliveryListResponse, error)
	Map(ctx context.Context, req *NotificationEmailDeliveryMapReq) (*NotificationEmailDeliveryMapResponse, error)
	Count(ctx context.Context, req *NotificationEmailDeliveryCountReq) (*NotificationEmailDeliveryCountResponse, error)
	Page(ctx context.Context, req *NotificationEmailDeliveryPageReq) (*NotificationEmailDeliveryPageResponse, error)
	Claim(ctx context.Context, req *NotificationEmailDeliveryClaimReq) (*NotificationEmailDeliveryClaimResponse, error)
	MarkSucceeded(ctx context.Context, req *NotificationEmailDeliveryMarkSucceededReq) (*NotificationEmailDeliveryMarkSucceededResponse, error)
	MarkFailed(ctx context.Context, req *NotificationEmailDeliveryMarkFailedReq) (*NotificationEmailDeliveryMarkFailedResponse, error)
	MarkUnknown(ctx context.Context, req *NotificationEmailDeliveryMarkUnknownReq) (*NotificationEmailDeliveryMarkUnknownResponse, error)
	MarkRateLimited(ctx context.Context, req *NotificationEmailDeliveryMarkRateLimitedReq) (*NotificationEmailDeliveryMarkRateLimitedResponse, error)
}

type NotificationTencentSMSDeliveryRepo interface {
	SaveOrGet(ctx context.Context, req *NotificationTencentSMSDeliverySaveOrGetReq) (*NotificationTencentSMSDeliverySaveOrGetResponse, error)
	Get(ctx context.Context, req *NotificationTencentSMSDeliveryGetReq) (*NotificationTencentSMSDeliveryGetResponse, error)
	List(ctx context.Context, req *NotificationTencentSMSDeliveryListReq) (*NotificationTencentSMSDeliveryListResponse, error)
	Map(ctx context.Context, req *NotificationTencentSMSDeliveryMapReq) (*NotificationTencentSMSDeliveryMapResponse, error)
	Count(ctx context.Context, req *NotificationTencentSMSDeliveryCountReq) (*NotificationTencentSMSDeliveryCountResponse, error)
	Page(ctx context.Context, req *NotificationTencentSMSDeliveryPageReq) (*NotificationTencentSMSDeliveryPageResponse, error)
	Claim(ctx context.Context, req *NotificationTencentSMSDeliveryClaimReq) (*NotificationTencentSMSDeliveryClaimResponse, error)
	MarkSucceeded(ctx context.Context, req *NotificationTencentSMSDeliveryMarkSucceededReq) (*NotificationTencentSMSDeliveryMarkSucceededResponse, error)
	MarkFailed(ctx context.Context, req *NotificationTencentSMSDeliveryMarkFailedReq) (*NotificationTencentSMSDeliveryMarkFailedResponse, error)
	MarkUnknown(ctx context.Context, req *NotificationTencentSMSDeliveryMarkUnknownReq) (*NotificationTencentSMSDeliveryMarkUnknownResponse, error)
	MarkRateLimited(ctx context.Context, req *NotificationTencentSMSDeliveryMarkRateLimitedReq) (*NotificationTencentSMSDeliveryMarkRateLimitedResponse, error)
}

type NotificationLarkWebhookDeliveryRepo interface {
	SaveOrGet(ctx context.Context, req *NotificationLarkWebhookDeliverySaveOrGetReq) (*NotificationLarkWebhookDeliverySaveOrGetResponse, error)
	Get(ctx context.Context, req *NotificationLarkWebhookDeliveryGetReq) (*NotificationLarkWebhookDeliveryGetResponse, error)
	List(ctx context.Context, req *NotificationLarkWebhookDeliveryListReq) (*NotificationLarkWebhookDeliveryListResponse, error)
	Map(ctx context.Context, req *NotificationLarkWebhookDeliveryMapReq) (*NotificationLarkWebhookDeliveryMapResponse, error)
	Count(ctx context.Context, req *NotificationLarkWebhookDeliveryCountReq) (*NotificationLarkWebhookDeliveryCountResponse, error)
	Page(ctx context.Context, req *NotificationLarkWebhookDeliveryPageReq) (*NotificationLarkWebhookDeliveryPageResponse, error)
	Claim(ctx context.Context, req *NotificationLarkWebhookDeliveryClaimReq) (*NotificationLarkWebhookDeliveryClaimResponse, error)
	MarkSucceeded(ctx context.Context, req *NotificationLarkWebhookDeliveryMarkSucceededReq) (*NotificationLarkWebhookDeliveryMarkSucceededResponse, error)
	MarkFailed(ctx context.Context, req *NotificationLarkWebhookDeliveryMarkFailedReq) (*NotificationLarkWebhookDeliveryMarkFailedResponse, error)
	MarkUnknown(ctx context.Context, req *NotificationLarkWebhookDeliveryMarkUnknownReq) (*NotificationLarkWebhookDeliveryMarkUnknownResponse, error)
}

type NotificationEmailDeliverySaveOrGetReq struct {
	Delivery *model.NotificationEmailDelivery
}

type NotificationEmailDeliverySaveOrGetResponse struct {
	Delivery *model.NotificationEmailDelivery
}

type NotificationEmailDeliveryGetReq struct {
	Query *NotificationEmailDeliveryQuery
}

type NotificationEmailDeliveryGetResponse struct {
	Item *model.NotificationEmailDelivery
}

type NotificationEmailDeliveryListReq struct {
	Query *NotificationEmailDeliveryQuery
}

type NotificationEmailDeliveryListResponse struct {
	Rows []*model.NotificationEmailDelivery
}

type NotificationEmailDeliveryMapReq struct {
	Query *NotificationEmailDeliveryQuery
}

type NotificationEmailDeliveryMapResponse struct {
	Rows map[int64]*model.NotificationEmailDelivery
}

type NotificationEmailDeliveryCountReq struct {
	Query *NotificationEmailDeliveryQuery
}

type NotificationEmailDeliveryCountResponse struct {
	Count int
}

type NotificationEmailDeliveryPageReq struct {
	Query *NotificationEmailDeliveryQuery
}

type NotificationEmailDeliveryPageResponse struct {
	Rows []*model.NotificationEmailDelivery
	Page *base.PageResponse
}

type NotificationEmailDeliveryClaimReq struct {
	ID                int64
	Now               time.Time
	ProcessingTimeout time.Duration
	RetryUnknown      bool
}

type NotificationEmailDeliveryClaimResponse struct {
	Claimed bool
}

type NotificationEmailDeliveryMarkSucceededReq struct {
	ID                int64
	ProviderMessageID *string
	ProviderResponse  *string
	SentAt            time.Time
}

type NotificationEmailDeliveryMarkSucceededResponse struct{}

type NotificationEmailDeliveryMarkFailedReq struct {
	ID               int64
	ProviderResponse *string
}

type NotificationEmailDeliveryMarkFailedResponse struct{}

type NotificationEmailDeliveryMarkUnknownReq struct {
	ID               int64
	ProviderResponse *string
}

type NotificationEmailDeliveryMarkUnknownResponse struct{}

type NotificationEmailDeliveryMarkRateLimitedReq struct {
	ID int64
}

type NotificationEmailDeliveryMarkRateLimitedResponse struct{}

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

type NotificationTencentSMSDeliverySaveOrGetReq struct {
	Delivery *model.NotificationTencentSMSDelivery
}

type NotificationTencentSMSDeliverySaveOrGetResponse struct {
	Delivery *model.NotificationTencentSMSDelivery
}

type NotificationTencentSMSDeliveryGetReq struct {
	Query *NotificationTencentSMSDeliveryQuery
}

type NotificationTencentSMSDeliveryGetResponse struct {
	Item *model.NotificationTencentSMSDelivery
}

type NotificationTencentSMSDeliveryListReq struct {
	Query *NotificationTencentSMSDeliveryQuery
}

type NotificationTencentSMSDeliveryListResponse struct {
	Rows []*model.NotificationTencentSMSDelivery
}

type NotificationTencentSMSDeliveryMapReq struct {
	Query *NotificationTencentSMSDeliveryQuery
}

type NotificationTencentSMSDeliveryMapResponse struct {
	Rows map[int64]*model.NotificationTencentSMSDelivery
}

type NotificationTencentSMSDeliveryCountReq struct {
	Query *NotificationTencentSMSDeliveryQuery
}

type NotificationTencentSMSDeliveryCountResponse struct {
	Count int
}

type NotificationTencentSMSDeliveryPageReq struct {
	Query *NotificationTencentSMSDeliveryQuery
}

type NotificationTencentSMSDeliveryPageResponse struct {
	Rows []*model.NotificationTencentSMSDelivery
	Page *base.PageResponse
}

type NotificationTencentSMSDeliveryClaimReq struct {
	ID                int64
	Now               time.Time
	ProcessingTimeout time.Duration
	RetryUnknown      bool
}

type NotificationTencentSMSDeliveryClaimResponse struct {
	Claimed bool
}

type NotificationTencentSMSDeliveryMarkSucceededReq struct {
	ID                int64
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
	SentAt            time.Time
}

type NotificationTencentSMSDeliveryMarkSucceededResponse struct{}

type NotificationTencentSMSDeliveryMarkFailedReq struct {
	ID                int64
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
}

type NotificationTencentSMSDeliveryMarkFailedResponse struct{}

type NotificationTencentSMSDeliveryMarkUnknownReq struct {
	ID                int64
	ProviderRequestID *string
	ProviderCode      *string
	ProviderMessage   *string
}

type NotificationTencentSMSDeliveryMarkUnknownResponse struct{}

type NotificationTencentSMSDeliveryMarkRateLimitedReq struct {
	ID int64
}

type NotificationTencentSMSDeliveryMarkRateLimitedResponse struct{}

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

type NotificationLarkWebhookDeliverySaveOrGetReq struct {
	Delivery *model.NotificationLarkWebhookDelivery
}

type NotificationLarkWebhookDeliverySaveOrGetResponse struct {
	Delivery *model.NotificationLarkWebhookDelivery
}

type NotificationLarkWebhookDeliveryGetReq struct {
	Query *NotificationLarkWebhookDeliveryQuery
}

type NotificationLarkWebhookDeliveryGetResponse struct {
	Item *model.NotificationLarkWebhookDelivery
}

type NotificationLarkWebhookDeliveryListReq struct {
	Query *NotificationLarkWebhookDeliveryQuery
}

type NotificationLarkWebhookDeliveryListResponse struct {
	Rows []*model.NotificationLarkWebhookDelivery
}

type NotificationLarkWebhookDeliveryMapReq struct {
	Query *NotificationLarkWebhookDeliveryQuery
}

type NotificationLarkWebhookDeliveryMapResponse struct {
	Rows map[int64]*model.NotificationLarkWebhookDelivery
}

type NotificationLarkWebhookDeliveryCountReq struct {
	Query *NotificationLarkWebhookDeliveryQuery
}

type NotificationLarkWebhookDeliveryCountResponse struct {
	Count int
}

type NotificationLarkWebhookDeliveryPageReq struct {
	Query *NotificationLarkWebhookDeliveryQuery
}

type NotificationLarkWebhookDeliveryPageResponse struct {
	Rows []*model.NotificationLarkWebhookDelivery
	Page *base.PageResponse
}

type NotificationLarkWebhookDeliveryClaimReq struct {
	ID                int64
	Now               time.Time
	ProcessingTimeout time.Duration
	RetryUnknown      bool
}

type NotificationLarkWebhookDeliveryClaimResponse struct {
	Claimed bool
}

type NotificationLarkWebhookDeliveryMarkSucceededReq struct {
	ID           int64
	HTTPStatus   *int
	ResponseBody *string
	SentAt       time.Time
}

type NotificationLarkWebhookDeliveryMarkSucceededResponse struct{}

type NotificationLarkWebhookDeliveryMarkFailedReq struct {
	ID           int64
	HTTPStatus   *int
	ResponseBody *string
}

type NotificationLarkWebhookDeliveryMarkFailedResponse struct{}

type NotificationLarkWebhookDeliveryMarkUnknownReq struct {
	ID           int64
	HTTPStatus   *int
	ResponseBody *string
}

type NotificationLarkWebhookDeliveryMarkUnknownResponse struct{}

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
