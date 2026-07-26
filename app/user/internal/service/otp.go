package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"strings"
	"user/internal/biz/usecase"
	"user/internal/config"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OtpService struct {
	v1.UnimplementedOtpServiceServer
	conf            *config.Bootstrap
	totpUsecase     *usecase.TotpUsecase
	emailOtpUsecase *usecase.EmailOtpUsecase
	smsOtpUsecase   *usecase.SmsOtpUsecase
}

func NewOtpService(
	conf *config.Bootstrap,
	totpUsecase *usecase.TotpUsecase,
	emailOtpUsecase *usecase.EmailOtpUsecase,
	smsOtpUsecase *usecase.SmsOtpUsecase,
) *OtpService {
	return &OtpService{
		conf:            conf,
		totpUsecase:     totpUsecase,
		emailOtpUsecase: emailOtpUsecase,
		smsOtpUsecase:   smsOtpUsecase,
	}
}

func (s *OtpService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterOtpServiceServer(gs, s)
}

func (s *OtpService) RegisterHttp(hs *http.Server) {
}

func (s *OtpService) ValidateTotp(ctx context.Context, req *v1.ValidateTotp_Req) (*v1.ValidateTotp_Resp, error) {
	if req == nil || req.GetUserId() == 0 || strings.TrimSpace(req.GetCode()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	verified, err := s.totpUsecase.ValidateByUserID(ctx, &usecase.ValidateTotpByUserIDReq{
		UserID: req.GetUserId(),
		Code:   strings.TrimSpace(req.GetCode()),
	})
	if err != nil {
		return nil, err
	}
	return &v1.ValidateTotp_Resp{
		Verified: verified,
	}, nil
}

func (s *OtpService) BeginEnableTotp(ctx context.Context, req *v1.BeginEnableTotp_Req) (*v1.BeginEnableTotp_Resp, error) {
	if req == nil || req.GetUserId() == 0 || strings.TrimSpace(req.GetAccountName()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := s.totpUsecase.BeginEnable(ctx, &usecase.BeginEnableTotpReq{
		UserID:      req.GetUserId(),
		AccountName: strings.TrimSpace(req.GetAccountName()),
	})
	if err != nil {
		return nil, err
	}
	return &v1.BeginEnableTotp_Resp{
		Url:    resp.URL,
		QrCode: resp.QRCode,
	}, nil
}

func (s *OtpService) ConfirmEnableTotp(ctx context.Context, req *v1.ConfirmEnableTotp_Req) (*v1.ConfirmEnableTotp_Resp, error) {
	if req == nil || req.GetUserId() == 0 || strings.TrimSpace(req.GetCode()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.totpUsecase.ConfirmEnable(ctx, &usecase.ConfirmEnableTotpReq{
		UserID: req.GetUserId(),
		Code:   strings.TrimSpace(req.GetCode()),
	})
	return &v1.ConfirmEnableTotp_Resp{}, err
}

func (s *OtpService) DisableTotp(ctx context.Context, req *v1.DisableTotp_Req) (*v1.DisableTotp_Resp, error) {
	if req == nil || req.GetUserId() == 0 || strings.TrimSpace(req.GetCode()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.totpUsecase.Disable(ctx, &usecase.DisableTotpReq{
		UserID: req.GetUserId(),
		Code:   strings.TrimSpace(req.GetCode()),
	})
	return &v1.DisableTotp_Resp{}, err
}

func (s *OtpService) GetTotp(ctx context.Context, req *v1.GetTotp_Req) (*v1.GetTotp_Resp, error) {
	if req == nil || req.GetUserId() == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := s.totpUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	totp := &v1.GetTotp_Resp_Totp{
		UserId: req.GetUserId(),
	}
	if resp != nil {
		totp.Enable = resp.Enable
		if resp.EnableTime != nil {
			totp.EnableTime = timestamppb.New(*resp.EnableTime)
		}
	}
	return &v1.GetTotp_Resp{
		Totp: totp,
	}, nil
}

func (s *OtpService) SendEmailOtp(ctx context.Context, req *v1.SendEmailOtp_Req) (*v1.SendEmailOtp_Resp, error) {
	if req == nil || strings.TrimSpace(req.GetEmail()) == "" || (req.UserId != nil && req.GetUserId() == 0) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := s.emailOtpUsecase.SendEmailOtp(ctx, &usecase.SendEmailOtpReq{
		UserID: req.UserId,
		Email:  req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.SendEmailOtp_Resp{}
	if resp != nil && resp.Code != "" && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}

func (s *OtpService) VerifyEmailOtp(ctx context.Context, req *v1.VerifyEmailOtp_Req) (*v1.VerifyEmailOtp_Resp, error) {
	if req == nil || strings.TrimSpace(req.GetEmail()) == "" || strings.TrimSpace(req.GetCode()) == "" || (req.UserId != nil && req.GetUserId() == 0) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.emailOtpUsecase.VerifyEmailOtp(ctx, &usecase.VerifyEmailOtpReq{
		UserID: req.UserId,
		Email:  req.GetEmail(),
		Code:   req.GetCode(),
	})
	return &v1.VerifyEmailOtp_Resp{}, err
}

func (s *OtpService) SendPhoneOtp(ctx context.Context, req *v1.SendPhoneOtp_Req) (*v1.SendPhoneOtp_Resp, error) {
	if req == nil || strings.TrimSpace(req.GetPhone()) == "" || (req.UserId != nil && req.GetUserId() == 0) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := s.smsOtpUsecase.SendPhoneOtp(ctx, &usecase.SendPhoneOtpReq{
		UserID: req.UserId,
		Phone:  req.GetPhone(),
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.SendPhoneOtp_Resp{}
	if resp != nil && resp.Code != "" && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}

func (s *OtpService) VerifyPhoneOtp(ctx context.Context, req *v1.VerifyPhoneOtp_Req) (*v1.VerifyPhoneOtp_Resp, error) {
	if req == nil || strings.TrimSpace(req.GetPhone()) == "" || strings.TrimSpace(req.GetCode()) == "" || (req.UserId != nil && req.GetUserId() == 0) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.smsOtpUsecase.VerifyPhoneOtp(ctx, &usecase.VerifyPhoneOtpReq{
		UserID: req.UserId,
		Phone:  req.GetPhone(),
		Code:   req.GetCode(),
	})
	return &v1.VerifyPhoneOtp_Resp{}, err
}
