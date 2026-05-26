package enum

import (
	v1 "common/api/gen/notify/v1"
	"common/pkg/enum"
)

type NotificationDeliveryStatus string

const (
	NotificationDeliveryStatusPending NotificationDeliveryStatus = "pending"
	NotificationDeliveryStatusSending NotificationDeliveryStatus = "sending"
	NotificationDeliveryStatusSent    NotificationDeliveryStatus = "sent"
	NotificationDeliveryStatusFailed  NotificationDeliveryStatus = "failed"
)

var NotificationDeliveryStatusMap = enum.NewMapping[NotificationDeliveryStatus, v1.NotificationDeliveryStatus](map[NotificationDeliveryStatus]enum.Entry[NotificationDeliveryStatus, v1.NotificationDeliveryStatus]{
	NotificationDeliveryStatusPending: {Proto: v1.NotificationDeliveryStatus_NOTIFICATION_DELIVERY_STATUS_PENDING},
	NotificationDeliveryStatusSending: {Proto: v1.NotificationDeliveryStatus_NOTIFICATION_DELIVERY_STATUS_SENDING},
	NotificationDeliveryStatusSent:    {Proto: v1.NotificationDeliveryStatus_NOTIFICATION_DELIVERY_STATUS_SENT},
	NotificationDeliveryStatusFailed:  {Proto: v1.NotificationDeliveryStatus_NOTIFICATION_DELIVERY_STATUS_FAILED},
})
