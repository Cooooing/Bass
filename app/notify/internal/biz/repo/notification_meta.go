package repo

import (
	cv1 "common/api/common/v1"
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
}
