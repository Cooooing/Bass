package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

type VerificationType string

const (
	VerificationTypeEmail VerificationType = "email" // 邮箱验证
	VerificationTypePhone VerificationType = "phone" // 手机验证
)

var VerificationTypeMap = commonenum.NewMapping[VerificationType, v1.VerificationType](map[VerificationType]commonenum.Entry[VerificationType, v1.VerificationType]{
	VerificationTypeEmail: {Proto: v1.VerificationType_VERIFICATION_TYPE_EMAIL},
	VerificationTypePhone: {Proto: v1.VerificationType_VERIFICATION_TYPE_PHONE},
})

func (e VerificationType) String() string {
	return string(e)
}
