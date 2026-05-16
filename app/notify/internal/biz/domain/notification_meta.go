package domain

import (
	"common/api/gen/common"
	"context"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
)

type NotificationMetaDomain struct {
	*domainbase.BaseDomain
	notificationMetaRepo repo.NotificationMetaRepo
}

func NewNotificationMetaDomain(
	base *domainbase.BaseDomain,
	notificationMetaRepo repo.NotificationMetaRepo) *NotificationMetaDomain {
	return &NotificationMetaDomain{
		BaseDomain:           base,
		notificationMetaRepo: notificationMetaRepo,
	}
}

func (d *NotificationMetaDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *common.PageReply, error) {
	return d.notificationMetaRepo.GetPage(ctx, d.Db, page, req)
}
