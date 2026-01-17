package constant

import (
	v1 "common/api/notify/v1"
	"fmt"
)

// Redis key
var (
	RequestNonce            = "RequestNonce::{%s}"            // 请求防重放
	TokenVerifyCode         = "TokenVerifyCode::{%s}::{%s}"   // 验证码 Token
	Token                   = "Token::{%s}"                   // Token
	NotificationTemplateMap = "NotificationTemplateMap"       // 通知模板
	TwoFactorAuthentication = "TwoFactorAuthentication::{%s}" // 2FA 验证码，首次启用 2FA 时使用
)

func GetKeyRequestNonce(nonce string) string {
	return fmt.Sprintf(RequestNonce, nonce)
}

func GetKeyTokenVerityCode(verifyCodeType VerifyCodeType, account string) string {
	return fmt.Sprintf(TokenVerifyCode, verifyCodeType, account)
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
	VerifyCodeTypeRegisterEmail VerifyCodeType = "RegisterEmail"
	VerifyCodeTypeRegisterPhone VerifyCodeType = "RegisterPhone"
)
