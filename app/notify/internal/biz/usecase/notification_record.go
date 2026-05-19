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
	"time"
)

type NotificationRecordUsecase struct {
	tx                     base.Tx
	notificationRecordRepo repo.NotificationRecordRepo
}

func NewNotificationRecordUsecase(
	tx base.Tx,
	notificationRecordRepo repo.NotificationRecordRepo,
) *NotificationRecordUsecase {
	return &NotificationRecordUsecase{
		tx:                     tx,
		notificationRecordRepo: notificationRecordRepo,
	}
}

func (d *NotificationRecordUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *common.PageReply, error) {
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, nil, errors.New("no client in context")
	}
	return d.notificationRecordRepo.GetPage(ctx, c, page, req)
}

func (d *NotificationRecordUsecase) Read(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return 0, errors.New("no client in context")
	}
	return d.notificationRecordRepo.Read(ctx, c, receiverId, startTime, endTime, notificationRecordIds)
}
