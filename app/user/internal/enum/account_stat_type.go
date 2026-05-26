package enum

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/enum"
)

type AccountStatType string

const (
	AccountStatTypeFollow   AccountStatType = "follow"
	AccountStatTypeFollower AccountStatType = "follower"
)

var AccountStatTypeMap = enum.NewMapping[AccountStatType, v1.AccountStatType](map[AccountStatType]enum.Entry[AccountStatType, v1.AccountStatType]{
	AccountStatTypeFollow:   {Proto: v1.AccountStatType_ACCOUNT_STAT_TYPE_FOLLOW},
	AccountStatTypeFollower: {Proto: v1.AccountStatType_ACCOUNT_STAT_TYPE_FOLLOWER},
})
