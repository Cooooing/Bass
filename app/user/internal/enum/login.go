package enum

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/enum"
)

type LoginMethod string

const (
	LoginMethodPassword LoginMethod = "password"
)

// LoginMethodMap 绑定数据库字符串枚举和对外 proto 枚举。
var LoginMethodMap = enum.NewMapping[LoginMethod, v1.LoginMethod](map[LoginMethod]enum.Entry[LoginMethod, v1.LoginMethod]{
	LoginMethodPassword: {Proto: v1.LoginMethod_LOGIN_METHOD_PASSWORD},
})

type LoginStatus string

const (
	LoginStatusSuccess LoginStatus = "success"
	LoginStatusFailed  LoginStatus = "failed"
)

// LoginStatusMap 保持登录审计落库值和 proto 枚举一致。
var LoginStatusMap = enum.NewMapping[LoginStatus, v1.LoginStatus](map[LoginStatus]enum.Entry[LoginStatus, v1.LoginStatus]{
	LoginStatusSuccess: {Proto: v1.LoginStatus_LOGIN_STATUS_SUCCESS},
	LoginStatusFailed:  {Proto: v1.LoginStatus_LOGIN_STATUS_FAILED},
})
