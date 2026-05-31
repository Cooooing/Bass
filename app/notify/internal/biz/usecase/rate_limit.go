package usecase

import (
	"context"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	notifyenum "notify/internal/enum"
	"time"
)

type RateLimitUsecase struct {
	notificationRateLimitCache repo.NotificationRateLimitCache
	enabled                    bool
	window                     time.Duration
	maxCount                   int64
}

func NewRateLimitUsecase(conf *conf.Bootstrap, notificationRateLimitCache repo.NotificationRateLimitCache) *RateLimitUsecase {
	enabled := true
	window := 5 * time.Minute
	maxCount := int64(5)
	if conf != nil && conf.Server != nil && conf.Server.NotificationRateLimit != nil {
		enabled = conf.Server.NotificationRateLimit.Enable
		if conf.Server.NotificationRateLimit.Window != nil && conf.Server.NotificationRateLimit.Window.AsDuration() > 0 {
			window = conf.Server.NotificationRateLimit.Window.AsDuration()
		}
		if conf.Server.NotificationRateLimit.MaxCount > 0 {
			maxCount = conf.Server.NotificationRateLimit.MaxCount
		}
	}
	return &RateLimitUsecase{
		notificationRateLimitCache: notificationRateLimitCache,
		enabled:                    enabled,
		window:                     window,
		maxCount:                   maxCount,
	}
}

func (u *RateLimitUsecase) Check(ctx context.Context, channel notifyenum.NotificationChannel, recipient string) (*repo.NotificationRateLimitState, error) {
	if !u.enabled {
		return &repo.NotificationRateLimitState{RemainingCount: u.maxCount}, nil
	}
	return u.notificationRateLimitCache.Check(ctx, &repo.NotificationRateLimitSpec{
		Channel:   channel,
		Recipient: recipient,
		Window:    u.window,
		MaxCount:  u.maxCount,
	})
}
