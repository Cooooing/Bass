package repo

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"time"
)

type NotificationStationMessageRepo interface {
	Save(ctx context.Context, message *model.NotificationStationMessage) (*model.NotificationStationMessage, error)
	Get(ctx context.Context, query *NotificationStationMessageQuery) (*model.NotificationStationMessage, error)
	List(ctx context.Context, query *NotificationStationMessageQuery) ([]*model.NotificationStationMessage, error)
	Map(ctx context.Context, query *NotificationStationMessageQuery) (map[int64]*model.NotificationStationMessage, error)
	Count(ctx context.Context, query *NotificationStationMessageQuery) (int, error)
	Page(ctx context.Context, query *NotificationStationMessageQuery) (*NotificationStationMessagePageResp, error)
	MarkRead(ctx context.Context, req *NotificationStationMessageMarkReadReq) (int, error)
	CountUnread(ctx context.Context, receiverID int64) (int, error)
}

type NotificationStationMessagePageResp struct {
	Rows []*model.NotificationStationMessage
	Page *base.PageResp
}

type NotificationStationMessageMarkReadReq struct {
	ReceiverID int64
	IDs        []int64
	StartTime  *time.Time
	EndTime    *time.Time
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
