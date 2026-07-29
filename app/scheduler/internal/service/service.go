package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewSchedulerScheduledTaskService,
	NewSchedulerDelayedTaskService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	schedulerScheduledTaskService *SchedulerScheduledTaskService,
	schedulerDelayedTaskService *SchedulerDelayedTaskService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		schedulerScheduledTaskService,
		schedulerDelayedTaskService,
	}
}
