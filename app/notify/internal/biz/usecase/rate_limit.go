package usecase

import (
	"context"
	"time"

	"notify/internal/biz/repo"
	"notify/internal/config"
	notifyenum "notify/internal/enum"
)

type RateLimitUsecase struct {
	notificationRateLimitCache repo.NotificationRateLimitCache
	enabled                    bool
	window                     time.Duration
	maxCount                   int64
}

func NewRateLimitUsecase(
	conf *config.Bootstrap,
	notificationRateLimitCache repo.NotificationRateLimitCache,
) *RateLimitUsecase {
	enabled := true
	window := 5 * time.Minute
	maxCount := int64(5)
	if conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetNotificationRateLimit() != nil {
		enabled = conf.GetBusiness().GetNotificationRateLimit().GetEnable()
		if conf.GetBusiness().GetNotificationRateLimit().GetWindow() != nil && conf.GetBusiness().GetNotificationRateLimit().GetWindow().AsDuration() > 0 {
			window = conf.GetBusiness().GetNotificationRateLimit().GetWindow().AsDuration()
		}
		if conf.GetBusiness().GetNotificationRateLimit().GetMaxCount() > 0 {
			maxCount = conf.GetBusiness().GetNotificationRateLimit().GetMaxCount()
		}
	}
	return &RateLimitUsecase{
		notificationRateLimitCache: notificationRateLimitCache,
		enabled:                    enabled,
		window:                     window,
		maxCount:                   maxCount,
	}
}

type RateLimitCheckReq struct {
	Channel   notifyenum.NotificationChannel
	Recipient string
}

type RateLimitCheckResp struct {
	Limited        bool
	RetryAfter     time.Duration
	RemainingCount int64
}

func (u *RateLimitUsecase) Check(ctx context.Context, req *RateLimitCheckReq) (*RateLimitCheckResp, error) {
	if req == nil {
		req = &RateLimitCheckReq{}
	}
	if !u.enabled {
		return &RateLimitCheckResp{
			RemainingCount: u.maxCount,
		}, nil
	}
	checkResp, err := u.notificationRateLimitCache.Check(ctx, &repo.NotificationRateLimitSpec{
		Channel:   req.Channel,
		Recipient: req.Recipient,
		Window:    u.window,
		MaxCount:  u.maxCount,
	})
	if err != nil {
		return nil, err
	}
	if checkResp == nil {
		return &RateLimitCheckResp{}, nil
	}
	return &RateLimitCheckResp{
		Limited:        checkResp.Limited,
		RetryAfter:     checkResp.RetryAfter,
		RemainingCount: checkResp.RemainingCount,
	}, nil
}
