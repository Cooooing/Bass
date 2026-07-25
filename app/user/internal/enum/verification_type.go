package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type VerificationType string

const (
	VerificationTypeEmail VerificationType = "email"
	VerificationTypePhone VerificationType = "phone"
)

var VerificationTypeMap = commonenum.NewMapping[VerificationType, v1.VerificationType](map[VerificationType]commonenum.Entry[VerificationType, v1.VerificationType]{
	VerificationTypeEmail: {Proto: v1.VerificationType_VERIFICATION_TYPE_EMAIL},
	VerificationTypePhone: {Proto: v1.VerificationType_VERIFICATION_TYPE_PHONE},
})
