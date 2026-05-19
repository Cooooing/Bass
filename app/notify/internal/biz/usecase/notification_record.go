package usecase

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"time"
)

type NotificationRecordUsecase struct {
	db                     *gen.Client
	notificationRecordRepo repo.NotificationRecordRepo
}

func NewNotificationRecordUsecase(
	db *gen.Client,
	notificationRecordRepo repo.NotificationRecordRepo,
) *NotificationRecordUsecase {
	return &NotificationRecordUsecase{
		db:                     db,
		notificationRecordRepo: notificationRecordRepo,
	}
}

func (d *NotificationRecordUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *common.PageReply, error) {
	return d.notificationRecordRepo.GetPage(ctx, d.db, page, req)
}

func (d *NotificationRecordUsecase) Read(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	return d.notificationRecordRepo.Read(ctx, d.db, receiverId, startTime, endTime, notificationRecordIds)
}
