package domain

import (
	cv1 "common/api/common/v1"
	"common/pkg/cutil/collections/dict"
	"context"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent"
	"notify/internal/data/ent/gen"
)

type NotificationTemplateDomain struct {
	*domainbase.BaseDomain
	notificationTemplateRepo repo.NotificationTemplateRepo
}

func NewNotificationTemplateDomain(base *domainbase.BaseDomain, notificationTemplateRepo repo.NotificationTemplateRepo) *NotificationTemplateDomain {
	return &NotificationTemplateDomain{
		BaseDomain:               base,
		notificationTemplateRepo: notificationTemplateRepo,
	}
}

func (d *NotificationTemplateDomain) Add(ctx context.Context, tpl *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	var update *model.NotificationTemplate
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		var err error
		update, err = d.notificationTemplateRepo.Save(ctx, tx, tpl)
		if err != nil {
			return err
		}
		return nil
	})
	return update, err
}

func (d *NotificationTemplateDomain) Update(ctx context.Context, tpl *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	var update *model.NotificationTemplate
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		var err error
		update, err = d.notificationTemplateRepo.Update(ctx, tx, tpl)
		if err != nil {
			return err
		}
		return nil
	})
	return update, err
}

func (d *NotificationTemplateDomain) GetMap(ctx context.Context, req *repo.NotificationTemplateGetReq) (dict.Map[string, *model.NotificationTemplate], error) {
	return d.notificationTemplateRepo.GetMap(ctx, d.Db, req)
}

func (d *NotificationTemplateDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *cv1.PageReply, error) {
	return d.notificationTemplateRepo.GetPage(ctx, d.Db, page, req)
}
