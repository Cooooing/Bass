package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type AccountStatus string

const (
	AccountStatusNormal    AccountStatus = "normal"
	AccountStatusBanned    AccountStatus = "banned"
	AccountStatusCancelled AccountStatus = "cancelled"
)

var AccountStatusMap = enum.NewMapping[AccountStatus, v1.AccountStatus](map[AccountStatus]enum.Entry[AccountStatus, v1.AccountStatus]{
	AccountStatusNormal:    {Proto: v1.AccountStatus_ACCOUNT_STATUS_NORMAL},
	AccountStatusBanned:    {Proto: v1.AccountStatus_ACCOUNT_STATUS_BANNED},
	AccountStatusCancelled: {Proto: v1.AccountStatus_ACCOUNT_STATUS_CANCELLED},
})