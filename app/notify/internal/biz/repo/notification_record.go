package repo

import (
	commonv1 "common/api/gen/common"
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/model"
	"time"
)

type NotificationRecordRepo interface {
	Save(ctx context.Context, u *model.NotificationRecord) (*model.NotificationRecord, error)
	Saves(ctx context.Context, u []*model.NotificationRecord) ([]*model.NotificationRecord, error)

	Read(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error)
	UnreadCount(ctx context.Context, receiverId int64) (int, error)

	Get(ctx context.Context, req *NotificationRecordGetReq) (*model.NotificationRecord, error)
	GetList(ctx context.Context, req *NotificationRecordGetReq) ([]*model.NotificationRecord, error)
	GetPage(ctx context.Context, page *commonv1.PageRequest, req *NotificationRecordGetReq) ([]*model.NotificationRecord, *commonv1.PageReply, error)
}

type NotificationRecordGetReq struct {
	NotificationRecordId  *int64
	NotificationRecordIds []int64
	NotificationMetaId    *int64
	NotificationMetaIds   []int64
	ReceiverId            *int64
	Status                *v1.NotificationStatus

	WithMeta bool
}
