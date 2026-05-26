package usecase

import (
	cerrors "common/api/gen/common/errors"
	"regexp"
	"unicode"
	"unicode/utf8"
	"user/internal/biz/repo"
)

const (
	maxProfileURLLength          = 2048
	maxProfileIntroductionLength = 512
)

// AccountValidationUsecase 统一维护账号输入校验规则。
type AccountValidationUsecase struct {
	passRe   *regexp.Regexp
	letterRe *regexp.Regexp
	numRe    *regexp.Regexp
	nameRe   *regexp.Regexp
}

func NewAccountValidationUsecase() *AccountValidationUsecase {
	return &AccountValidationUsecase{
		passRe:   regexp.MustCompile(`^[!-~]+$`),
		letterRe: regexp.MustCompile(`[A-Za-z]`),
		numRe:    regexp.MustCompile(`[0-9]`),
		nameRe:   regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
	}
}

func (s *AccountValidationUsecase) ValidateRegister(name string, nickname *string, password string) error {
	if !s.ValidName(name) {
		return cerrors.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if nickname != nil && !s.ValidNickname(*nickname) {
		return cerrors.ErrorBadRequest("nickname must be 2-32 characters long and contain at least one non-digit character")
	}
	if !s.ValidPassword(password) {
		return cerrors.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number")
	}
	return nil
}

func (s *AccountValidationUsecase) ValidateProfileUpdate(patch *repo.AccountProfilePatch) error {
	if patch.Nickname.Set && patch.Nickname.Value != "" && !s.ValidNickname(patch.Nickname.Value) {
		return cerrors.ErrorBadRequest("nickname must be 2-32 characters long and contain at least one non-digit character")
	}
	if patch.AvatarURL.Set && utf8.RuneCountInString(patch.AvatarURL.Value) > maxProfileURLLength {
		return cerrors.ErrorBadRequest("avatar_url is too long")
	}
	if patch.URL.Set && utf8.RuneCountInString(patch.URL.Value) > maxProfileURLLength {
		return cerrors.ErrorBadRequest("url is too long")
	}
	if patch.Introduction.Set && utf8.RuneCountInString(patch.Introduction.Value) > maxProfileIntroductionLength {
		return cerrors.ErrorBadRequest("introduction is too long")
	}
	return nil
}

func (s *AccountValidationUsecase) ValidPassword(password string) bool {
	length := utf8.RuneCountInString(password)
	return length >= 6 && length <= 64 &&
		s.passRe.MatchString(password) &&
		s.letterRe.MatchString(password) &&
		s.numRe.MatchString(password)
}

func (s *AccountValidationUsecase) ValidName(name string) bool {
	length := utf8.RuneCountInString(name)
	return length >= 4 && length <= 32 &&
		s.nameRe.MatchString(name)
}

func (s *AccountValidationUsecase) ValidNickname(nickname string) bool {
	length := utf8.RuneCountInString(nickname)
	if length < 2 || length > 32 {
		return false
	}
	for _, r := range nickname {
		if !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
