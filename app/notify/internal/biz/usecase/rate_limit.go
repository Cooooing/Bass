package usecase

import (
	"context"
	"notify/internal/biz/repo"
	"notify/internal/config"
	notifyenum "notify/internal/enum"
	"time"
)

type RateLimitUsecase struct {
	notificationRateLimitCache repo.NotificationRateLimitCache
	enabled                    bool
	window                     time.Duration
	maxCount                   int64
}

func NewRateLimitUsecase(conf *config.Bootstrap, notificationRateLimitCache repo.NotificationRateLimitCache) *RateLimitUsecase {
	enabled := true
	window := 5 * time.Minute
	maxCount := int64(5)
	if conf != nil && conf.Notify != nil && conf.Notify.NotificationRateLimit != nil {
		enabled = conf.Notify.NotificationRateLimit.Enable
		if conf.Notify.NotificationRateLimit.Window != nil && conf.Notify.NotificationRateLimit.Window.AsDuration() > 0 {
			window = conf.Notify.NotificationRateLimit.Window.AsDuration()
		}
		if conf.Notify.NotificationRateLimit.MaxCount > 0 {
			maxCount = conf.Notify.NotificationRateLimit.MaxCount
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
