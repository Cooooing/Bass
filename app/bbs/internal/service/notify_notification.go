package service

import (
	"bbs/internal/biz/usecase"
	bbsnotifyv1 "common/proto/gen/bbs/v1/notify"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type NotificationService struct {
	bbsnotifyv1.UnimplementedNotificationServiceServer
	notificationUsecase *usecase.NotificationUsecase
}

func NewNotificationService(notificationUsecase *usecase.NotificationUsecase) *NotificationService {
	return &NotificationService{notificationUsecase: notificationUsecase}
}

func (s *NotificationService) RegisterGrpc(gs *grpc.Server) {}

func (s *NotificationService) RegisterHttp(hs *http.Server) {
	bbsnotifyv1.RegisterNotificationServiceHTTPServer(hs, s)
}

func (s *NotificationService) List(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.notificationUsecase.ListNotifications(ctx, &usecase.ListNotificationsReq{UserID: userID, Page: req.GetPage()})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.ListNotifications_Response{Page: response.Page, Rows: response.Rows}, nil
}
func (s *NotificationService) MarkRead(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.notificationUsecase.MarkReadNotification(ctx, &usecase.MarkReadNotificationReq{UserID: userID, IDs: req.GetIds()})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.MarkReadNotification_Response{Count: response.Count}, nil
}
func (s *NotificationService) CountUnread(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.notificationUsecase.CountUnreadNotifications(ctx, &usecase.CountUnreadNotificationsReq{UserID: userID})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.CountUnreadNotifications_Response{Count: response.Count}, nil
}
