package service

import (
	"common/pkg/util/server"

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

func ProvideServices(
	systemService *SystemService,
	notificationMetaService *NotificationMetaService,
	notificationRecordService *NotificationRecordService,
	notificationTemplateService *NotificationTemplateService,
	ossService *OssService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		notificationMetaService,
		notificationRecordService,
		notificationTemplateService,
		ossService,
	}
}
