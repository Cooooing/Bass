package task

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	schedulerv1 "common/proto/gen/scheduler/v1"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"
	schedulerenum "scheduler/internal/enum"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserSettleLiveness 在每日 01:00 结算前一业务日活跃度并自动签到。
type UserSettleLiveness struct{ client *rpc.UserClient }

func NewUserSettleLiveness(client *rpc.UserClient) *UserSettleLiveness {
	return &UserSettleLiveness{client: client}
}
func (*UserSettleLiveness) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameUserSettleLiveness
}
func (*UserSettleLiveness) Title() string { return "结算每日活跃度" }
func (*UserSettleLiveness) Description() string {
	return "结算前一业务日活跃度，并为满足阈值的用户自动签到。"
}
func (t *UserSettleLiveness) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{{Title: t.Title(), Description: t.Description(), Enabled: true, CronSpec: "0 0 1 * * ?", Payload: commonscheduler.MustMarshalPayload(&schedulerv1.SettleUserLivenessTaskPayload{PageSize: 100}), Timeout: 30 * time.Minute, MaxAttempts: 3, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest, AllowOverlap: false}}
}
func (*UserSettleLiveness) DefaultDelayedTasks() []*DefaultDelayedTask { return nil }
func (t *UserSettleLiveness) Execute(ctx context.Context, execution Execution) error {
	data := &schedulerv1.SettleUserLivenessTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(execution.Payload, data); err != nil {
		return err
	}
	if execution.ScheduledAt.IsZero() {
		return fmt.Errorf("scheduled liveness settlement requires scheduled_at")
	}
	date := execution.ScheduledAt.In(time.Local).AddDate(0, 0, -1)
	_, err := t.client.Liveness.SettleLiveness(ctx, &userv1.SettleUserLiveness_Req{Date: timestamppb.New(date), PageSize: data.GetPageSize()})
	return err
}
