package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/notify/enum/v1"
)

type SMSType string

const (
	SMSTypeTencent SMSType = "tencent" // 腾讯云短信服务
)

var SMSTypeMap = commonenum.NewMapping[SMSType, v1.SMSType](map[SMSType]commonenum.Entry[SMSType, v1.SMSType]{
	SMSTypeTencent: {Proto: v1.SMSType_SMS_TYPE_TENCENT},
})

func (e SMSType) String() string {
	return string(e)
}
