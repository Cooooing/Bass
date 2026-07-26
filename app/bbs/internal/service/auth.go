package service

import (
	"bbs/internal/biz/usecase"
	"bbs/internal/enum"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	bbsuserv1enum "common/proto/gen/bbs/v1/user/enum"
	cerrors "common/proto/gen/common/errors"
	"context"
	"net/mail"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	bbsuserv1.UnimplementedAuthServiceServer
	authUsecase *usecase.AuthUsecase
	phoneRe     *regexp.Regexp
	nameRe      *regexp.Regexp
	codeRe      *regexp.Regexp
}

func NewAuthService(
	authUsecase *usecase.AuthUsecase,
) *AuthService {
	return &AuthService{
		authUsecase: authUsecase,
		phoneRe:     regexp.MustCompile(`^1[3-9]\d{9}$`),
		nameRe:      regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`),
		codeRe:      regexp.MustCompile(`^[A-Za-z0-9]{6}$`),
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
}

func (s *AuthService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAuthServiceHTTPServer(hs, s)
}

func (s *AuthService) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Req) (*bbsuserv1.StartEmailRegistration_Resp, error) {
	if err := s.validateStartEmailRegistration(req); err != nil {
		return nil, err
	}
	resp, err := s.authUsecase.StartEmailRegistration(ctx, &usecase.StartEmailRegistrationReq{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	reply := &bbsuserv1.StartEmailRegistration_Resp{}
	if resp.Code != "" {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}

func (s *AuthService) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Req) (*bbsuserv1.VerifyEmailRegistration_Resp, error) {
	if err := s.validateVerifyEmailRegistration(req); err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyEmailRegistration_Resp{}, s.authUsecase.VerifyEmailRegistration(ctx, &usecase.VerifyEmailRegistrationReq{
		Email: req.GetEmail(),
		Code:  req.GetCode(),
	})
}

func (s *AuthService) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Req) (*bbsuserv1.StartPhoneRegistration_Resp, error) {
	if err := s.validateStartPhoneRegistration(req); err != nil {
		return nil, err
	}
	resp, err := s.authUsecase.StartPhoneRegistration(ctx, &usecase.StartPhoneRegistrationReq{
		Phone:    req.GetPhone(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	reply := &bbsuserv1.StartPhoneRegistration_Resp{}
	if resp.Code != "" {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}

func (s *AuthService) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Req) (*bbsuserv1.VerifyPhoneRegistration_Resp, error) {
	if err := s.validateVerifyPhoneRegistration(req); err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyPhoneRegistration_Resp{}, s.authUsecase.VerifyPhoneRegistration(ctx, &usecase.VerifyPhoneRegistrationReq{
		Phone: req.GetPhone(),
		Code:  req.GetCode(),
	})
}

func (s *AuthService) StartEmailLogin(ctx context.Context, req *bbsuserv1.StartEmailLogin_Req) (*bbsuserv1.StartEmailLogin_Resp, error) {
	email, err := s.normalizeEmail(req.GetEmail())
	if err != nil {
		return nil, err
	}
	resp, err := s.authUsecase.StartEmailLogin(ctx, email)
	if err != nil {
		return nil, err
	}
	reply := &bbsuserv1.StartEmailLogin_Resp{}
	if resp.Code != "" {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}

func (s *AuthService) StartPhoneLogin(ctx context.Context, req *bbsuserv1.StartPhoneLogin_Req) (*bbsuserv1.StartPhoneLogin_Resp, error) {
	phone, err := s.normalizePhone(req.GetPhone())
	if err != nil {
		return nil, err
	}
	resp, err := s.authUsecase.StartPhoneLogin(ctx, phone)
	if err != nil {
		return nil, err
	}
	reply := &bbsuserv1.StartPhoneLogin_Resp{}
	if resp.Code != "" {
		reply.Code = new(resp.Code)
	}
	return reply, nil
}

func (s *AuthService) Login(ctx context.Context, req *bbsuserv1.Login_Req) (*bbsuserv1.Login_Resp, error) {
	ucReq, err := s.validateLogin(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.authUsecase.Login(ctx, ucReq)
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.Login_Resp_Account
	if resp.Account != nil {
		account = &bbsuserv1.Login_Resp_Account{}
		if profile := resp.Account.Profile; profile != nil {
			account.Basic = &bbsuserv1.Login_Resp_AccountBasic{
				Id:            profile.ID,
				Name:          profile.Name,
				Nickname:      profile.Nickname,
				Url:           profile.URL,
				AvatarUrl:     profile.AvatarURL,
				Introduction:  profile.Introduction,
				Status:        bbsuserv1enum.AccountStatus(profile.Status),
				Mbti:          bbsuserv1enum.MBTI(profile.MBTI),
				FollowCount:   profile.FollowCount,
				FollowerCount: profile.FollowerCount,
			}
			if profile.CreatedAt != nil {
				account.Basic.CreatedAt = timestamppb.New(*profile.CreatedAt)
			}
			if profile.UpdatedAt != nil {
				account.Basic.UpdatedAt = timestamppb.New(*profile.UpdatedAt)
			}
		}
		if contact := resp.Account.Contact; contact != nil {
			account.Contact = &bbsuserv1.Login_Resp_AccountContact{
				UserId: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &bbsuserv1.Login_Resp{
		AccessToken:           resp.AccessToken,
		RefreshToken:          resp.RefreshToken,
		AccessTokenExpiresAt:  timestamppb.New(resp.AccessTokenExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(resp.RefreshTokenExpiresAt),
		SessionExpiresAt:      timestamppb.New(resp.SessionExpiresAt),
		Account:               account,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *bbsuserv1.RefreshToken_Req) (*bbsuserv1.RefreshToken_Resp, error) {
	token := strings.TrimSpace(req.GetRefreshToken())
	if token == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := s.authUsecase.RefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.RefreshToken_Resp{
		AccessToken:           resp.AccessToken,
		RefreshToken:          resp.RefreshToken,
		AccessTokenExpiresAt:  timestamppb.New(resp.AccessTokenExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(resp.RefreshTokenExpiresAt),
		SessionExpiresAt:      timestamppb.New(resp.SessionExpiresAt),
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *bbsuserv1.Logout_Req) (*bbsuserv1.Logout_Resp, error) {
	token, ok := util.GetContextValue[string](ctx, constant.CtxToken)
	if !ok || token == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return &bbsuserv1.Logout_Resp{}, s.authUsecase.Logout(ctx, token)
}

func (s *AuthService) CancelAccount(ctx context.Context, req *bbsuserv1.CancelAccount_Req) (*bbsuserv1.CancelAccount_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if strings.TrimSpace(req.GetPassword()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return &bbsuserv1.CancelAccount_Resp{}, s.authUsecase.CancelAccount(ctx, user.ID, req.GetPassword(), req.GetCode())
}

func (s *AuthService) validateLogin(req *bbsuserv1.Login_Req) (*usecase.LoginReq, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	loginType, ok := enum.LoginTypeMap.ToEnum(req.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	out := &usecase.LoginReq{
		Type: loginType,
	}
	switch loginType {
	case enum.LoginTypePassword:
		cred := req.GetPasswordCredential()
		if cred == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		account, err := s.normalizeLoginAccount(cred.GetAccount())
		if err != nil {
			return nil, err
		}
		if strings.TrimSpace(cred.GetPassword()) == "" {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		out.Account, out.Password, out.Code = account, cred.GetPassword(), strings.TrimSpace(cred.GetCode())
	case enum.LoginTypeEmail:
		cred := req.GetEmailCredential()
		if cred == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		email, err := s.normalizeEmail(cred.GetEmail())
		if err != nil {
			return nil, err
		}
		code, err := s.normalizeCode(cred.GetCode())
		if err != nil {
			return nil, err
		}
		out.Email, out.Code = email, code
	case enum.LoginTypePhone:
		cred := req.GetPhoneCredential()
		if cred == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		phone, err := s.normalizePhone(cred.GetPhone())
		if err != nil {
			return nil, err
		}
		code, err := s.normalizeCode(cred.GetCode())
		if err != nil {
			return nil, err
		}
		out.Phone, out.Code = phone, code
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return out, nil
}

func (s *AuthService) validateStartEmailRegistration(req *bbsuserv1.StartEmailRegistration_Req) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	email, err := s.normalizeEmail(req.GetEmail())
	if err != nil {
		return err
	}
	name, err := s.normalizeName(req.GetName())
	if err != nil {
		return err
	}
	nickname, err := s.normalizeNickname(req.Nickname)
	if err != nil {
		return err
	}
	if err := s.validatePassword(req.GetPassword()); err != nil {
		return err
	}
	req.Email, req.Name, req.Nickname = email, name, nickname
	return nil
}

func (s *AuthService) validateStartPhoneRegistration(req *bbsuserv1.StartPhoneRegistration_Req) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	phone, err := s.normalizePhone(req.GetPhone())
	if err != nil {
		return err
	}
	name, err := s.normalizeName(req.GetName())
	if err != nil {
		return err
	}
	nickname, err := s.normalizeNickname(req.Nickname)
	if err != nil {
		return err
	}
	if err := s.validatePassword(req.GetPassword()); err != nil {
		return err
	}
	req.Phone, req.Name, req.Nickname = phone, name, nickname
	return nil
}

func (s *AuthService) validateVerifyEmailRegistration(req *bbsuserv1.VerifyEmailRegistration_Req) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	email, err := s.normalizeEmail(req.GetEmail())
	if err != nil {
		return err
	}
	code, err := s.normalizeCode(req.GetCode())
	if err != nil {
		return err
	}
	req.Email, req.Code = email, code
	return nil
}

func (s *AuthService) validateVerifyPhoneRegistration(req *bbsuserv1.VerifyPhoneRegistration_Req) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	phone, err := s.normalizePhone(req.GetPhone())
	if err != nil {
		return err
	}
	code, err := s.normalizeCode(req.GetCode())
	if err != nil {
		return err
	}
	req.Phone, req.Code = phone, code
	return nil
}

func (s *AuthService) normalizeEmail(email string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || utf8.RuneCountInString(email) > 254 {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	parsed, err := mail.ParseAddress(email)
	if err != nil || parsed.Address != email || !strings.Contains(email, "@") {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return email, nil
}

func (s *AuthService) normalizePhone(phone string) (string, error) {
	phone = strings.TrimSpace(phone)
	if !s.phoneRe.MatchString(phone) {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return phone, nil
}

func (s *AuthService) normalizeName(name string) (string, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	length := utf8.RuneCountInString(name)
	if length < 4 || length > 32 || !s.nameRe.MatchString(name) {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return name, nil
}

func (s *AuthService) normalizeNickname(nickname *string) (*string, error) {
	if nickname == nil {
		return nil, nil
	}
	value := strings.TrimSpace(*nickname)
	length := utf8.RuneCountInString(value)
	if length < 2 || length > 32 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	for _, r := range value {
		if !unicode.IsDigit(r) {
			return new(value), nil
		}
	}
	return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
}

func (s *AuthService) validatePassword(password string) error {
	if len(password) < 6 || len(password) > 64 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range password {
		if r < '!' || r > '~' {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
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
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return nil
}

func (s *AuthService) normalizeCode(code string) (string, error) {
	code = strings.TrimSpace(code)
	if !s.codeRe.MatchString(code) {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return code, nil
}

func (s *AuthService) normalizeLoginAccount(account string) (string, error) {
	account = strings.TrimSpace(account)
	if account == "" {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if strings.Contains(account, "@") {
		return s.normalizeEmail(account)
	}
	if s.phoneRe.MatchString(account) {
		return account, nil
	}
	return s.normalizeName(account)
}
