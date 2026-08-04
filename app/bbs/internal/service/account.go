package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	bbsuserv1enum "common/proto/gen/bbs/v1/user/enum"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"net/mail"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type AccountService struct {
	bbsuserv1.UnimplementedAccountServiceServer
	accountUsecase *usecase.AccountUsecase
	phoneRe        *regexp.Regexp
	codeRe         *regexp.Regexp
}

func NewAccountService(
	accountUsecase *usecase.AccountUsecase,
) *AccountService {
	return &AccountService{
		accountUsecase: accountUsecase,
		phoneRe:        regexp.MustCompile("^1[3-9]\\d{9}$"),
		codeRe:         regexp.MustCompile("^[A-Za-z0-9]{6}$"),
	}
}
func (s *AccountService) RegisterGrpc(gs *grpc.Server) {
}

func (s *AccountService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAccountServiceHTTPServer(hs, s)
}

func (s *AccountService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Req) (*bbsuserv1.GetCurrentAccount_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	account, err := s.accountUsecase.GetCurrentAccount(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentAccount_Resp{
		Account: account,
	}, nil
}

func (s *AccountService) GetProfile(ctx context.Context, req *bbsuserv1.GetProfileAccount_Req) (*bbsuserv1.GetProfileAccount_Resp, error) {
	profile, err := s.accountUsecase.GetProfileAccount(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetProfileAccount_Resp{
		Profile: profile,
	}, nil
}

func (s *AccountService) UpdateProfile(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Req) (*bbsuserv1.UpdateProfileAccount_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
	}

	if req.Nickname != nil {
		value := strings.TrimSpace(*req.Nickname)
		if value != "" {
			length := utf8.RuneCountInString(value)
			if length < 2 || length > 32 {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
			}
			hasNonDigit := false
			for _, r := range value {
				if !unicode.IsDigit(r) {
					hasNonDigit = true
					break
				}
			}
			if !hasNonDigit {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
			}
		}
		req.Nickname = new(value)
	}
	if req.Url != nil {
		value := strings.TrimSpace(*req.Url)
		if utf8.RuneCountInString(value) > 2048 {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
		}
		req.Url = new(value)
	}
	if req.Introduction != nil && utf8.RuneCountInString(*req.Introduction) > 512 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
	}
	if req.Mbti != nil {
		if _, ok := bbsuserv1enum.MBTI_name[int32(*req.Mbti)]; !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
		}
	}
	profile, err := s.accountUsecase.UpdateProfileAccount(ctx, &usecase.UpdateProfileAccountReq{
		UserID:        user.ID,
		AvatarAssetID: req.AvatarAssetId,
		Nickname:      req.Nickname,
		URL:           req.Url,
		Introduction:  req.Introduction,
		Mbti:          req.Mbti,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdateProfileAccount_Resp{
		Profile: profile,
	}, nil
}

func (s *AccountService) UpdatePassword(ctx context.Context, req *bbsuserv1.UpdatePasswordAccount_Req) (*bbsuserv1.UpdatePasswordAccount_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	oldPassword := req.GetOldPassword()
	newPassword := req.GetNewPassword()
	if len(oldPassword) < 6 || len(oldPassword) > 64 || len(newPassword) < 6 || len(newPassword) > 64 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range newPassword {
		if r < '!' || r > '~' {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		switch {
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		case r >= '0' && r <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.accountUsecase.UpdatePasswordAccount(ctx, &usecase.UpdatePasswordAccountReq{
		UserID:      user.ID,
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}); err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdatePasswordAccount_Resp{}, nil
}

func (s *AccountService) UpdateEmail(ctx context.Context, req *bbsuserv1.UpdateEmailAccount_Req) (*bbsuserv1.UpdateEmailAccount_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	email := strings.ToLower(strings.TrimSpace(req.GetEmail()))
	parsed, err := mail.ParseAddress(email)
	if err != nil || parsed.Address != email || !strings.Contains(email, "@") || utf8.RuneCountInString(email) > 254 || !s.codeRe.MatchString(strings.TrimSpace(req.GetCode())) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.accountUsecase.UpdateEmailAccount(ctx, &usecase.UpdateEmailAccountReq{
		UserID: user.ID,
		Email:  email,
		Code:   strings.TrimSpace(req.GetCode()),
	}); err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdateEmailAccount_Resp{}, nil
}

func (s *AccountService) UpdatePhone(ctx context.Context, req *bbsuserv1.UpdatePhoneAccount_Req) (*bbsuserv1.UpdatePhoneAccount_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	phone := strings.TrimSpace(req.GetPhone())
	if !s.phoneRe.MatchString(phone) || !s.codeRe.MatchString(strings.TrimSpace(req.GetCode())) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.accountUsecase.UpdatePhoneAccount(ctx, &usecase.UpdatePhoneAccountReq{
		UserID: user.ID,
		Phone:  phone,
		Code:   strings.TrimSpace(req.GetCode()),
	}); err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdatePhoneAccount_Resp{}, nil
}
func (s *AccountService) Avatar(ctx context.Context, req *bbsuserv1.AvatarAccount_Req) (*common.ImageResp, error) {
	resp, err := s.accountUsecase.AvatarAccount(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return &common.ImageResp{
		Data:        resp.Data,
		ContentType: resp.ContentType,
	}, nil
}
