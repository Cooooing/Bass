package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/notify/v1"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/ent/gen"
	"time"
)

type NotificationRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.NotificationRecord) (*model.NotificationRecord, error)
	Saves(ctx context.Context, tx *gen.Client, u []*model.NotificationRecord) ([]*model.NotificationRecord, error)

	Read(ctx context.Context, tx *gen.Client, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error)

	GetOne(ctx context.Context, tx *gen.Client, req *NotificationRecordGetReq) (*model.NotificationRecord, error)
	GetList(ctx context.Context, tx *gen.Client, req *NotificationRecordGetReq) ([]*model.NotificationRecord, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *NotificationRecordGetReq) ([]*model.NotificationRecord, *cv1.PageReply, error)
}

type NotificationRecordGetReq struct {
	NotificationType      *v1.NotificationType
	NotificationRecordId  *int64
	NotificationRecordIds []int64
	NotificationMetaId    *int64
	NotificationMetaIds   []int64
	SenderId              *int64
	ReceiverId            *int64
	Status                *v1.NotificationStatus

	WithMeta bool
}
