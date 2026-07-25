package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type AccountStatType string

const (
	AccountStatTypeFollow   AccountStatType = "follow"
	AccountStatTypeFollower AccountStatType = "follower"
)

var AccountStatTypeMap = commonenum.NewMapping[AccountStatType, v1.AccountStatType](map[AccountStatType]commonenum.Entry[AccountStatType, v1.AccountStatType]{
	AccountStatTypeFollow:   {Proto: v1.AccountStatType_ACCOUNT_STAT_TYPE_FOLLOW},
	AccountStatTypeFollower: {Proto: v1.AccountStatType_ACCOUNT_STAT_TYPE_FOLLOWER},
})
