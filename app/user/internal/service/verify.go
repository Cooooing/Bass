package service

import (
	"regexp"
	"unicode/utf8"
)

// VerifyService 参数校验
type VerifyService struct {
	passRe   *regexp.Regexp
	letterRe *regexp.Regexp
	numRe    *regexp.Regexp
	nameRe   *regexp.Regexp
}

// NewVerifyService 预编译正则
func NewVerifyService() *VerifyService {
	return &VerifyService{
		passRe:   regexp.MustCompile(`^[!-~]+$`),
		letterRe: regexp.MustCompile(`[A-Za-z]`),
		numRe:    regexp.MustCompile(`[0-9]`),
		nameRe:   regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
	}
}

// VerifyPassword 校验密码
// 规则：长度 4-64，允许字符集，至少包含一个字母和一个数字
func (s *VerifyService) VerifyPassword(password string) bool {
	length := utf8.RuneCountInString(password)
	return length >= 4 && length <= 64 &&
		s.passRe.MatchString(password) &&
		s.letterRe.MatchString(password) &&
		s.numRe.MatchString(password)
}

// VerifyName 校验用户名
// 规则：长度 1-32，字母数字或单连字符，不能以连字符开头或结尾
func (s *VerifyService) VerifyName(name string) bool {
	length := utf8.RuneCountInString(name)
	return length >= 1 && length <= 32 &&
		s.nameRe.MatchString(name)
}

// VerifyNickname 校验昵称
// 规则：长度 1-32
func (s *VerifyService) VerifyNickname(nickname string) bool {
	length := utf8.RuneCountInString(nickname)
	return length >= 1 && length <= 32
}
