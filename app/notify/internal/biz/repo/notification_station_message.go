package repo

import (
	"common/proto/gen/common"
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"time"
)

type NotificationStationMessageRepo interface {
	Save(ctx context.Context, message *model.NotificationStationMessage) (*model.NotificationStationMessage, error)
	Get(ctx context.Context, req *NotificationStationMessageQuery) (*model.NotificationStationMessage, error)
	List(ctx context.Context, req *NotificationStationMessageQuery) ([]*model.NotificationStationMessage, error)
	Map(ctx context.Context, req *NotificationStationMessageQuery) (map[int64]*model.NotificationStationMessage, error)
	Count(ctx context.Context, req *NotificationStationMessageQuery) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *NotificationStationMessageQuery) ([]*model.NotificationStationMessage, *common.PageReply, error)
	MarkRead(ctx context.Context, receiverID int64, ids []int64, startTime *time.Time, endTime *time.Time) (int, error)
	CountUnread(ctx context.Context, receiverID int64) (int, error)
}

type NotificationStationMessageQuery struct {
	ID         *int64
	IDs        []int64
	EventIDs   []string
	ReceiverID *int64
	EventType  *enums.EventType
	Unread     *bool
}
