package service

import (
	commonClient "common/pkg/client"
	"common/pkg/util/server"
	"notify/internal/conf"
	"notify/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewNotificationMetaService,
	NewNotificationRecordService,
	NewNotificationTemplateService,
	NewOssService,
	ProvideServices,
)

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(gs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	notificationMetaService *NotificationMetaService,
	notificationRecordService *NotificationRecordService,
	notificationTemplateService *NotificationTemplateService,
	ossService *OssService,
) []server.Service {
	return []server.Service{
		systemService,
		notificationMetaService,
		notificationRecordService,
		notificationTemplateService,
		ossService,
	}
}
