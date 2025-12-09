package constant

import (
	v1 "common/api/notify/v1"
	"common/pkg/cutil/base"
	"fmt"
	"strings"
)

var Authentication = "Authorization" // token 请求头名称

// Redis key
var (
	TokenEmailCode          = "TokenEmailCode::{%s}"
	Token                   = "Token::{%s}"
	NotificationTemplateMap = "NotificationTemplateMap"
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

func GetNotificationTemplateTypeFromKey(key string) *v1.NotificationType {
	s := strings.Split(key, "_")[0]
	if v, ok := v1.NotificationType_value[s]; ok {
		return base.Ptr(v1.NotificationType(v))
	}
	return nil
}

func GetNotificationTemplateChannelFromKey(key string) *v1.NotificationChannel {
	s := strings.Split(key, "_")[1]
	if v, ok := v1.NotificationChannel_value[s]; ok {
		return base.Ptr(v1.NotificationChannel(v))
	}
	return nil
}
