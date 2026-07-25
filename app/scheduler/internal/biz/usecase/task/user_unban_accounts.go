package task

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"fmt"
)

type UserUnbanAccounts struct{ consul *commonClient.ConsulClient }

func NewUserUnbanAccounts(
	consul *commonClient.ConsulClient,
) *UserUnbanAccounts {
	return &UserUnbanAccounts{
		consul: consul,
	}
}

func (t *UserUnbanAccounts) Name() string {
	return "user.unban_accounts"
}

func (t *UserUnbanAccounts) Title() string {
	return "用户过期封禁解封"
}

func (t *UserUnbanAccounts) Description() string {
	return "调用 user.UnbanAccounts 解封已到期的临时封禁账号。"
}

type userUnbanAccountsPayload struct {
	UserIDs []int64 `json:"user_ids"`
}

func (t *UserUnbanAccounts) Execute(ctx context.Context, payload string) error {
	var data userUnbanAccountsPayload
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return err
	}
	if len(data.UserIDs) == 0 {
		return fmt.Errorf("invalid user unban payload")
	}
	for _, userID := range data.UserIDs {
		if userID == 0 {
			return fmt.Errorf("invalid user unban payload")
		}
	}
	conn, err := t.consul.GetGrpcConn(constant.UserServiceName.String())
	if err != nil {
		return err
	}
	_, err = userv1.NewAuthServiceClient(conn).UnbanAccounts(ctx, &userv1.UnbanAccounts_Req{
		UserIds: data.UserIDs,
	})
	return err
}
