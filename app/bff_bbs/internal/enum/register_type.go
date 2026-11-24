package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/bff_bbs/user/enum/v1"
	userv1 "common/proto/gen/user/enum/v1"
)

type RegisterType string

const (
	RegisterTypeEmail RegisterType = "email" // 邮箱注册
	RegisterTypePhone RegisterType = "phone" // 手机注册
)

var RegisterTypeMap = commonenum.NewMapping[RegisterType, v1.RegisterType](map[RegisterType]commonenum.Entry[RegisterType, v1.RegisterType]{
	RegisterTypeEmail: {Proto: v1.RegisterType_REGISTER_TYPE_EMAIL},
	RegisterTypePhone: {Proto: v1.RegisterType_REGISTER_TYPE_PHONE},
})

func (e RegisterType) String() string {
	return string(e)
}

func (e RegisterType) ToUserProto() userv1.RegisterType {
	switch e {
	case RegisterTypeEmail:
		return userv1.RegisterType_REGISTER_TYPE_EMAIL
	case RegisterTypePhone:
		return userv1.RegisterType_REGISTER_TYPE_PHONE
	default:
		return userv1.RegisterType_REGISTER_TYPE_UNSPECIFIED
	}
}
