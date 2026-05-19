package sender

import (
	v1 "common/api/gen/notify/v1"
	"context"
)

// UserInfo 用户联系信息（由 NotifyUsecase 通过 UserResolver 获取）
type UserInfo struct {
	Phone string
	Email string
}

// ChannelSender 渠道发送器接口
type ChannelSender interface {
	Channel() v1.NotificationChannel
	Send(ctx context.Context, req *SendRequest) error
}

// SendRequest 发送请求
type SendRequest struct {
	ReceiverID int64
	Title      string
	Content    string
	UserInfo   UserInfo
}

// Registry 渠道发送器注册表
type Registry struct {
	senders map[v1.NotificationChannel]ChannelSender
}

// NewRegistry 创建注册表
func NewRegistry(senders ...ChannelSender) *Registry {
	m := make(map[v1.NotificationChannel]ChannelSender, len(senders))
	for _, s := range senders {
		if s != nil {
			m[s.Channel()] = s
		}
	}
	return &Registry{senders: m}
}

// ProvideRegistry wire provider
func ProvideRegistry(smtp *SmtpSender, sms *TencentSmsSender) *Registry {
	var senders []ChannelSender
	senders = append(senders, NewStationSender())
	if smtp != nil {
		senders = append(senders, smtp)
	}
	if sms != nil {
		senders = append(senders, sms)
	}
	return NewRegistry(senders...)
}

// Get 根据渠道获取发送器
func (r *Registry) Get(channel v1.NotificationChannel) (ChannelSender, bool) {
	s, ok := r.senders[channel]
	return s, ok
}

// Send 通过渠道直接发送
func (r *Registry) Send(ctx context.Context, channel v1.NotificationChannel, req *SendRequest) error {
	s, ok := r.Get(channel)
	if !ok {
		return nil
	}
	return s.Send(ctx, req)
}

// ChannelToProto 将 ent enum (string) 转为 proto enum (int32)
func ChannelToProto(ch string) v1.NotificationChannel {
	return v1.NotificationChannel(v1.NotificationChannel_value[ch])
}

// StatusToProto 将 ent enum (string) 转为 proto enum (int32)
func StatusToProto(s string) v1.NotificationStatus {
	return v1.NotificationStatus(v1.NotificationStatus_value[s])
}
