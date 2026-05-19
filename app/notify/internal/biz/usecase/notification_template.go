package usecase

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/client"
	"notify/internal/data/gen"
)

type NotificationTemplateUsecase struct {
	db                       *gen.Client
	notificationTemplateRepo repo.NotificationTemplateRepo
}

func NewNotificationTemplateUsecase(
	db *gen.Client,
	notificationTemplateRepo repo.NotificationTemplateRepo,
) *NotificationTemplateUsecase {
	return &NotificationTemplateUsecase{
		db:                       db,
		notificationTemplateRepo: notificationTemplateRepo,
	}
}

func (d *NotificationTemplateUsecase) Add(ctx context.Context, tpl *model.NotificationTemplate) (*model.NotificationTemplate, error) {
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

func (d *NotificationTemplateUsecase) Update(ctx context.Context, tpl *model.NotificationTemplate) (*model.NotificationTemplate, error) {
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

func (d *NotificationTemplateUsecase) GetMap(ctx context.Context, req *repo.NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error) {
	return d.notificationTemplateRepo.GetMap(ctx, d.db, req)
}

func (d *NotificationTemplateUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *common.PageReply, error) {
	return d.notificationTemplateRepo.GetPage(ctx, d.db, page, req)
}
