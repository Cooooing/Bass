package constant

import (
	v1 "common/api/notify/v1"
	"fmt"
)

var Authentication = "Authorization" // token 请求头名称

// Redis key
var (
	TokenEmailCode          = "TokenEmailCode::{%s}"
	Token                   = "Token::{%s}"
	NotificationTemplateMap = "NotificationTemplateMap"
	TwoFactorAuthentication = "TwoFactorAuthentication::{%s}"
)

func GetKeyTokenEmailCode(email string) string {
	return fmt.Sprintf(TokenEmailCode, email)
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
