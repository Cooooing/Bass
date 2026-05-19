package sender

import (
	v1 "common/api/gen/notify/v1"
	"context"
)

// StationSender 站内信发送器（记录由 NotifyUsecase 创建，此处为占位）
type StationSender struct{}

func NewStationSender() *StationSender {
	return &StationSender{}
}

func (s *StationSender) Channel() v1.NotificationChannel {
	return v1.NotificationChannel_NOTIFICATION_CHANNEL_STATION
}

func (s *StationSender) Send(_ context.Context, _ *SendRequest) error {
	return nil
}
