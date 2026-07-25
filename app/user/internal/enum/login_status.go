package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type LoginStatus string

const (
	LoginStatusSuccess LoginStatus = "success"
	LoginStatusFailed  LoginStatus = "failed"
)

var LoginStatusMap = commonenum.NewMapping[LoginStatus, v1.LoginStatus](map[LoginStatus]commonenum.Entry[LoginStatus, v1.LoginStatus]{
	LoginStatusSuccess: {Proto: v1.LoginStatus_LOGIN_STATUS_SUCCESS},
	LoginStatusFailed:  {Proto: v1.LoginStatus_LOGIN_STATUS_FAILED},
})
