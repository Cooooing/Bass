package domain

import (
	"common/api/gen/common"
	"context"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"time"
)

type NotificationRecordDomain struct {
	*domainbase.BaseDomain
	notificationRecordRepo repo.NotificationRecordRepo
}

func NewNotificationRecordDomain(
	base *domainbase.BaseDomain,
	notificationRecordRepo repo.NotificationRecordRepo) *NotificationRecordDomain {
	return &NotificationRecordDomain{
		BaseDomain:             base,
		notificationRecordRepo: notificationRecordRepo,
	}
}

func (d *NotificationRecordDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *common.PageReply, error) {
	return d.notificationRecordRepo.GetPage(ctx, d.Db, page, req)
}

func (d *NotificationRecordDomain) Read(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	return d.notificationRecordRepo.Read(ctx, d.Db, receiverId, startTime, endTime, notificationRecordIds)
}
