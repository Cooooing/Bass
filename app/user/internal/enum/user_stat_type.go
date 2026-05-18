package enum

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/enum"
)

type UserStatType string

const (
	UserStatTypeFollow   UserStatType = "follow"
	UserStatTypeFollower UserStatType = "follower"
	UserStatTypeBlock    UserStatType = "block"
	UserStatTypeBlocked  UserStatType = "blocked"
)

var UserStatTypeMap = enum.NewMapping[UserStatType, v1.UserStatType](map[UserStatType]enum.Entry[UserStatType, v1.UserStatType]{
	UserStatTypeFollow:   {Proto: v1.UserStatType_USER_STAT_TYPE_FOLLOW},
	UserStatTypeFollower: {Proto: v1.UserStatType_USER_STAT_TYPE_FOLLOWER},
	UserStatTypeBlock:    {Proto: v1.UserStatType_USER_STAT_TYPE_BLOCK},
	UserStatTypeBlocked:  {Proto: v1.UserStatType_USER_STAT_TYPE_BLOCKED},
})
