package usecase

import (
	"common/api/gen/common"
	utilent "common/pkg/util/ent"
	"context"
	"errors"
	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
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
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		var err error
		update, err = d.notificationTemplateRepo.Save(ctx, c, tpl)
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
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		var err error
		update, err = d.notificationTemplateRepo.Update(ctx, c, tpl)
		if err != nil {
			return err
		}
		return nil
	})
	return update, err
}

func (d *NotificationTemplateUsecase) GetMap(ctx context.Context, req *repo.NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error) {
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, errors.New("no client in context")
	}
	return d.notificationTemplateRepo.GetMap(ctx, c, req)
}

func (d *NotificationTemplateUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *common.PageReply, error) {
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, nil, errors.New("no client in context")
	}
	return d.notificationTemplateRepo.GetPage(ctx, c, page, req)
}
