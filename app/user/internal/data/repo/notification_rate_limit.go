package repo

import (
	"common/pkg/client/rpc"
	notifyv1 "common/proto/gen/notify/v1"
	notifyv1enum "common/proto/gen/notify/v1/enum"
	"context"
	bizrepo "user/internal/biz/repo"
)

var _ bizrepo.NotificationRateLimitClient = (*NotificationRateLimitClient)(nil)

type NotificationRateLimitClient struct {
	notifyClient *rpc.NotifyClient
}

func NewNotificationRateLimitClient(
	notifyClient *rpc.NotifyClient,
) bizrepo.NotificationRateLimitClient {
	return &NotificationRateLimitClient{
		notifyClient: notifyClient,
	}
}

func (c *NotificationRateLimitClient) CheckEmail(ctx context.Context, email string) (*bizrepo.NotificationRateLimitState, error) {
	reply, err := c.notifyClient.RateLimit.Check(ctx, &notifyv1.CheckNotificationRateLimit_Req{
		Channel:   notifyv1enum.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL,
		Recipient: email,
	})
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationRateLimitState{
		Limited:           reply.GetLimited(),
		RetryAfterSeconds: reply.GetRetryAfterSeconds(),
	}, nil
}

func (c *NotificationRateLimitClient) CheckPhone(ctx context.Context, phone string) (*bizrepo.NotificationRateLimitState, error) {
	reply, err := c.notifyClient.RateLimit.Check(ctx, &notifyv1.CheckNotificationRateLimit_Req{
		Channel:   notifyv1enum.NotificationChannel_NOTIFICATION_CHANNEL_TENCENT_SMS,
		Recipient: phone,
	})
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationRateLimitState{
		Limited:           reply.GetLimited(),
		RetryAfterSeconds: reply.GetRetryAfterSeconds(),
	}, nil
}
