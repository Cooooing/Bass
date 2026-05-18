package repo

import (
	commonv1 "common/api/gen/common"
	"common/api/gen/common/enums"
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/gen"
)

type NotificationMetaRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.NotificationMeta) (*model.NotificationMeta, error)

	GetOne(ctx context.Context, tx *gen.Client, req *NotificationMetaGetReq) (*model.NotificationMeta, error)
	GetList(ctx context.Context, tx *gen.Client, req *NotificationMetaGetReq) ([]*model.NotificationMeta, error)
	GetPage(ctx context.Context, tx *gen.Client, page *commonv1.PageRequest, req *NotificationMetaGetReq) ([]*model.NotificationMeta, *commonv1.PageReply, error)
}

type NotificationMetaGetReq struct {
	NotificationMetaId  *int64
	NotificationMeraIds []int64
	UUID                *string
	UUIDs               []string
	EventType           *enums.EventType
	Status              *v1.NotificationStatus
}
