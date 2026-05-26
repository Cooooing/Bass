package channel

import (
	"fmt"
	bizchannel "notify/internal/biz/channel"
	notifyenum "notify/internal/enum"
)

// Client 根据通知渠道返回对应发送实现。
type Client struct {
	emailChannel   *EmailChannel
	smsChannel     *SMSChannel
	webhookChannel *WebhookChannel
}

func NewClient(
	emailChannel *EmailChannel,
	smsChannel *SMSChannel,
	webhookChannel *WebhookChannel,
) bizchannel.Client {
	return &Client{
		emailChannel:   emailChannel,
		smsChannel:     smsChannel,
		webhookChannel: webhookChannel,
	}
}

func (c *Client) GetChannel(channel notifyenum.NotificationChannel) (bizchannel.Channel, error) {
	switch channel {
	case notifyenum.NotificationChannelEmail:
		if c.emailChannel != nil {
			return c.emailChannel, nil
		}
	case notifyenum.NotificationChannelSMS:
		if c.smsChannel != nil {
			return c.smsChannel, nil
		}
	case notifyenum.NotificationChannelWebhook:
		if c.webhookChannel != nil {
			return c.webhookChannel, nil
		}
	}
	return nil, fmt.Errorf("notification channel is not configured: %s", channel)
}
