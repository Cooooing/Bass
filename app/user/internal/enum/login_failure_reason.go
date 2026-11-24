package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

type LoginFailureReason string

const (
	LoginFailureReasonInvalidCredentials   LoginFailureReason = "invalid_credentials"     // #nosec G101 -- 登录失败枚举值，不包含敏感凭据。
	LoginFailureReasonTotpInvalid          LoginFailureReason = "totp_invalid"            // 双因素验证码无效
	LoginFailureReasonCodeInvalidOrExpired LoginFailureReason = "code_invalid_or_expired" // 验证码无效或已过期
	LoginFailureReasonAccountNotNormal     LoginFailureReason = "account_not_normal"      // 账号状态不允许登录
	LoginFailureReasonNotImplemented       LoginFailureReason = "not_implemented"         // 登录方式尚未实现
	LoginFailureReasonInternal             LoginFailureReason = "internal"                // 系统内部登录错误
)

var LoginFailureReasonMap = commonenum.NewMapping[LoginFailureReason, v1.LoginFailureReason](map[LoginFailureReason]commonenum.Entry[LoginFailureReason, v1.LoginFailureReason]{
	LoginFailureReasonInvalidCredentials:   {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_INVALID_CREDENTIALS},
	LoginFailureReasonTotpInvalid:          {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_TOTP_INVALID},
	LoginFailureReasonCodeInvalidOrExpired: {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_CODE_INVALID_OR_EXPIRED},
	LoginFailureReasonAccountNotNormal:     {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_ACCOUNT_NOT_NORMAL},
	LoginFailureReasonNotImplemented:       {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_NOT_IMPLEMENTED},
	LoginFailureReasonInternal:             {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_INTERNAL},
})

func (e LoginFailureReason) String() string {
	return string(e)
}
