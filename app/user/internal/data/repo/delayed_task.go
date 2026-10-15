package repo

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	schedulerv1 "common/proto/gen/scheduler/v1"
	schedulerv1enum "common/proto/gen/scheduler/v1/enum"
	"context"
	"errors"
	"time"

	bizrepo "user/internal/biz/repo"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type DelayedTaskClient struct{ schedulerClient *rpc.SchedulerClient }

func NewDelayedTaskClient(
	schedulerClient *rpc.SchedulerClient,
) bizrepo.DelayedTaskClient {
	return &DelayedTaskClient{
		schedulerClient: schedulerClient,
	}
}

func (c *DelayedTaskClient) RegisterUnbanAccounts(ctx context.Context, userID int64, executeAt *time.Time) error {
	if executeAt == nil {
		return errors.New("delayed task execute_at is required")
	}
	payload, err := commonscheduler.MarshalPayload(&schedulerv1.UnbanAccountsTaskPayload{UserIds: []int64{userID}})
	if err != nil {
		return err
	}
	_, err = c.schedulerClient.DelayedTask.Schedule(ctx, &schedulerv1.ScheduleSchedulerDelayedTask_Req{
		HandlerName: schedulerv1enum.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_USER_UNBAN_ACCOUNTS,
		Payload:     payload,
		ScheduledAt: timestamppb.New(*executeAt),
	})
	return err
}
