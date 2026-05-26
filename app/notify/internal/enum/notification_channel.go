package enum

import (
	v1 "common/api/gen/notify/v1"
	"common/pkg/enum"
)

type NotificationChannel string

const (
	NotificationChannelEmail   NotificationChannel = "email"
	NotificationChannelSMS     NotificationChannel = "sms"
	NotificationChannelStation NotificationChannel = "station"
	NotificationChannelWebhook NotificationChannel = "webhook"
)

var NotificationChannelMap = enum.NewMapping[NotificationChannel, v1.NotificationChannel](map[NotificationChannel]enum.Entry[NotificationChannel, v1.NotificationChannel]{
	NotificationChannelEmail:   {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL},
	NotificationChannelSMS:     {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_SMS},
	NotificationChannelStation: {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_STATION},
	NotificationChannelWebhook: {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_WEBHOOK},
})
