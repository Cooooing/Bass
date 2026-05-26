package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewNotificationMetaService,
	NewNotificationRecordService,
	NewNotificationSettingService,
	NewNotificationTemplateService,
	NewOssService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	notificationMetaService *NotificationMetaService,
	notificationRecordService *NotificationRecordService,
	notificationSettingService *NotificationSettingService,
	notificationTemplateService *NotificationTemplateService,
	ossService *OssService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		notificationMetaService,
		notificationRecordService,
		notificationSettingService,
		notificationTemplateService,
		ossService,
	}
}
