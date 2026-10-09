package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/notify/v1/enum"
)

type SMSType string

const (
	SMSTypeTencent SMSType = "tencent" // tencent
)

var SMSTypeMap = commonenum.NewMapping[SMSType, v1.SMSType](map[SMSType]commonenum.Entry[SMSType, v1.SMSType]{
	SMSTypeTencent: {Proto: v1.SMSType_SMS_TYPE_TENCENT},
})

func (e SMSType) String() string {
	return string(e)
}
