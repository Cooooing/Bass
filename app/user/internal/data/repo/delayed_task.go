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

type unbanExpiredPayload struct {
	UserID      int64 `json:"user_id"`
	BanRecordID int64 `json:"ban_record_id"`
}

func (c *DelayedTaskClient) RegisterUnbanExpired(
	ctx context.Context,
	userID int64,
	banRecordID int64,
	executeAt time.Time,
) error {
	payload, err := json.Marshal(&unbanExpiredPayload{
		UserID:      userID,
		BanRecordID: banRecordID,
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
		IdempotencyKey: fmt.Sprintf("user.unban_expired:%d", banRecordID),
		TaskName:       "user.unban_expired",
		Payload:        string(payload),
		ExecuteAt:      timestamppb.New(executeAt),
		MaxAttempts:    10,
		TimeoutSeconds: 30,
	})
	return err
}
