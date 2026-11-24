package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/notify/enum/v1"
)

type NotificationChannelStatus string

const (
	NotificationChannelStatusProcessing    NotificationChannelStatus = "processing"     // 正在投递
	NotificationChannelStatusSucceeded     NotificationChannelStatus = "succeeded"      // 投递成功
	NotificationChannelStatusSkipped       NotificationChannelStatus = "skipped"        // 无需投递
	NotificationChannelStatusFailed        NotificationChannelStatus = "failed"         // 投递失败
	NotificationChannelStatusUnknown       NotificationChannelStatus = "unknown"        // 状态未知
	NotificationChannelStatusInternalError NotificationChannelStatus = "internal_error" // 通道内部错误
	NotificationChannelStatusRateLimited   NotificationChannelStatus = "rate_limited"   // 触发通道限流
)

var NotificationChannelStatusMap = enum.NewMapping[NotificationChannelStatus, v1.NotificationChannelStatus](map[NotificationChannelStatus]enum.Entry[NotificationChannelStatus, v1.NotificationChannelStatus]{
	NotificationChannelStatusProcessing:    {Proto: v1.NotificationChannelStatus_NOTIFICATION_CHANNEL_STATUS_PROCESSING},
	NotificationChannelStatusSucceeded:     {Proto: v1.NotificationChannelStatus_NOTIFICATION_CHANNEL_STATUS_SUCCEEDED},
	NotificationChannelStatusSkipped:       {Proto: v1.NotificationChannelStatus_NOTIFICATION_CHANNEL_STATUS_SKIPPED},
	NotificationChannelStatusFailed:        {Proto: v1.NotificationChannelStatus_NOTIFICATION_CHANNEL_STATUS_FAILED},
	NotificationChannelStatusUnknown:       {Proto: v1.NotificationChannelStatus_NOTIFICATION_CHANNEL_STATUS_UNKNOWN},
	NotificationChannelStatusInternalError: {Proto: v1.NotificationChannelStatus_NOTIFICATION_CHANNEL_STATUS_INTERNAL_ERROR},
	NotificationChannelStatusRateLimited:   {Proto: v1.NotificationChannelStatus_NOTIFICATION_CHANNEL_STATUS_RATE_LIMITED},
})

// Blocking 表示该状态会阻断当前通道后续投递。
func (s NotificationChannelStatus) Blocking() bool {
	return s == NotificationChannelStatusProcessing ||
		s == NotificationChannelStatusFailed ||
		s == NotificationChannelStatusInternalError
}

// Merge 合并同一通知规则下多个投递目标的通道状态。
func (s NotificationChannelStatus) Merge(next NotificationChannelStatus) NotificationChannelStatus {
	if s == NotificationChannelStatusSkipped {
		return next
	}
	if s.Blocking() || next == NotificationChannelStatusSkipped {
		return s
	}
	if next.Blocking() {
		return next
	}
	if next == NotificationChannelStatusUnknown {
		return next
	}
	if next == NotificationChannelStatusRateLimited {
		return next
	}
	return s
}

func (s NotificationChannelStatus) String() string {
	return string(s)
}
