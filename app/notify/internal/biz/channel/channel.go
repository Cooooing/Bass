package channel

import (
	commonenum "common/pkg/enum"
	"context"

	notifyenum "notify/internal/enum"
)

type Client interface {
	GetChannel(channel notifyenum.NotificationChannel) (Channel, error)
}

// Channel 是 biz 层面对具体投递渠道的统一发送能力。
type Channel interface {
	Send(ctx context.Context, req *SendReq) error
}

// SendReq 是外部投递任务在发送阶段使用的扁平请求。
type SendReq struct {
	EventID        string
	EventType      commonenum.EventType
	ReceiverID     *int64
	Channel        notifyenum.NotificationChannel
	Target         string
	Title          string
	Content        string
	TemplateID     string
	TemplateParams []string
}
