package usecase

import (
	"common/api/gen/common"
	"context"
	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
)

type NotificationTemplateUsecase struct {
	tx                       base.Tx
	notificationTemplateRepo repo.NotificationTemplateRepo
}

func NewNotificationTemplateUsecase(
	tx base.Tx,
	notificationTemplateRepo repo.NotificationTemplateRepo,
) *NotificationTemplateUsecase {
	return &NotificationTemplateUsecase{
		tx:                       tx,
		notificationTemplateRepo: notificationTemplateRepo,
	}
}

func (d *NotificationTemplateUsecase) Add(ctx context.Context, tpl *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	var update *model.NotificationTemplate
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		update, err = d.notificationTemplateRepo.Save(ctx, tpl)
		if err != nil {
			return err
		}
		return nil
	})
	return update, err
}

func (d *NotificationTemplateUsecase) Update(ctx context.Context, tpl *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	var update *model.NotificationTemplate
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		update, err = d.notificationTemplateRepo.Update(ctx, tpl)
		if err != nil {
			return err
		}
		return nil
	})
	return update, err
}

func (d *NotificationTemplateUsecase) GetMap(ctx context.Context, req *repo.NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error) {
	var result map[string]*model.NotificationTemplate
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		result, err = d.notificationTemplateRepo.GetMap(ctx, req)
		return err
	})
	return result, err
}

func (d *NotificationTemplateUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *common.PageReply, error) {
	var (
		rows      []*model.NotificationTemplate
		pageReply *common.PageReply
	)
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		rows, pageReply, err = d.notificationTemplateRepo.GetPage(ctx, page, req)
		return err
	})
	return rows, pageReply, err
}
