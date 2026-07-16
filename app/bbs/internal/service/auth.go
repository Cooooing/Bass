package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
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

type AuthService struct {
	bbsuserv1.UnimplementedAuthServiceServer
	authUsecase *usecase.AuthUsecase
	phoneRe     *regexp.Regexp
	nameRe      *regexp.Regexp
	codeRe      *regexp.Regexp
}

func NewAuthService(authUsecase *usecase.AuthUsecase) *AuthService {
	return &AuthService{
		authUsecase: authUsecase,
		phoneRe:     regexp.MustCompile(`^1[3-9]\d{9}$`),
		nameRe:      regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`),
		codeRe:      regexp.MustCompile(`^[A-Za-z0-9]{6}$`),
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {}

func (s *AuthService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAuthServiceHTTPServer(hs, s)
}

func (s *AuthService) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Request) (*bbsuserv1.StartEmailRegistration_Response, error) {
	if err := s.validateStartEmailRegistration(req); err != nil {
		return nil, err
	}
	response, err := s.authUsecase.StartEmailRegistration(ctx, &usecase.StartEmailRegistrationReq{Email: req.GetEmail(), Password: req.GetPassword(), Name: req.GetName(), Nickname: req.Nickname})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.StartEmailRegistration_Response{CodeToken: response.CodeToken, Code: response.Code}, nil
}

func (s *AuthService) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Request) (*bbsuserv1.VerifyEmailRegistration_Response, error) {
	if err := s.validateVerifyEmailRegistration(req); err != nil {
		return nil, err
	}
	err := s.authUsecase.VerifyEmailRegistration(ctx, &usecase.VerifyEmailRegistrationReq{Code: req.GetCode(), CodeToken: req.GetCodeToken()})
	return &bbsuserv1.VerifyEmailRegistration_Response{}, err
}

func (s *AuthService) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Request) (*bbsuserv1.StartPhoneRegistration_Response, error) {
	if err := s.validateStartPhoneRegistration(req); err != nil {
		return nil, err
	}
	response, err := s.authUsecase.StartPhoneRegistration(ctx, &usecase.StartPhoneRegistrationReq{Phone: req.GetPhone(), Password: req.GetPassword(), Name: req.GetName(), Nickname: req.Nickname})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.StartPhoneRegistration_Response{CodeToken: response.CodeToken, Code: response.Code}, nil
}

func (s *AuthService) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Request) (*bbsuserv1.VerifyPhoneRegistration_Response, error) {
	if err := s.validateVerifyPhoneRegistration(req); err != nil {
		return nil, err
	}
	err := s.authUsecase.VerifyPhoneRegistration(ctx, &usecase.VerifyPhoneRegistrationReq{Code: req.GetCode(), CodeToken: req.GetCodeToken()})
	return &bbsuserv1.VerifyPhoneRegistration_Response{}, err
}

func (s *AuthService) LoginByPassword(ctx context.Context, req *bbsuserv1.LoginByPassword_Request) (*bbsuserv1.LoginByPassword_Response, error) {
	if err := s.validateLoginByPassword(req); err != nil {
		return nil, err
	}
	response, err := s.authUsecase.LoginByPassword(ctx, &usecase.LoginByPasswordReq{Account: req.GetAccount(), Password: req.GetPassword()})
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.LoginByPassword_Response_Account
	if response.Account != nil {
		account = &bbsuserv1.LoginByPassword_Response_Account{}
		if profile := response.Account.Basic; profile != nil {
			account.Basic = &bbsuserv1.LoginByPassword_Response_AccountBasic{Id: profile.Id, Name: profile.Name, Nickname: profile.Nickname, Url: profile.Url, AvatarUrl: profile.AvatarUrl, Introduction: profile.Introduction, Status: bbsuserv1.AccountStatus(profile.Status), Mbti: bbsuserv1.MBTI(profile.Mbti), FollowCount: profile.FollowCount, FollowerCount: profile.FollowerCount, CreatedAt: profile.CreatedAt, UpdatedAt: profile.UpdatedAt}
		}
		if contact := response.Account.Contact; contact != nil {
			account.Contact = &bbsuserv1.LoginByPassword_Response_AccountContact{UserId: contact.UserId, Email: contact.Email, Phone: contact.Phone}
		}
	}
	return &bbsuserv1.LoginByPassword_Response{Token: response.Token, Account: account}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Response, error) {
	token, err := currentToken(ctx)
	if err != nil {
		return nil, err
	}
	err = s.authUsecase.Logout(ctx, &usecase.LogoutReq{Token: token})
	return &bbsuserv1.Logout_Response{}, err
}

func (s *AuthService) validateStartEmailRegistration(req *bbsuserv1.StartEmailRegistration_Request) error {
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
	req.Email = email
	req.Name = name
	req.Nickname = nickname
	return nil
}

func (s *AuthService) validateStartPhoneRegistration(req *bbsuserv1.StartPhoneRegistration_Request) error {
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
	req.Phone = phone
	req.Name = name
	req.Nickname = nickname
	return nil
}

func (s *AuthService) validateVerifyEmailRegistration(req *bbsuserv1.VerifyEmailRegistration_Request) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	code, codeToken, err := s.normalizeVerification(req.GetCode(), req.GetCodeToken())
	if err != nil {
		return err
	}
	req.Code = code
	req.CodeToken = codeToken
	return nil
}

func (s *AuthService) validateVerifyPhoneRegistration(req *bbsuserv1.VerifyPhoneRegistration_Request) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	code, codeToken, err := s.normalizeVerification(req.GetCode(), req.GetCodeToken())
	if err != nil {
		return err
	}
	req.Code = code
	req.CodeToken = codeToken
	return nil
}

func (s *AuthService) validateLoginByPassword(req *bbsuserv1.LoginByPassword_Request) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	account, err := s.normalizeLoginAccount(req.GetAccount())
	if err != nil {
		return err
	}
	if strings.TrimSpace(req.GetPassword()) == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	req.Account = account
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
	if len(password) < 8 || len(password) > 64 {
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

func (s *AuthService) normalizeVerification(code string, codeToken string) (string, string, error) {
	code = strings.TrimSpace(code)
	codeToken = strings.TrimSpace(codeToken)
	if !s.codeRe.MatchString(code) {
		return "", "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if codeToken == "" {
		return "", "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return code, codeToken, nil
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
