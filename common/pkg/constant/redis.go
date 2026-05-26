package constant

import (
	"fmt"
)

// RedisKey 定义 Redis 键模板。
var (
	RequestNonce     = "Auth:RequestNonce:{%s}"         // 请求防重放。
	TokenVerifyCode  = "Auth:TokenVerifyCode:{%s}:{%s}" // 验证码令牌。
	Token            = "Auth:Token:{%s}"                // 登录令牌。
	TwoFactorAuth    = "Auth:TwoFactorAuth:{%s}"        // 2FA 验证码，首次启用 2FA 时使用。
	AsynqTaskVersion = "Asynq:TaskVersion"              // Asynq 任务版本号映射，任务名 -> 版本。
)

func GetKeyRequestNonce(nonce string) string {
	return fmt.Sprintf(RequestNonce, nonce)
}

type VerifyCodeType string

func (v VerifyCodeType) String() string {
	return string(v)
}

const (
	VerifyCodeTypeRegisterEmail VerifyCodeType = "RegisterEmail"
	VerifyCodeTypeRegisterPhone VerifyCodeType = "RegisterPhone"
)

func GetKeyTokenVerityCode(verifyCodeType VerifyCodeType, account string) string {
	return fmt.Sprintf(TokenVerifyCode, verifyCodeType, account)
}

func GetKeyToken(token string) string {
	return fmt.Sprintf(Token, token)
}

func GetKeyTwoFactorAuth(name string) string {
	return fmt.Sprintf(TwoFactorAuth, name)
}
