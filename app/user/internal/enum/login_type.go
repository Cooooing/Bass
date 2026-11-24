package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

type LoginType string

const (
	LoginTypePassword LoginType = "password" // 账号密码登录
	LoginTypeEmail    LoginType = "email"    // 邮箱验证码登录
	LoginTypePhone    LoginType = "phone"    // 手机验证码登录
)

var LoginTypeMap = commonenum.NewMapping[LoginType, v1.LoginType](map[LoginType]commonenum.Entry[LoginType, v1.LoginType]{
	LoginTypePassword: {Proto: v1.LoginType_LOGIN_TYPE_PASSWORD},
	LoginTypeEmail:    {Proto: v1.LoginType_LOGIN_TYPE_EMAIL},
	LoginTypePhone:    {Proto: v1.LoginType_LOGIN_TYPE_PHONE},
})

func (e LoginType) String() string {
	return string(e)
}
