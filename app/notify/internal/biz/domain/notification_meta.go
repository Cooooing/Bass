package domain

import (
	cv1 "common/api/common/v1"
	"context"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
)

type NotificationMetaDomain struct {
	*domainbase.BaseDomain
	notificationMetaRepo repo.NotificationMetaRepo
}

func NewNotificationMetaDomain(base *domainbase.BaseDomain, notificationMetaRepo repo.NotificationMetaRepo) *NotificationMetaDomain {
	return &NotificationMetaDomain{
		BaseDomain:           base,
		notificationMetaRepo: notificationMetaRepo,
	}
}

func (d *NotificationMetaDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *cv1.PageReply, error) {
	return d.notificationMetaRepo.GetPage(ctx, d.Db, page, req)
}
