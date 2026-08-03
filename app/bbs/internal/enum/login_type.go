package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/bbs/v1/user/enum"
	userv1 "common/proto/gen/user/v1/enum"
)

type LoginType string

const (
	LoginTypePassword LoginType = "password"
	LoginTypeEmail    LoginType = "email"
	LoginTypePhone    LoginType = "phone"
)

var LoginTypeMap = commonenum.NewMapping[LoginType, v1.LoginType](map[LoginType]commonenum.Entry[LoginType, v1.LoginType]{
	LoginTypePassword: {Proto: v1.LoginType_LOGIN_TYPE_PASSWORD},
	LoginTypeEmail:    {Proto: v1.LoginType_LOGIN_TYPE_EMAIL},
	LoginTypePhone:    {Proto: v1.LoginType_LOGIN_TYPE_PHONE},
})

func (e LoginType) String() string {
	return string(e)
}

func (e LoginType) ToUserProto() userv1.LoginType {
	switch e {
	case LoginTypePassword:
		return userv1.LoginType_LOGIN_TYPE_PASSWORD
	case LoginTypeEmail:
		return userv1.LoginType_LOGIN_TYPE_EMAIL
	case LoginTypePhone:
		return userv1.LoginType_LOGIN_TYPE_PHONE
	default:
		return userv1.LoginType_LOGIN_TYPE_UNSPECIFIED
	}
}
