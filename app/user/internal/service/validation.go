package service

import (
	cerrors "common/api/gen/common/errors"
	"unicode"
	"unicode/utf8"
)

const (
	maxProfileURLLength          = 2048
	maxProfileIntroductionLength = 512
)

func (s *AuthService) validateRegister(name string, nickname *string, password string) error {
	if !s.validName(name) {
		return cerrors.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if nickname != nil && !s.validNickname(*nickname) {
		return cerrors.ErrorBadRequest("nickname must be 2-32 characters long and contain at least one non-digit character")
	}
	if !s.validPassword(password) {
		return cerrors.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number")
	}
	return nil
}

func (s *AuthService) validPassword(password string) bool {
	length := utf8.RuneCountInString(password)
	return length >= 6 && length <= 64 &&
		s.passRe.MatchString(password) &&
		s.letterRe.MatchString(password) &&
		s.numRe.MatchString(password)
}

func (s *AuthService) validName(name string) bool {
	length := utf8.RuneCountInString(name)
	return length >= 4 && length <= 32 &&
		s.nameRe.MatchString(name)
}

func (s *AuthService) validNickname(nickname string) bool {
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

func (s *AccountService) validateProfileUpdate(avatarURL *string, nickname *string, url *string, introduction *string) error {
	if nickname != nil && *nickname != "" && !s.validNickname(*nickname) {
		return cerrors.ErrorBadRequest("nickname must be 2-32 characters long and contain at least one non-digit character")
	}
	if avatarURL != nil && utf8.RuneCountInString(*avatarURL) > maxProfileURLLength {
		return cerrors.ErrorBadRequest("avatar_url is too long")
	}
	if url != nil && utf8.RuneCountInString(*url) > maxProfileURLLength {
		return cerrors.ErrorBadRequest("url is too long")
	}
	if introduction != nil && utf8.RuneCountInString(*introduction) > maxProfileIntroductionLength {
		return cerrors.ErrorBadRequest("introduction is too long")
	}
	return nil
}

func (s *AccountService) validNickname(nickname string) bool {
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
