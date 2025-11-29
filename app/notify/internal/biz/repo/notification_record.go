package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/notify/v1"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/ent/gen"
)

type NotificationRecordRepo interface {
	Save(ctx context.Context, client *gen.Client, u *model.NotificationRecord) (*model.NotificationRecord, error)
	Saves(ctx context.Context, client *gen.Client, u []*model.NotificationRecord) ([]*model.NotificationRecord, error)

	GetOne(ctx context.Context, client *gen.Client, req *NotificationRecordGetReq) (*model.NotificationRecord, error)
	GetList(ctx context.Context, client *gen.Client, req *NotificationRecordGetReq) ([]*model.NotificationRecord, error)
	GetPage(ctx context.Context, client *gen.Client, page *cv1.PageRequest, req *NotificationRecordGetReq) ([]*model.NotificationRecord, *cv1.PageReply, error)
}

type NotificationRecordGetReq struct {
	NotificationType      *v1.NotificationMetaType
	NotificationRecordId  *int64
	NotificationRecordIds []int64
	NotificationMetaId    *int64
	NotificationMetaIds   []int64
	SenderId              *int64
	ReceiverId            *int64
	Status                *v1.NotificationMetaStatus

	WithMeta bool
}
