package usecase

import (
	"common/api/gen/common"
	"context"
	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
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
	var (
		rows      []*model.NotificationRecord
		pageReply *common.PageReply
	)
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		rows, pageReply, err = d.notificationRecordRepo.GetPage(ctx, page, req)
		return err
	})
	return rows, pageReply, err
}

func (d *NotificationRecordUsecase) MarkRead(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	var count int
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		count, err = d.notificationRecordRepo.Read(ctx, receiverId, startTime, endTime, notificationRecordIds)
		return err
	})
	return count, err
}

func (d *NotificationRecordUsecase) CountUnread(ctx context.Context, receiverId int64) (int, error) {
	var count int
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		count, err = d.notificationRecordRepo.UnreadCount(ctx, receiverId)
		return err
	})
	return count, err
}
