package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	"context"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OtpService struct {
	bbsuserv1.UnimplementedOtpServiceServer
	otpUsecase *usecase.OtpUsecase
	phoneRe    *regexp.Regexp
	codeRe     *regexp.Regexp
}

func NewOtpService(
	otpUsecase *usecase.OtpUsecase,
) *OtpService {
	return &OtpService{
		otpUsecase: otpUsecase,
		phoneRe:    regexp.MustCompile(`^1[3-9]\d{9}$`),
		codeRe:     regexp.MustCompile(`^[A-Za-z0-9]{6}$`),
	}
}

func (s *OtpService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterOtpServiceHTTPServer(hs, s)
}

func (s *OtpService) RegisterGrpc(gs *grpc.Server) {
}

func (s *OtpService) BeginEnableTotp(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Req) (*bbsuserv1.BeginEnableTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.otpUsecase.BeginEnableTotp(ctx, &usecase.BeginEnableTotpReq{
		UserID:      user.ID,
		AccountName: user.Name,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BeginEnableTotp_Resp{
		Url:    resp.URL,
		QrCode: resp.QRCode,
	}, nil
}

func (s *OtpService) ConfirmEnableTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Req) (*bbsuserv1.ConfirmEnableTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	code := strings.TrimSpace(req.GetCode())
	if !s.codeRe.MatchString(code) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.otpUsecase.ConfirmEnableTotp(ctx, &usecase.ConfirmEnableTotpReq{
		UserID: user.ID,
		Code:   code,
	})
	return &bbsuserv1.ConfirmEnableTotp_Resp{}, err
}

func (s *OtpService) DisableTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Req) (*bbsuserv1.DisableTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	code := strings.TrimSpace(req.GetCode())
	if !s.codeRe.MatchString(code) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.otpUsecase.DisableTotp(ctx, &usecase.DisableTotpReq{
		UserID: user.ID,
		Code:   code,
	})
	return &bbsuserv1.DisableTotp_Resp{}, err
}

func (s *OtpService) GetCurrentTotp(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Req) (*bbsuserv1.GetCurrentTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	totp, err := s.otpUsecase.GetCurrentTotp(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	var replyTotp *bbsuserv1.GetCurrentTotp_Resp_Totp
	if totp != nil {
		replyTotp = &bbsuserv1.GetCurrentTotp_Resp_Totp{
			UserId: totp.UserID,
			Enable: totp.Enable,
		}
		if totp.EnableTime != nil {
			replyTotp.EnableTime = timestamppb.New(*totp.EnableTime)
		}
	}
	return &bbsuserv1.GetCurrentTotp_Resp{
		Totp: replyTotp,
	}, nil
}

func (s *OtpService) SendEmailOtp(ctx context.Context, req *bbsuserv1.SendEmailOtp_Req) (*bbsuserv1.SendEmailOtp_Resp, error) {
	email := strings.ToLower(strings.TrimSpace(req.GetEmail()))
	parsed, err := mail.ParseAddress(email)
	if email == "" || utf8.RuneCountInString(email) > 254 || err != nil || parsed.Address != email || !strings.Contains(email, "@") {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var userID *int64
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		userID = new(user.ID)
	}
	resp, err := s.otpUsecase.SendEmailOtp(ctx, &usecase.SendEmailOtpReq{
		UserID: userID,
		Email:  email,
	})
	if err != nil {
		return nil, err
	}
	reply := &bbsuserv1.SendEmailOtp_Resp{}
	if resp.Code != "" {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}

func (s *OtpService) SendPhoneOtp(ctx context.Context, req *bbsuserv1.SendPhoneOtp_Req) (*bbsuserv1.SendPhoneOtp_Resp, error) {
	phone := strings.TrimSpace(req.GetPhone())
	if !s.phoneRe.MatchString(phone) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var userID *int64
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		userID = new(user.ID)
	}
	resp, err := s.otpUsecase.SendPhoneOtp(ctx, &usecase.SendPhoneOtpReq{
		UserID: userID,
		Phone:  phone,
	})
	if err != nil {
		return nil, err
	}
	reply := &bbsuserv1.SendPhoneOtp_Resp{}
	if resp.Code != "" {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}
