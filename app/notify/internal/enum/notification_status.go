package enum

import (
	v1 "common/api/gen/notify/v1"
	"common/pkg/enum"
)

type NotificationStatus string

const (
	NotificationStatusNormal NotificationStatus = "normal"
	NotificationStatusCancel NotificationStatus = "cancel"
)

var NotificationStatusMap = enum.NewMapping[NotificationStatus, v1.NotificationStatus](map[NotificationStatus]enum.Entry[NotificationStatus, v1.NotificationStatus]{
	NotificationStatusNormal: {Proto: v1.NotificationStatus_NOTIFICATION_STATUS_NORMAL},
	NotificationStatusCancel: {Proto: v1.NotificationStatus_NOTIFICATION_STATUS_CANCEL},
})
