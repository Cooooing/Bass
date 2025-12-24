package constant

import (
	v1 "common/api/notify/v1"
	"fmt"
)

var Authentication = "Authorization" // token 请求头名称

// Redis key
var (
	TokenEmailCode          = "TokenVerifyCode::{%s}::{%s}"
	Token                   = "Token::{%s}"
	NotificationTemplateMap = "NotificationTemplateMap"
	TwoFactorAuthentication = "TwoFactorAuthentication::{%s}"
)

func GetKeyTokenVerityCode(verifyCodeType VerifyCodeType, account string) string {
	return fmt.Sprintf(TokenEmailCode, verifyCodeType, account)
}

func GetKeyToken(token string) string {
	return fmt.Sprintf(Token, token)
}

func GetKeyNotificationTemplateMap() string {
	return NotificationTemplateMap
}

func GetKeyNotificationTemplate(notificationType *v1.NotificationType, channel *v1.NotificationChannel) string {
	if notificationType == nil || channel == nil {
		return ""
	}
	return fmt.Sprintf("%s_%s", notificationType.String(), channel.String())
}

func GetKeyTwoFactorAuthentication(name string) string {
	return fmt.Sprintf(TwoFactorAuthentication, name)
}

type VerifyCodeType string

func (v VerifyCodeType) String() string {
	return string(v)
}

const (
	VerifyCodeTypeRegisterEmail VerifyCodeType = "register_email"
	VerifyCodeTypeRegisterPhone VerifyCodeType = "register_phone"
)
