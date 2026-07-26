package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type SmsOtpUsecase struct{}

func NewSmsOtpUsecase() *SmsOtpUsecase {
	return &SmsOtpUsecase{}
}

type SendPhoneOtpReq struct {
	UserID *int64
	Phone  string
}

type SendPhoneOtpResp struct {
	Code string
}

func (u *SmsOtpUsecase) SendPhoneOtp(ctx context.Context, req *SendPhoneOtpReq) (*SendPhoneOtpResp, error) {
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
