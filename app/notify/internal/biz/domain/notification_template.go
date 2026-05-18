package domain

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/client"
	"notify/internal/data/gen"
)

type NotificationTemplateDomain struct {
	db                       *gen.Client
	notificationTemplateRepo repo.NotificationTemplateRepo
}

func NewNotificationTemplateDomain(
	db *gen.Client,
	notificationTemplateRepo repo.NotificationTemplateRepo,
) *NotificationTemplateDomain {
	return &NotificationTemplateDomain{
		db:                       db,
		notificationTemplateRepo: notificationTemplateRepo,
	}
}

func (d *NotificationTemplateDomain) Add(ctx context.Context, tpl *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	var update *model.NotificationTemplate
	err := client.WithTx(ctx, d.db, func(tx *gen.Client) error {
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
	err := client.WithTx(ctx, d.db, func(tx *gen.Client) error {
		var err error
		update, err = d.notificationTemplateRepo.Update(ctx, tx, tpl)
		if err != nil {
			return err
		}
		return nil
	})
	return update, err
}

func (d *NotificationTemplateDomain) GetMap(ctx context.Context, req *repo.NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error) {
	return d.notificationTemplateRepo.GetMap(ctx, d.db, req)
}

func (d *NotificationTemplateDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *common.PageReply, error) {
	return d.notificationTemplateRepo.GetPage(ctx, d.db, page, req)
}
