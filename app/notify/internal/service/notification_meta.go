package service

import (
	v1 "common/api/notify/v1"
	"context"
	"notify/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NotificationMetaService struct {
	v1.UnimplementedNotifyNotificationMetaServiceServer
	*BaseService
	notificationMetaDomain *domain.NotificationMetaDomain
}

func NewNotificationMetaService(baseService *BaseService, notificationMetaDomain *domain.NotificationMetaDomain) *NotificationMetaService {
	return &NotificationMetaService{
		BaseService:            baseService,
		notificationMetaDomain: notificationMetaDomain,
	}
}

func (s *NotificationMetaService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyNotificationMetaServiceServer(gs, s)
}

func (s *NotificationMetaService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyNotificationMetaServiceHTTPServer(hs, s)
}

func (s *NotificationMetaService) Page(ctx context.Context, req *v1.NotificationMetaPageRequest) (rsp *v1.NotificationMetaPageReply, err error) {

	return &v1.NotificationMetaPageReply{}, nil
}
