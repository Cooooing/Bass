package task

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	schedulerv1 "common/proto/gen/scheduler/v1"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"
	"time"

	schedulerenum "scheduler/internal/enum"
)

type UserUnbanAccounts struct {
	userClient  *rpc.UserClient
	title       string
	description string
}

func NewUserUnbanAccounts(
	userClient *rpc.UserClient,
) *UserUnbanAccounts {
	return &UserUnbanAccounts{
		userClient:  userClient,
		title:       "用户批量解封",
		description: "调用 user.UnbanAccounts 解封过期临时封禁账号。",
	}
}

func (t *UserUnbanAccounts) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameUserUnbanAccounts
}

func (t *UserUnbanAccounts) Title() string {
	return t.title
}

func (t *UserUnbanAccounts) Description() string {
	return t.description
}

func (t *UserUnbanAccounts) DefaultScheduledTasks() []*DefaultScheduledTask {
	return nil
}

func (t *UserUnbanAccounts) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{
		{
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			Timeout:       30 * time.Second,
			MaxAttempts:   3,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll,
		},
	}
}

func (t *UserUnbanAccounts) Execute(ctx context.Context, execution Execution) (*Result, error) {
	data := &schedulerv1.UnbanAccountsTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(execution.Payload, data); err != nil {
		return nil, err
	}
	if len(data.GetUserIds()) == 0 {
		return nil, fmt.Errorf("invalid user unban payload")
	}
	response, err := t.userClient.Auth.UnbanAccounts(ctx, &userv1.UnbanAccounts_Req{
		UserIds: data.GetUserIds(),
	})
	if err != nil || response == nil {
		if err == nil {
			err = fmt.Errorf("user unban returned an empty response")
		}
		return nil, err
	}
	return &Result{Summary: fmt.Sprintf("user_count=%d", len(data.GetUserIds()))}, nil
}
