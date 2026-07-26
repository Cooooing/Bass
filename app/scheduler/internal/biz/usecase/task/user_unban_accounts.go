package task

import (
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type UserUnbanAccounts struct {
	userClient  *rpc.UserClient
	name        string
	title       string
	description string
}

func NewUserUnbanAccounts(
	userClient *rpc.UserClient,
) *UserUnbanAccounts {
	return &UserUnbanAccounts{
		userClient:  userClient,
		name:        "user.unban_accounts",
		title:       "User unban accounts",
		description: "Call user.UnbanAccounts to unban expired temporary bans.",
	}
}

func (t *UserUnbanAccounts) Name() string {
	return t.name
}

func (t *UserUnbanAccounts) Title() string {
	return t.title
}

func (t *UserUnbanAccounts) Description() string {
	return t.description
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
