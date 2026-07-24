package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type LoginType string

const (
	LoginTypePassword LoginType = "password"
	LoginTypeEmail    LoginType = "email"
	LoginTypePhone    LoginType = "phone"
)

var LoginTypeMap = enum.NewMapping[LoginType, v1.LoginType](map[LoginType]enum.Entry[LoginType, v1.LoginType]{
	LoginTypePassword: {Proto: v1.LoginType_LOGIN_TYPE_PASSWORD},
	LoginTypeEmail:    {Proto: v1.LoginType_LOGIN_TYPE_EMAIL},
	LoginTypePhone:    {Proto: v1.LoginType_LOGIN_TYPE_PHONE},
})

type LoginRealm string

const (
	LoginRealmBBS      LoginRealm = "bbs"
	LoginRealmBBSAdmin LoginRealm = "bbs_admin"
)

var LoginRealmMap = enum.NewMapping[LoginRealm, v1.LoginRealm](map[LoginRealm]enum.Entry[LoginRealm, v1.LoginRealm]{
	LoginRealmBBS:      {Proto: v1.LoginRealm_LOGIN_REALM_BBS},
	LoginRealmBBSAdmin: {Proto: v1.LoginRealm_LOGIN_REALM_BBS_ADMIN},
})

type ClientType string

const (
	ClientTypeUnknown  ClientType = "unknown"
	ClientTypeWeb      ClientType = "web"
	ClientTypeAdminWeb ClientType = "admin_web"
	ClientTypeMobile   ClientType = "mobile"
)

var ClientTypeMap = enum.NewMapping[ClientType, v1.ClientType](map[ClientType]enum.Entry[ClientType, v1.ClientType]{
	ClientTypeUnknown:  {Proto: v1.ClientType_CLIENT_TYPE_UNKNOWN},
	ClientTypeWeb:      {Proto: v1.ClientType_CLIENT_TYPE_WEB},
	ClientTypeAdminWeb: {Proto: v1.ClientType_CLIENT_TYPE_ADMIN_WEB},
	ClientTypeMobile:   {Proto: v1.ClientType_CLIENT_TYPE_MOBILE},
})

type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "unknown"
	DeviceTypeDesktop DeviceType = "desktop"
	DeviceTypeMobile  DeviceType = "mobile"
	DeviceTypeTablet  DeviceType = "tablet"
	DeviceTypeBot     DeviceType = "bot"
)

var DeviceTypeMap = enum.NewMapping[DeviceType, v1.DeviceType](map[DeviceType]enum.Entry[DeviceType, v1.DeviceType]{
	DeviceTypeUnknown: {Proto: v1.DeviceType_DEVICE_TYPE_UNKNOWN},
	DeviceTypeDesktop: {Proto: v1.DeviceType_DEVICE_TYPE_DESKTOP},
	DeviceTypeMobile:  {Proto: v1.DeviceType_DEVICE_TYPE_MOBILE},
	DeviceTypeTablet:  {Proto: v1.DeviceType_DEVICE_TYPE_TABLET},
	DeviceTypeBot:     {Proto: v1.DeviceType_DEVICE_TYPE_BOT},
})

type LoginStatus string

const (
	LoginStatusSuccess LoginStatus = "success"
	LoginStatusFailed  LoginStatus = "failed"
)

var LoginStatusMap = enum.NewMapping[LoginStatus, v1.LoginStatus](map[LoginStatus]enum.Entry[LoginStatus, v1.LoginStatus]{
	LoginStatusSuccess: {Proto: v1.LoginStatus_LOGIN_STATUS_SUCCESS},
	LoginStatusFailed:  {Proto: v1.LoginStatus_LOGIN_STATUS_FAILED},
})

type LoginFailureReason string

const (
	LoginFailureReasonInvalidCredentials    LoginFailureReason = "invalid_credentials"
	LoginFailureReasonTotpInvalid           LoginFailureReason = "totp_invalid"
	LoginFailureReasonCodeInvalidOrExpired  LoginFailureReason = "code_invalid_or_expired"
	LoginFailureReasonAccountNotNormal      LoginFailureReason = "account_not_normal"
	LoginFailureReasonNotImplemented        LoginFailureReason = "not_implemented"
	LoginFailureReasonInternal              LoginFailureReason = "internal"
)

var LoginFailureReasonMap = enum.NewMapping[LoginFailureReason, v1.LoginFailureReason](map[LoginFailureReason]enum.Entry[LoginFailureReason, v1.LoginFailureReason]{
	LoginFailureReasonInvalidCredentials:   {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_INVALID_CREDENTIALS},
	LoginFailureReasonTotpInvalid:          {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_TOTP_INVALID},
	LoginFailureReasonCodeInvalidOrExpired: {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_CODE_INVALID_OR_EXPIRED},
	LoginFailureReasonAccountNotNormal:     {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_ACCOUNT_NOT_NORMAL},
	LoginFailureReasonNotImplemented:       {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_NOT_IMPLEMENTED},
	LoginFailureReasonInternal:             {Proto: v1.LoginFailureReason_LOGIN_FAILURE_REASON_INTERNAL},
})