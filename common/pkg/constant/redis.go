package constant

import "fmt"

var Authentication = "Authorization" // token 请求头名称
var CtxToken = "Token"               // 上下文token key
var CtxUserInfo = "CtxUserInfo"      // 上下文用户信息 key

// Redis key
var (
	TokenEmailCode = "TokenEmailCode::{%s}"
	Token          = "Token::{%s}"
)

func GetKeyTokenEmailCode(token string) string {
	return fmt.Sprintf(TokenEmailCode, token)
}

func GetKeyToken(token string) string {
	return fmt.Sprintf(Token, token)
}
