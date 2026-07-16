package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/notify/v1"
	"context"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type RateLimitService struct {
	v1.UnimplementedNotifyRateLimitServiceServer
	rateLimitUsecase *usecase.RateLimitUsecase
}

func NewRateLimitService(rateLimitUsecase *usecase.RateLimitUsecase) *RateLimitService {
	return &RateLimitService{rateLimitUsecase: rateLimitUsecase}
}

func (s *RateLimitService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyRateLimitServiceServer(gs, s)
}

func (s *RateLimitService) RegisterHttp(hs *http.Server) {}

func (s *RateLimitService) Check(ctx context.Context, req *v1.CheckNotificationRateLimit_Request) (*v1.CheckNotificationRateLimit_Response, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_NOTIFY_RATE_LIMIT_REQUEST_INVALID)
	}
	channel, ok := notifyenum.NotificationChannelMap.ToEnum(req.GetChannel())
	if !ok || (channel != notifyenum.NotificationChannelEmail && channel != notifyenum.NotificationChannelTencentSMS) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_NOTIFY_CHANNEL_INVALID)
	}
	recipient := strings.TrimSpace(req.GetRecipient())
	if recipient == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_NOTIFY_RECIPIENT_INVALID)
	}
	state, err := s.rateLimitUsecase.Check(ctx, &usecase.RateLimitCheckReq{Channel: channel, Recipient: recipient})
	if err != nil {
		return nil, err
	}
	retryAfterSeconds := int64(0)
	if state.RetryAfter > 0 {
		retryAfterSeconds = int64((state.RetryAfter + time.Second - 1) / time.Second)
	}
	return &v1.CheckNotificationRateLimit_Response{
		Limited:           state.Limited,
		RetryAfterSeconds: retryAfterSeconds,
		RemainingCount:    state.RemainingCount,
	}, nil
}
