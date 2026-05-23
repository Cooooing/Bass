package service

import (
	"regexp"
	"unicode/utf8"
)

// VerifyService 在请求进入认证 usecase 前校验账号输入。
type VerifyService struct {
	passRe   *regexp.Regexp
	letterRe *regexp.Regexp
	numRe    *regexp.Regexp
	nameRe   *regexp.Regexp
}

// NewVerifyService 预编译服务使用的校验表达式。
func NewVerifyService() *VerifyService {
	return &VerifyService{
		passRe:   regexp.MustCompile(`^[!-~]+$`),
		letterRe: regexp.MustCompile(`[A-Za-z]`),
		numRe:    regexp.MustCompile(`[0-9]`),
		nameRe:   regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
	}
}

// VerifyPassword 校验密码是否为可打印 ASCII，且至少包含一个字母和一个数字。
func (s *VerifyService) VerifyPassword(password string) bool {
	length := utf8.RuneCountInString(password)
	return length >= 4 && length <= 64 &&
		s.passRe.MatchString(password) &&
		s.letterRe.MatchString(password) &&
		s.numRe.MatchString(password)
}

// VerifyName 校验登录和注册使用的账号名格式。
func (s *VerifyService) VerifyName(name string) bool {
	length := utf8.RuneCountInString(name)
	return length >= 1 && length <= 32 &&
		s.nameRe.MatchString(name)
}

// VerifyNickname 校验展示昵称长度。
func (s *VerifyService) VerifyNickname(nickname string) bool {
	length := utf8.RuneCountInString(nickname)
	return length >= 1 && length <= 32
}
