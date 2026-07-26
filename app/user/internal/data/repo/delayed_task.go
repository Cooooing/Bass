package repo

import (
	"common/pkg/client/rpc"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	bizrepo "user/internal/biz/repo"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ bizrepo.DelayedTaskClient = (*DelayedTaskClient)(nil)

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

func (c *DelayedTaskClient) RegisterUnbanAccounts(ctx context.Context, userID int64, banRecordID int64, executeAt *time.Time) error {
	if executeAt == nil {
		return errors.New("delayed task execute_at is required")
	}
	payload, err := json.Marshal(&unbanAccountsPayload{
		UserIDs: []int64{userID},
	})
	if err != nil {
		return err
	}
	_, err = c.schedulerClient.DelayedTask.Register(ctx, &schedulerv1.RegisterSchedulerDelayedTask_Req{
		IdempotencyKey: fmt.Sprintf("user.unban_accounts:%d", banRecordID),
		TaskName:       "user.unban_accounts",
		Payload:        string(payload),
		ExecuteAt:      timestamppb.New(*executeAt),
		MaxAttempts:    10,
		TimeoutSeconds: 30,
	})
	return err
}
