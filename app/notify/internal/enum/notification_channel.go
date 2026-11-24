package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/notify/enum/v1"
)

type NotificationChannel string

const (
	NotificationChannelStation     NotificationChannel = "station"      // 站内消息通道
	NotificationChannelEmail       NotificationChannel = "email"        // 电子邮件通道
	NotificationChannelTencentSMS  NotificationChannel = "tencent_sms"  // 腾讯云短信通道
	NotificationChannelLarkWebhook NotificationChannel = "lark_webhook" // 飞书机器人回调通道
)

var NotificationChannelMap = enum.NewMapping[NotificationChannel, v1.NotificationChannel](map[NotificationChannel]enum.Entry[NotificationChannel, v1.NotificationChannel]{
	NotificationChannelStation:     {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_STATION},
	NotificationChannelEmail:       {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL},
	NotificationChannelTencentSMS:  {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_TENCENT_SMS},
	NotificationChannelLarkWebhook: {Proto: v1.NotificationChannel_NOTIFICATION_CHANNEL_LARK_WEBHOOK},
})

func (e NotificationChannel) String() string {
	return string(e)
}
