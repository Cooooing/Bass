package rpc

import (
	schedulerv1 "common/proto/gen/scheduler/v1"

	"google.golang.org/grpc"
)

type SchedulerClient struct {
	ScheduledTask schedulerv1.SchedulerScheduledTaskServiceClient
	DelayedTask   schedulerv1.SchedulerDelayedTaskServiceClient
}

func NewSchedulerClient(
	conn *grpc.ClientConn,
) *SchedulerClient {
	return &SchedulerClient{
		ScheduledTask: schedulerv1.NewSchedulerScheduledTaskServiceClient(conn),
		DelayedTask:   schedulerv1.NewSchedulerDelayedTaskServiceClient(conn),
	}
}
