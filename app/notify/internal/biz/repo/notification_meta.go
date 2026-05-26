package repo

import (
	commonv1 "common/api/gen/common"
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/model"
)

type NotificationMetaRepo interface {
	Save(ctx context.Context, u *model.NotificationMeta) (*model.NotificationMeta, error)

	Get(ctx context.Context, req *NotificationMetaGetReq) (*model.NotificationMeta, error)
	GetList(ctx context.Context, req *NotificationMetaGetReq) ([]*model.NotificationMeta, error)
	GetPage(ctx context.Context, page *commonv1.PageRequest, req *NotificationMetaGetReq) ([]*model.NotificationMeta, *commonv1.PageReply, error)
}

type NotificationMetaGetReq struct {
	NotificationMetaId  *int64
	NotificationMetaIds []int64
	Status              *v1.NotificationStatus
}
