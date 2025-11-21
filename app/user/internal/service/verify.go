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
	nickRe   *regexp.Regexp
	allNumRe *regexp.Regexp
}

// NewVerifyService 预编译正则
func NewVerifyService() *VerifyService {
	return &VerifyService{
		passRe:   regexp.MustCompile(`^[\]A-Za-z0-9@#$%^&*!()_+\-=[{};:'",.<>/?` + "`~|\\\\]+$"),
		letterRe: regexp.MustCompile(`[A-Za-z]`),
		numRe:    regexp.MustCompile(`[0-9]`),
		nameRe:   regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
		nickRe:   regexp.MustCompile(`^[\p{L}\p{N}_\-]+$`),
		allNumRe: regexp.MustCompile(`^\d+$`),
	}
}

// verifyPassword 校验密码
// 规则：长度 6-64，允许字符集，至少包含一个字母和一个数字
func (s *VerifyService) verifyPassword(password string) bool {
	length := utf8.RuneCountInString(password)
	return length >= 6 && length <= 64 &&
		s.passRe.MatchString(password) &&
		s.letterRe.MatchString(password) &&
		s.numRe.MatchString(password)
}

// verifyName 校验用户名
// 规则：长度 4-32，字母数字或单连字符，不能以连字符开头或结尾
func (s *VerifyService) verifyName(name string) bool {
	length := utf8.RuneCountInString(name)
	return length >= 4 && length <= 32 &&
		s.nameRe.MatchString(name)
}

// verifyNickname 校验昵称
// 规则：长度 2-32，允许 Unicode 字符（含 emoji）、下划线、连字符，至少包含一个非数字字符
func (s *VerifyService) verifyNickname(nickname string) bool {
	length := utf8.RuneCountInString(nickname)
	return length >= 2 && length <= 32 &&
		s.nickRe.MatchString(nickname) &&
		!s.allNumRe.MatchString(nickname)
}
