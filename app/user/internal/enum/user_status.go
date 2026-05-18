package enum

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/enum"
)

type UserStatus string

const (
	UserStatusNormal  UserStatus = "normal"
	UserStatusBanned  UserStatus = "banned"
	UserStatusDeleted UserStatus = "deleted"
)

var UserStatusMap = enum.NewMapping[UserStatus, v1.UserStatus](map[UserStatus]enum.Entry[UserStatus, v1.UserStatus]{
	UserStatusNormal:  {Proto: v1.UserStatus_USER_STATUS_NORMAL},
	UserStatusBanned:  {Proto: v1.UserStatus_USER_STATUS_BANNED},
	UserStatusDeleted: {Proto: v1.UserStatus_USER_STATUS_DELETED},
})
