package domain

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
)

type NotificationMetaDomain struct {
	db                   *gen.Client
	notificationMetaRepo repo.NotificationMetaRepo
}

func NewNotificationMetaDomain(
	db *gen.Client,
	notificationMetaRepo repo.NotificationMetaRepo,
) *NotificationMetaDomain {
	return &NotificationMetaDomain{
		db:                   db,
		notificationMetaRepo: notificationMetaRepo,
	}
}

func (d *NotificationMetaDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *common.PageReply, error) {
	return d.notificationMetaRepo.GetPage(ctx, d.db, page, req)
}
