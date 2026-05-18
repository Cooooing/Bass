package repo

import (
	commonv1 "common/api/gen/common"
	"common/api/gen/common/enums"
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/gen"
	"time"
)

type NotificationRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.NotificationRecord) (*model.NotificationRecord, error)
	Saves(ctx context.Context, tx *gen.Client, u []*model.NotificationRecord) ([]*model.NotificationRecord, error)

	Read(ctx context.Context, tx *gen.Client, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error)

	GetOne(ctx context.Context, tx *gen.Client, req *NotificationRecordGetReq) (*model.NotificationRecord, error)
	GetList(ctx context.Context, tx *gen.Client, req *NotificationRecordGetReq) ([]*model.NotificationRecord, error)
	GetPage(ctx context.Context, tx *gen.Client, page *commonv1.PageRequest, req *NotificationRecordGetReq) ([]*model.NotificationRecord, *commonv1.PageReply, error)
}

type NotificationRecordGetReq struct {
	EventType             *enums.EventType
	NotificationRecordId  *int64
	NotificationRecordIds []int64
	NotificationMetaId    *int64
	NotificationMetaIds   []int64
	ReceiverId            *int64
	Status                *v1.NotificationStatus

	WithMeta bool
}
