package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"strings"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/usecase"
	"user/internal/config"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer
	conf        *config.Bootstrap
	authUsecase *usecase.AuthUsecase
}

func NewAuthService(
	conf *config.Bootstrap,
	authUsecase *usecase.AuthUsecase,
) *AuthService {
	return &AuthService{
		conf:        conf,
		authUsecase: authUsecase,
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAuthServiceServer(gs, s)
}

func (s *AuthService) RegisterHttp(hs *http.Server) {
}

func (s *AuthService) StartEmailRegistration(ctx context.Context, req *v1.StartEmailRegistration_Req) (*v1.StartEmailRegistration_Resp, error) {
	res, err := s.authUsecase.StartEmailRegistration(ctx, &model.Account{
		Email:    &req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	reply := &v1.StartEmailRegistration_Resp{}
	if err == nil && res != nil && res.Code != "" && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = new(res.Code)
	}
	return reply, err
}

func (s *AuthService) VerifyEmailRegistration(ctx context.Context, req *v1.VerifyEmailRegistration_Req) (*v1.VerifyEmailRegistration_Resp, error) {
	err := s.authUsecase.VerifyEmailRegistration(ctx, &usecase.VerifyEmailRegistrationReq{
		Email: req.GetEmail(),
		Code:  req.GetCode(),
	})
	return &v1.VerifyEmailRegistration_Resp{}, err
}

func (s *AuthService) StartPhoneRegistration(ctx context.Context, req *v1.StartPhoneRegistration_Req) (*v1.StartPhoneRegistration_Resp, error) {
	res, err := s.authUsecase.StartPhoneRegistration(ctx, &model.Account{
		Phone:    &req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	reply := &v1.StartPhoneRegistration_Resp{}
	if err == nil && res != nil && res.Code != "" && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = new(res.Code)
	}
	return reply, err
}

func (s *AuthService) VerifyPhoneRegistration(ctx context.Context, req *v1.VerifyPhoneRegistration_Req) (*v1.VerifyPhoneRegistration_Resp, error) {
	err := s.authUsecase.VerifyPhoneRegistration(ctx, &usecase.VerifyPhoneRegistrationReq{
		Phone: req.GetPhone(),
		Code:  req.GetCode(),
	})
	return &v1.VerifyPhoneRegistration_Resp{}, err
}

func (s *AuthService) Login(ctx context.Context, req *v1.Login_Req) (*v1.Login_Resp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	loginType, ok := enum.LoginTypeMap.ToEnum(req.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	client := req.GetClient()
	clientContext := &model.LoginContext{
		ClientType: enum.ClientTypeUnknown,
		DeviceType: enum.DeviceTypeUnknown,
	}
	if client != nil {
		clientType, ok := enum.ClientTypeMap.ToEnum(client.GetClientType())
		if !ok {
			clientType = enum.ClientTypeUnknown
		}
		deviceType, ok := enum.DeviceTypeMap.ToEnum(client.GetDeviceType())
		if !ok {
			deviceType = enum.DeviceTypeUnknown
		}
		clientContext = &model.LoginContext{
			IP:             client.GetIp(),
			UserAgent:      client.GetUserAgent(),
			ClientType:     clientType,
			DeviceType:     deviceType,
			OSName:         client.GetOsName(),
			OSVersion:      client.GetOsVersion(),
			BrowserName:    client.GetBrowserName(),
			BrowserVersion: client.GetBrowserVersion(),
			AppName:        client.GetAppName(),
			AppVersion:     client.GetAppVersion(),
		}
	}
	ucReq := &usecase.LoginReq{
		Type:   loginType,
		Realm:  realm,
		Client: clientContext,
	}
	switch loginType {
	case enum.LoginTypePassword:
		cred := req.GetPasswordCredential()
		if cred == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		ucReq.PasswordAccount = cred.GetAccount()
		ucReq.Password = cred.GetPassword()
		ucReq.Code = cred.GetCode()
	case enum.LoginTypeEmail:
		cred := req.GetEmailCredential()
		if cred == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		ucReq.Email = cred.GetEmail()
		ucReq.Code = cred.GetCode()
	case enum.LoginTypePhone:
		cred := req.GetPhoneCredential()
		if cred == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		ucReq.Phone = cred.GetPhone()
		ucReq.Code = cred.GetCode()
	}
	res, err := s.authUsecase.Login(ctx, ucReq)
	if err != nil {
		return nil, err
	}
	var account *v1.Login_Resp_Account
	if res.Account != nil {
		basic := &v1.Login_Resp_AccountBasic{
			Id:            res.Account.ID,
			Name:          res.Account.Name,
			Nickname:      res.Account.Nickname,
			Url:           res.Account.URL,
			AvatarUrl:     res.Account.AvatarURL,
			Introduction:  res.Account.Introduction,
			FollowCount:   res.Account.FollowCount,
			FollowerCount: res.Account.FollowerCount,
		}
		if res.Account.Mbti != nil {
			basic.Mbti = enum.MBTIMap.MustToProto(*res.Account.Mbti)
		}
		if res.Account.Status != nil {
			basic.Status = enum.AccountStatusMap.MustToProto(*res.Account.Status)
		}
		if res.Account.CreatedAt != nil {
			basic.CreatedAt = timestamppb.New(*res.Account.CreatedAt)
		}
		if res.Account.UpdatedAt != nil {
			basic.UpdatedAt = timestamppb.New(*res.Account.UpdatedAt)
		}
		account = &v1.Login_Resp_Account{
			Basic: basic,
			Contact: &v1.Login_Resp_AccountContact{
				UserId: res.Account.ID,
				Email:  res.Account.Email,
				Phone:  res.Account.Phone,
			},
		}
	}
	reply := &v1.Login_Resp{
		AccessToken:  res.TokenPair.AccessToken,
		RefreshToken: res.TokenPair.RefreshToken,
		Account:      account,
	}
	if res.TokenPair.AccessTokenExpiresAt != nil {
		reply.AccessTokenExpiresAt = timestamppb.New(*res.TokenPair.AccessTokenExpiresAt)
	}
	if res.TokenPair.RefreshTokenExpiresAt != nil {
		reply.RefreshTokenExpiresAt = timestamppb.New(*res.TokenPair.RefreshTokenExpiresAt)
	}
	if res.TokenPair.SessionExpiresAt != nil {
		reply.SessionExpiresAt = timestamppb.New(*res.TokenPair.SessionExpiresAt)
	}
	return reply, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *v1.RefreshToken_Req) (*v1.RefreshToken_Resp, error) {
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	res, err := s.authUsecase.RefreshToken(ctx, req.GetRefreshToken(), realm)
	if err != nil {
		return nil, err
	}
	reply := &v1.RefreshToken_Resp{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
	if res.AccessTokenExpiresAt != nil {
		reply.AccessTokenExpiresAt = timestamppb.New(*res.AccessTokenExpiresAt)
	}
	if res.RefreshTokenExpiresAt != nil {
		reply.RefreshTokenExpiresAt = timestamppb.New(*res.RefreshTokenExpiresAt)
	}
	if res.SessionExpiresAt != nil {
		reply.SessionExpiresAt = timestamppb.New(*res.SessionExpiresAt)
	}
	return reply, nil
}

func (s *AuthService) Logout(ctx context.Context, req *v1.Logout_Req) (*v1.Logout_Resp, error) {
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.authUsecase.Logout(ctx, req.GetAccessToken(), realm)
	return &v1.Logout_Resp{}, err
}

func (s *AuthService) ParseToken(ctx context.Context, req *v1.ParseToken_Req) (*v1.ParseToken_Resp, error) {
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	res, err := s.authUsecase.ParseToken(ctx, req.GetAccessToken(), realm)
	if err != nil {
		return nil, err
	}
	user := res.User
	return &v1.ParseToken_Resp{
		User: &v1.ParseToken_Resp_TokenUser{
			Id:        user.ID,
			Name:      user.Name,
			Nickname:  user.Nickname,
			Language:  user.Language,
			Timezone:  user.Timezone,
			SessionId: res.SessionID,
			Realm:     commonenum.LoginRealmMap.MustToProto(res.Realm),
		},
	}, nil
}

func (s *AuthService) CancelAccount(ctx context.Context, req *v1.CancelAccount_Req) (*v1.CancelAccount_Resp, error) {
	if req == nil || req.GetUserId() == 0 || strings.TrimSpace(req.GetPassword()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.authUsecase.CancelAccount(ctx, &usecase.CancelAccountReq{
		UserID:   req.GetUserId(),
		Password: req.GetPassword(),
		Code:     req.GetCode(),
	})
	return &v1.CancelAccount_Resp{}, err
}

func (s *AuthService) BanAccount(ctx context.Context, req *v1.BanAccount_Req) (*v1.BanAccount_Resp, error) {
	if req == nil || req.GetUserId() == 0 || req.GetOperatorId() == 0 || strings.TrimSpace(req.GetReason()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetOperatorRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var bannedUntil *time.Time
	if req.BannedUntil != nil {
		bannedUntil = new(req.BannedUntil.AsTime())
	}
	res, err := s.authUsecase.BanAccount(ctx, &usecase.BanAccountReq{
		UserID:        req.GetUserId(),
		OperatorID:    req.GetOperatorId(),
		OperatorRealm: realm,
		Reason:        req.GetReason(),
		Remark:        req.GetRemark(),
		BannedUntil:   bannedUntil,
	})
	if err != nil {
		return nil, err
	}
	return &v1.BanAccount_Resp{
		BanRecordId: res.BanRecordID,
	}, nil
}

func (s *AuthService) UnbanAccounts(ctx context.Context, req *v1.UnbanAccounts_Req) (*v1.UnbanAccounts_Resp, error) {
	if req == nil || len(req.GetUserIds()) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.authUsecase.UnbanAccounts(ctx, req.GetUserIds()); err != nil {
		return nil, err
	}
	return &v1.UnbanAccounts_Resp{}, nil
}
