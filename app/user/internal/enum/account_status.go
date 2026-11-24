package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

type AccountStatus string

const (
	AccountStatusNormal    AccountStatus = "normal"    // 账号状态正常
	AccountStatusBanned    AccountStatus = "banned"    // 账号已被封禁
	AccountStatusCancelled AccountStatus = "cancelled" // 账号已注销
)

var AccountStatusMap = commonenum.NewMapping[AccountStatus, v1.AccountStatus](map[AccountStatus]commonenum.Entry[AccountStatus, v1.AccountStatus]{
	AccountStatusNormal:    {Proto: v1.AccountStatus_ACCOUNT_STATUS_NORMAL},
	AccountStatusBanned:    {Proto: v1.AccountStatus_ACCOUNT_STATUS_BANNED},
	AccountStatusCancelled: {Proto: v1.AccountStatus_ACCOUNT_STATUS_CANCELLED},
})

func (e AccountStatus) String() string {
	return string(e)
}
