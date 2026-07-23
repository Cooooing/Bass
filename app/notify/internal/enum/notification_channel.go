package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/notify/v1/enum"
)

type NotificationChannel string

const (
	NotificationChannelStation     NotificationChannel = "station"
	NotificationChannelEmail       NotificationChannel = "email"
	NotificationChannelTencentSMS  NotificationChannel = "tencent_sms"
	NotificationChannelLarkWebhook NotificationChannel = "lark_webhook"
)

var NotificationChannelMap = enum.NewMapping[NotificationChannel, v1.NotificationChannel](map[NotificationChannel]enum.Entry[NotificationChannel, v1.NotificationChannel]{
	NotificationChannelStation:     {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_STATION},
	NotificationChannelEmail:       {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL},
	NotificationChannelTencentSMS:  {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_TENCENT_SMS},
	NotificationChannelLarkWebhook: {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_LARK_WEBHOOK},
})
