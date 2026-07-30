package task

import (
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"fmt"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"
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
			TaskKey:       schedulerenum.TaskKeyUserUnbanAccountsDefault,
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			Timeout:       30 * time.Second,
			MaxAttempts:   3,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll,
		},
	}
}

type userUnbanAccountsPayload struct {
	UserIDs []int64 `json:"user_ids"`
}

func (t *UserUnbanAccounts) Execute(ctx context.Context, payload string) error {
	var data userUnbanAccountsPayload
	if err := json.Unmarshal([]byte(strings.TrimSpace(payload)), &data); err != nil {
		return err
	}
	if len(data.UserIDs) == 0 {
		return fmt.Errorf("invalid user unban payload")
	}
	_, err := t.userClient.Auth.UnbanAccounts(ctx, &userv1.UnbanAccounts_Req{
		UserIds: data.UserIDs,
	})
	return err
}
