package usecase

import (
	"bbs/internal/biz/repo"
	bbsnotifyv1 "common/api/gen/bbs/v1/notify"
	"context"
)

type NotificationUsecase struct {
	notificationRepo repo.NotificationRepo
}

func NewNotificationUsecase(notificationRepo repo.NotificationRepo) *NotificationUsecase {
	return &NotificationUsecase{notificationRepo: notificationRepo}
}

func (u *NotificationUsecase) ListNotifications(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error) {
	return u.notificationRepo.ListNotifications(ctx, req)
}

func (u *NotificationUsecase) MarkReadNotification(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error) {
	return u.notificationRepo.MarkReadNotification(ctx, req)
}

func (u *NotificationUsecase) CountUnreadNotifications(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error) {
	return u.notificationRepo.CountUnreadNotifications(ctx, req)
}
