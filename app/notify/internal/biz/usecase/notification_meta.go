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
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, nil, errors.New("no client in context")
	}
	return d.notificationMetaRepo.GetPage(ctx, c, page, req)
}
