package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

type AccountStatType string

const (
	AccountStatTypeFollow   AccountStatType = "follow"   // 用户关注数量
	AccountStatTypeFollower AccountStatType = "follower" // 用户粉丝数量
)

var AccountStatTypeMap = commonenum.NewMapping[AccountStatType, v1.AccountStatType](map[AccountStatType]commonenum.Entry[AccountStatType, v1.AccountStatType]{
	AccountStatTypeFollow:   {Proto: v1.AccountStatType_ACCOUNT_STAT_TYPE_FOLLOW},
	AccountStatTypeFollower: {Proto: v1.AccountStatType_ACCOUNT_STAT_TYPE_FOLLOWER},
})

func (e AccountStatType) String() string {
	return string(e)
}
