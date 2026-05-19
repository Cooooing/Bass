package usecase

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
)

type NotificationMetaUsecase struct {
	db                   *gen.Client
	notificationMetaRepo repo.NotificationMetaRepo
}

func NewNotificationMetaUsecase(
	db *gen.Client,
	notificationMetaRepo repo.NotificationMetaRepo,
) *NotificationMetaUsecase {
	return &NotificationMetaUsecase{
		db:                   db,
		notificationMetaRepo: notificationMetaRepo,
	}
}

func (d *NotificationMetaUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *common.PageReply, error) {
	return d.notificationMetaRepo.GetPage(ctx, d.db, page, req)
}
