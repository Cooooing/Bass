package repo

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	schedulerv1 "common/proto/gen/scheduler/v1"
	schedulerv1enum "common/proto/gen/scheduler/v1/enum"
	"context"
	"encoding/json"
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

type unbanAccountsPayload struct {
	UserIDs []int64 `json:"user_ids"`
}

func (c *DelayedTaskClient) RegisterUnbanAccounts(ctx context.Context, userID int64, executeAt *time.Time) error {
	if executeAt == nil {
		return errors.New("delayed task execute_at is required")
	}
	payload, err := json.Marshal(&unbanAccountsPayload{
		UserIDs: []int64{userID},
	})
	if err != nil {
		return err
	}
	_, err = c.schedulerClient.DelayedTask.Schedule(ctx, &schedulerv1.ScheduleSchedulerDelayedTask_Req{
		TaskKey: commonscheduler.TaskKeyMap.MustToEnum(
			schedulerv1enum.SchedulerTaskKey_SCHEDULER_TASK_KEY_USER_UNBAN_ACCOUNTS_DEFAULT,
		).String(),
		Payload:     string(payload),
		ScheduledAt: timestamppb.New(*executeAt),
	})
	return err
}
