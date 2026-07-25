package rpc

import (
	schedulerv1 "common/proto/gen/scheduler/v1"

	"google.golang.org/grpc"
)

type SchedulerClient struct {
	DelayedTask schedulerv1.SchedulerDelayedTaskServiceClient
}

func NewSchedulerClient(
	conn *grpc.ClientConn,
) *SchedulerClient {
	return &SchedulerClient{
		DelayedTask: schedulerv1.NewSchedulerDelayedTaskServiceClient(conn),
	}
}
