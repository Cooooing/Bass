package service

import (
	"bbs/internal/biz/usecase"
	bbsnotifyv1 "common/api/gen/bbs/v1/notify"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NotificationService struct {
	bbsnotifyv1.UnimplementedNotificationServiceServer
	notifyUsecase *usecase.NotifyUsecase
}

func NewNotificationService(notifyUsecase *usecase.NotifyUsecase) *NotificationService {
	return &NotificationService{notifyUsecase: notifyUsecase}
}

func (s *NotificationService) RegisterGrpc(gs *grpc.Server) {
	bbsnotifyv1.RegisterNotificationServiceServer(gs, s)
}

func (s *NotificationService) RegisterHttp(hs *http.Server) {
	bbsnotifyv1.RegisterNotificationServiceHTTPServer(hs, s)
}

func (s *NotificationService) List(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error) {
	return s.notifyUsecase.ListNotifications(ctx, req)
}

func (s *NotificationService) MarkRead(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error) {
	return s.notifyUsecase.MarkReadNotification(ctx, req)
}

func (s *NotificationService) CountUnread(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error) {
	return s.notifyUsecase.CountUnreadNotifications(ctx, req)
}
