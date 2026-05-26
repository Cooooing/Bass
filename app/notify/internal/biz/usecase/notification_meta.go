package usecase

import (
	"common/api/gen/common"
	"context"
	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
)

type NotificationMetaUsecase struct {
	tx                   base.Tx
	notificationMetaRepo repo.NotificationMetaRepo
}

func NewNotificationMetaUsecase(
	tx base.Tx,
	notificationMetaRepo repo.NotificationMetaRepo,
) *NotificationMetaUsecase {
	return &NotificationMetaUsecase{
		tx:                   tx,
		notificationMetaRepo: notificationMetaRepo,
	}
}

func (d *NotificationMetaUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *common.PageReply, error) {
	var (
		rows      []*model.NotificationMeta
		pageReply *common.PageReply
	)
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		rows, pageReply, err = d.notificationMetaRepo.GetPage(ctx, page, req)
		return err
	})
	return rows, pageReply, err
}
