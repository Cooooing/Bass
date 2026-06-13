package usecase

import (
	"bbs/internal/biz/repo"
	bbsnotifyv1 "common/proto/gen/bbs/v1/notify"
	"context"
)

type NotificationUsecase struct {
	notificationClient repo.NotificationClient
}

func NewNotificationUsecase(notificationClient repo.NotificationClient) *NotificationUsecase {
	return &NotificationUsecase{notificationClient: notificationClient}
}

func (u *NotificationUsecase) ListNotifications(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error) {
	return u.notificationClient.ListNotifications(ctx, req)
}

func (u *NotificationUsecase) MarkReadNotification(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error) {
	return u.notificationClient.MarkReadNotification(ctx, req)
}

func (u *NotificationUsecase) CountUnreadNotifications(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error) {
	return u.notificationClient.CountUnreadNotifications(ctx, req)
}
