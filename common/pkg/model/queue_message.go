package model

import notifyv1 "common/api/notify/v1"

type Notification struct {
	UUID     string                          `json:"uuid"`
	Type     *notifyv1.NotificationType      `json:"type"`
	SenderId int64                           `json:"sender_id"`
	Channels []*notifyv1.NotificationChannel `json:"channels"`
	Meta     map[string]any                  `json:"meta"`
	Status   notifyv1.NotificationStatus     `json:"status"`

	ReceiverIds   []int64 `json:"-"` // 接收者ID
	Content       string  `json:"-"` // 模板内容
	ContentRender string  `json:"-"` // 模板内容渲染结果(持久化)
}
