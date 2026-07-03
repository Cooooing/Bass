package service

import (
	"bbs/internal/biz/usecase"
	bbsnotifyv1 "common/proto/gen/bbs/v1/notify"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type NotificationService struct {
	bbsnotifyv1.UnimplementedNotificationServiceServer
	notificationUsecase *usecase.NotificationUsecase
}

func NewNotificationService(notificationUsecase *usecase.NotificationUsecase) *NotificationService {
	return &NotificationService{notificationUsecase: notificationUsecase}
}

func (s *NotificationService) RegisterHttp(hs *http.Server) {
	bbsnotifyv1.RegisterNotificationServiceHTTPServer(hs, s)
}

func (s *NotificationService) List(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error) {
	return s.notificationUsecase.ListNotifications(ctx, req)
}

func (s *NotificationService) MarkRead(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error) {
	return s.notificationUsecase.MarkReadNotification(ctx, req)
}

func (s *NotificationService) CountUnread(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error) {
	return s.notificationUsecase.CountUnreadNotifications(ctx, req)
}
