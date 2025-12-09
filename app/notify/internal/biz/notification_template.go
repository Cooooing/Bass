package biz

import (
	cv1 "common/api/common/v1"
	"common/pkg/cutil/collections/dict"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"
)

type NotificationTemplateDomain struct {
	*base.BaseDomain
	notificationTemplateRepo repo.NotificationTemplateRepo
}

func NewNotificationTemplateDomain(base *base.BaseDomain, notificationTemplateRepo repo.NotificationTemplateRepo) *NotificationTemplateDomain {
	return &NotificationTemplateDomain{
		BaseDomain:               base,
		notificationTemplateRepo: notificationTemplateRepo,
	}
}

func (d *NotificationTemplateDomain) GetMap(ctx context.Context, tx *gen.Client, req *repo.NotificationTemplateGetReq) (dict.Map[string, *model.NotificationTemplate], error) {
	return d.notificationTemplateRepo.GetMap(ctx, tx, req)
}

func (d *NotificationTemplateDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *cv1.PageReply, error) {
	return d.notificationTemplateRepo.GetPage(ctx, d.Db, page, req)
}
