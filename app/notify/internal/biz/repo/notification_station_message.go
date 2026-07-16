package repo

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"time"
)

type NotificationStationMessageRepo interface {
	Save(ctx context.Context, req *NotificationStationMessageSaveReq) (*NotificationStationMessageSaveResponse, error)
	Get(ctx context.Context, req *NotificationStationMessageGetReq) (*NotificationStationMessageGetResponse, error)
	List(ctx context.Context, req *NotificationStationMessageListReq) (*NotificationStationMessageListResponse, error)
	Map(ctx context.Context, req *NotificationStationMessageMapReq) (*NotificationStationMessageMapResponse, error)
	Count(ctx context.Context, req *NotificationStationMessageCountReq) (*NotificationStationMessageCountResponse, error)
	Page(ctx context.Context, req *NotificationStationMessagePageReq) (*NotificationStationMessagePageResponse, error)
	MarkRead(ctx context.Context, req *NotificationStationMessageMarkReadReq) (*NotificationStationMessageMarkReadResponse, error)
	CountUnread(ctx context.Context, req *NotificationStationMessageCountUnreadReq) (*NotificationStationMessageCountUnreadResponse, error)
}

type NotificationStationMessageSaveReq struct {
	Message *model.NotificationStationMessage
}

type NotificationStationMessageSaveResponse struct {
	Message *model.NotificationStationMessage
}

type NotificationStationMessageGetReq struct {
	Query *NotificationStationMessageQuery
}

type NotificationStationMessageGetResponse struct {
	Item *model.NotificationStationMessage
}

type NotificationStationMessageListReq struct {
	Query *NotificationStationMessageQuery
}

type NotificationStationMessageListResponse struct {
	Rows []*model.NotificationStationMessage
}

type NotificationStationMessageMapReq struct {
	Query *NotificationStationMessageQuery
}

type NotificationStationMessageMapResponse struct {
	Rows map[int64]*model.NotificationStationMessage
}

type NotificationStationMessageCountReq struct {
	Query *NotificationStationMessageQuery
}

type NotificationStationMessageCountResponse struct {
	Count int
}

type NotificationStationMessagePageReq struct {
	Query *NotificationStationMessageQuery
}

type NotificationStationMessagePageResponse struct {
	Rows []*model.NotificationStationMessage
	Page *base.PageResponse
}

type NotificationStationMessageMarkReadReq struct {
	ReceiverID int64
	IDs        []int64
	StartTime  *time.Time
	EndTime    *time.Time
}

type NotificationStationMessageMarkReadResponse struct {
	Count int
}

type NotificationStationMessageCountUnreadReq struct {
	ReceiverID int64
}

type NotificationStationMessageCountUnreadResponse struct {
	Count int
}

type NotificationStationMessageQuery struct {
	Page       *base.PageRequest
	ID         *int64
	IDs        []int64
	EventIDs   []string
	ReceiverID *int64
	EventType  *enums.EventType
	Unread     *bool
}
