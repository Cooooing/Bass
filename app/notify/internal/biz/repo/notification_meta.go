package repo

import (
	cv1 "common/api/gen/common/v1"
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/ent/gen"
)

type NotificationMetaRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.NotificationMeta) (*model.NotificationMeta, error)

	GetOne(ctx context.Context, tx *gen.Client, req *NotificationMetaGetReq) (*model.NotificationMeta, error)
	GetList(ctx context.Context, tx *gen.Client, req *NotificationMetaGetReq) ([]*model.NotificationMeta, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *NotificationMetaGetReq) ([]*model.NotificationMeta, *cv1.PageReply, error)
}

type NotificationMetaGetReq struct {
	NotificationMetaId   *int64
	NotificationMeraIds  []int64
	UUID                 *string
	UUIDs                []string
	NotificationMetaType *v1.NotificationType
	SenderId             *int64
	SenderIds            []int64
	Status               *v1.NotificationStatus
}
