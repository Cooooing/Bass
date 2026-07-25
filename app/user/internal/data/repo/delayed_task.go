package repo

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"encoding/json"
	"fmt"
	"time"
	bizrepo "user/internal/biz/repo"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ bizrepo.DelayedTaskClient = (*DelayedTaskClient)(nil)

type DelayedTaskClient struct{ consul *commonClient.ConsulClient }

func NewDelayedTaskClient(
	consul *commonClient.ConsulClient,
) bizrepo.DelayedTaskClient {
	return &DelayedTaskClient{
		consul: consul,
	}
}

type unbanAccountsPayload struct {
	UserIDs []int64 `json:"user_ids"`
}

func (c *DelayedTaskClient) RegisterUnbanAccounts(ctx context.Context, userID int64, banRecordID int64, executeAt time.Time) error {
	payload, err := json.Marshal(&unbanAccountsPayload{
		UserIDs: []int64{userID},
	})
	if err != nil {
		return err
	}
	conn, err := c.consul.GetGrpcConn(constant.SchedulerServiceName.String())
	if err != nil {
		return err
	}
	client := schedulerv1.NewSchedulerDelayedTaskServiceClient(conn)
	_, err = client.Register(ctx, &schedulerv1.RegisterSchedulerDelayedTask_Req{
		IdempotencyKey: fmt.Sprintf("user.unban_accounts:%d", banRecordID),
		TaskName:       "user.unban_accounts",
		Payload:        string(payload),
		ExecuteAt:      timestamppb.New(executeAt),
		MaxAttempts:    10,
		TimeoutSeconds: 30,
	})
	return err
}
