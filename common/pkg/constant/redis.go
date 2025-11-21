package constant

import "fmt"

var Authentication = "Authorization" // token 请求头名称

// Redis key
var (
	TokenEmailCode = "TokenEmailCode::{%s}"
	Token          = "Token::{%s}"
)

func GetKeyTokenEmailCode(email string) string {
	return fmt.Sprintf(TokenEmailCode, email)
}

func GetKeyToken(token string) string {
	return fmt.Sprintf(Token, token)
}
