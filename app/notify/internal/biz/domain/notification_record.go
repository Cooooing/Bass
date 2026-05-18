package domain

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"time"
)

type NotificationRecordDomain struct {
	db                     *gen.Client
	notificationRecordRepo repo.NotificationRecordRepo
}

func NewNotificationRecordDomain(
	db *gen.Client,
	notificationRecordRepo repo.NotificationRecordRepo,
) *NotificationRecordDomain {
	return &NotificationRecordDomain{
		db:                     db,
		notificationRecordRepo: notificationRecordRepo,
	}
}

func (d *NotificationRecordDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *common.PageReply, error) {
	return d.notificationRecordRepo.GetPage(ctx, d.db, page, req)
}

func (d *NotificationRecordDomain) Read(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	return d.notificationRecordRepo.Read(ctx, d.db, receiverId, startTime, endTime, notificationRecordIds)
}
