package repo

import (
	"common/api/gen/common"
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"time"
)

type NotificationStationMessageRepo interface {
	Save(ctx context.Context, message *model.NotificationStationMessage) (*model.NotificationStationMessage, error)
	Page(ctx context.Context, page *common.PageRequest, req *NotificationStationMessageQuery) ([]*model.NotificationStationMessage, *common.PageReply, error)
	MarkRead(ctx context.Context, receiverID int64, ids []int64, startTime *time.Time, endTime *time.Time) (int, error)
	CountUnread(ctx context.Context, receiverID int64) (int, error)
}

type NotificationStationMessageQuery struct {
	IDs        []int64
	ReceiverID *int64
	EventType  *enums.EventType
	Unread     *bool
}
