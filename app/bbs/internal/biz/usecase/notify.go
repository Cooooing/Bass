package usecase

import (
	"bbs/internal/biz/repo"
	bbsnotifyv1 "common/api/gen/bbs/v1/notify"
	"context"
)

type NotifyUsecase struct {
	notifyRepo repo.NotifyRepo
}

func NewNotifyUsecase(notifyRepo repo.NotifyRepo) *NotifyUsecase {
	return &NotifyUsecase{notifyRepo: notifyRepo}
}

func (u *NotifyUsecase) ListNotifications(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error) {
	return u.notifyRepo.ListNotifications(ctx, req)
}

func (u *NotifyUsecase) MarkReadNotification(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error) {
	return u.notifyRepo.MarkReadNotification(ctx, req)
}

func (u *NotifyUsecase) CountUnreadNotifications(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error) {
	return u.notifyRepo.CountUnreadNotifications(ctx, req)
}
