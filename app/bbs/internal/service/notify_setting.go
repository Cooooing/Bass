package service

import (
	"bbs/internal/biz/usecase"
	bbsnotifyv1 "common/api/gen/bbs/v1/notify"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NotificationSettingService struct {
	bbsnotifyv1.UnimplementedNotificationSettingServiceServer
	notifyUsecase *usecase.NotifyUsecase
}

func NewNotificationSettingService(notifyUsecase *usecase.NotifyUsecase) *NotificationSettingService {
	return &NotificationSettingService{notifyUsecase: notifyUsecase}
}

func (s *NotificationSettingService) RegisterGrpc(gs *grpc.Server) {
	bbsnotifyv1.RegisterNotificationSettingServiceServer(gs, s)
}

func (s *NotificationSettingService) RegisterHttp(hs *http.Server) {
	bbsnotifyv1.RegisterNotificationSettingServiceHTTPServer(hs, s)
}

func (s *NotificationSettingService) ListCurrent(ctx context.Context, req *bbsnotifyv1.ListCurrentNotificationSetting_Request) (*bbsnotifyv1.ListCurrentNotificationSetting_Reply, error) {
	return s.notifyUsecase.ListCurrentNotificationSetting(ctx, req)
}

func (s *NotificationSettingService) UpdateCurrent(ctx context.Context, req *bbsnotifyv1.UpdateCurrentNotificationSetting_Request) (*bbsnotifyv1.UpdateCurrentNotificationSetting_Reply, error) {
	return s.notifyUsecase.UpdateCurrentNotificationSetting(ctx, req)
}
