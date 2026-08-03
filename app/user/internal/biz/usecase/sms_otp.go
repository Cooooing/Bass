package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"strings"
	"user/internal/biz/repo"
)

type SmsOtpUsecase struct {
	notificationRateLimitClient repo.NotificationRateLimitClient
}

func NewSmsOtpUsecase(
	notificationRateLimitClient repo.NotificationRateLimitClient,
) *SmsOtpUsecase {
	return &SmsOtpUsecase{
		notificationRateLimitClient: notificationRateLimitClient,
	}
}

type SendPhoneOtpReq struct {
	UserID *int64
	Phone  string
}

type SendPhoneOtpResp struct {
	Code string
}

func (u *SmsOtpUsecase) SendPhoneOtp(ctx context.Context, req *SendPhoneOtpReq) (*SendPhoneOtpResp, error) {
	rateLimitState, err := u.notificationRateLimitClient.CheckPhone(ctx, strings.TrimSpace(req.Phone))
	if err != nil {
		return nil, err
	}
	if rateLimitState != nil && rateLimitState.Limited {
		return nil, apperror.New(
			cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_SEND_TOO_FREQUENT,
			apperror.WithData(&cerrors.RetryAfterErrorData{
				RetryAfterSeconds: rateLimitState.RetryAfterSeconds,
			}),
		)
	}
	return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
}

type VerifyPhoneOtpReq struct {
	UserID *int64
	Phone  string
	Code   string
}

func (u *SmsOtpUsecase) VerifyPhoneOtp(ctx context.Context, req *VerifyPhoneOtpReq) error {
	return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
}
