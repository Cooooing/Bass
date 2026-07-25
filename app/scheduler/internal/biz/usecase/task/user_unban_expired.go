package task

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"fmt"
)

const UserUnbanExpiredTaskName = "user.unban_expired"

type UserUnbanExpired struct{ consul *commonClient.ConsulClient }

func NewUserUnbanExpired(
	consul *commonClient.ConsulClient,
) *UserUnbanExpired {
	return &UserUnbanExpired{
		consul: consul,
	}
}

func (t *UserUnbanExpired) Name() string {
	return UserUnbanExpiredTaskName
}

func (t *UserUnbanExpired) Title() string {
	return "用户过期封禁解封"
}

func (t *UserUnbanExpired) Description() string {
	return "调用 user.UnbanExpired 解封已到期的临时封禁账号。"
}

type userUnbanExpiredPayload struct {
	UserID      int64 `json:"user_id"`
	BanRecordID int64 `json:"ban_record_id"`
}

func (t *UserUnbanExpired) Execute(
	ctx context.Context,
	payload string,
) error {
	var data userUnbanExpiredPayload
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return err
	}
	if data.UserID == 0 || data.BanRecordID == 0 {
		return fmt.Errorf("invalid user unban payload")
	}
	conn, err := t.consul.GetGrpcConn(constant.UserServiceName.String())
	if err != nil {
		return err
	}
	_, err = userv1.NewAuthServiceClient(conn).UnbanExpired(ctx, &userv1.UnbanExpired_Req{
		UserId:      data.UserID,
		BanRecordId: data.BanRecordID,
	})
	return err
}
