package domain

import (
	cv1 "common/api/common/v1"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"time"
)

type NotificationRecordDomain struct {
	*base.BaseDomain
	notificationRecordRepo repo.NotificationRecordRepo
}

func NewNotificationRecordDomain(base *base.BaseDomain, notificationRecordRepo repo.NotificationRecordRepo) *NotificationRecordDomain {
	return &NotificationRecordDomain{
		BaseDomain:             base,
		notificationRecordRepo: notificationRecordRepo,
	}
}

func (d *NotificationRecordDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *cv1.PageReply, error) {
	return d.notificationRecordRepo.GetPage(ctx, d.Db, page, req)
}

func (d *NotificationRecordDomain) Read(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	return d.notificationRecordRepo.Read(ctx, d.Db, receiverId, startTime, endTime, notificationRecordIds)
}
