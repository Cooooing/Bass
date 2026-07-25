package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type LoginFailureReason string

const (
	LoginFailureReasonInvalidCredentials   LoginFailureReason = "invalid_credentials"
	LoginFailureReasonTotpInvalid          LoginFailureReason = "totp_invalid"
	LoginFailureReasonCodeInvalidOrExpired LoginFailureReason = "code_invalid_or_expired"
	LoginFailureReasonAccountNotNormal     LoginFailureReason = "account_not_normal"
	LoginFailureReasonNotImplemented       LoginFailureReason = "not_implemented"
	LoginFailureReasonInternal             LoginFailureReason = "internal"
)

var LoginFailureReasonMap = commonenum.NewMapping[LoginFailureReason, v1.LoginFailureReason](map[LoginFailureReason]commonenum.Entry[LoginFailureReason, v1.LoginFailureReason]{
	LoginFailureReasonInvalidCredentials:   {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_INVALID_CREDENTIALS},
	LoginFailureReasonTotpInvalid:          {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_TOTP_INVALID},
	LoginFailureReasonCodeInvalidOrExpired: {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_CODE_INVALID_OR_EXPIRED},
	LoginFailureReasonAccountNotNormal:     {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_ACCOUNT_NOT_NORMAL},
	LoginFailureReasonNotImplemented:       {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_NOT_IMPLEMENTED},
	LoginFailureReasonInternal:             {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_INTERNAL},
})
