package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	cerrors "common/api/gen/common/errors"
	"net/mail"
	"strings"
	"unicode"
	"unicode/utf8"
)

func (s *AuthService) validateStartEmailRegistration(req *bbsuserv1.StartEmailRegistration_Request) error {
	if req == nil {
		return cerrors.ErrorBadRequest("registration request is invalid")
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
		return cerrors.ErrorBadRequest("registration request is invalid")
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
		return cerrors.ErrorBadRequest("verification request is invalid")
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
		return cerrors.ErrorBadRequest("verification request is invalid")
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
		return cerrors.ErrorBadRequest("login request is invalid")
	}
	account, err := s.normalizeLoginAccount(req.GetAccount())
	if err != nil {
		return err
	}
	if strings.TrimSpace(req.GetPassword()) == "" {
		return cerrors.ErrorBadRequest("password is required")
	}
	req.Account = account
	return nil
}

func (s *AuthService) normalizeEmail(email string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || utf8.RuneCountInString(email) > 254 {
		return "", cerrors.ErrorBadRequest("email is invalid")
	}
	parsed, err := mail.ParseAddress(email)
	if err != nil || parsed.Address != email || !strings.Contains(email, "@") {
		return "", cerrors.ErrorBadRequest("email is invalid")
	}
	return email, nil
}

func (s *AuthService) normalizePhone(phone string) (string, error) {
	phone = strings.TrimSpace(phone)
	if !s.phoneRe.MatchString(phone) {
		return "", cerrors.ErrorBadRequest("phone is invalid")
	}
	return phone, nil
}

func (s *AuthService) normalizeName(name string) (string, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	length := utf8.RuneCountInString(name)
	if length < 4 || length > 32 || !s.nameRe.MatchString(name) {
		return "", cerrors.ErrorBadRequest("name is invalid")
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
		return nil, cerrors.ErrorBadRequest("nickname is invalid")
	}
	for _, r := range value {
		if !unicode.IsDigit(r) {
			return new(value), nil
		}
	}
	return nil, cerrors.ErrorBadRequest("nickname is invalid")
}

func (s *AuthService) validatePassword(password string) error {
	if len(password) < 8 || len(password) > 64 {
		return cerrors.ErrorBadRequest("password is invalid")
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range password {
		if r < '!' || r > '~' {
			return cerrors.ErrorBadRequest("password is invalid")
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
		return cerrors.ErrorBadRequest("password is invalid")
	}
	return nil
}

func (s *AuthService) normalizeVerification(code string, codeToken string) (string, string, error) {
	code = strings.TrimSpace(code)
	codeToken = strings.TrimSpace(codeToken)
	if !s.codeRe.MatchString(code) {
		return "", "", cerrors.ErrorBadRequest("verification code is invalid")
	}
	if codeToken == "" {
		return "", "", cerrors.ErrorBadRequest("code_token is required")
	}
	return code, codeToken, nil
}

func (s *AuthService) normalizeLoginAccount(account string) (string, error) {
	account = strings.TrimSpace(account)
	if account == "" {
		return "", cerrors.ErrorBadRequest("account is required")
	}
	if strings.Contains(account, "@") {
		return s.normalizeEmail(account)
	}
	if s.phoneRe.MatchString(account) {
		return account, nil
	}
	return s.normalizeName(account)
}
