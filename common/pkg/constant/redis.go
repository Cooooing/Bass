package constant

import "fmt"

// RedisKey 定义 Redis 键模板。
var (
	RequestNonce        = "Auth:RequestNonce:{%s}"               // 请求防重放。
	TokenVerifyCode     = "Auth:TokenVerifyCode:{%s}:{%s}"       // 旧验证码令牌，兼容旧 TokenCache。
	Token               = "Auth:Token:{%s}"                      // 旧登录令牌，兼容旧 TokenCache。
	AuthCode            = "Auth:Code:{%s}:{%s}"                  // 通用验证码缓存。
	AuthRegisterDraft   = "Auth:RegisterDraft:{%s}:{%s}"         // 注册草稿缓存。
	AuthRefreshSession  = "Auth:Refresh:{%s}"                    // refresh session 缓存。
	AuthUserSessions    = "Auth:UserSessions:{%s}"               // 用户 session 索引。
	AuthRbacPermissions = "Auth:Rbac:{%s}:{%d}"                  // 用户 RBAC 权限缓存。
	TotpSecret          = "Auth:TotpSecret:{%s}"                 // TOTP 临时密钥。
	AsynqTaskVersion    = "Asynq:TaskVersion"                    // Asynq 任务版本映射。
	OutboxPublisherLock = "Event:OutboxPublisherLock:{%s}"       // outbox 单轮投递锁。
	DeadLetterAlert     = "Event:DeadLetterAlert:{%s}:{%s}:{%s}" // 死信告警去重。
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

func GetKeyAuthCode(codeType string, account string) string {
	return fmt.Sprintf(AuthCode, codeType, account)
}

func GetKeyAuthRegisterDraft(codeType string, account string) string {
	return fmt.Sprintf(AuthRegisterDraft, codeType, account)
}

func GetKeyAuthRefreshSession(sessionID string) string {
	return fmt.Sprintf(AuthRefreshSession, sessionID)
}

func GetKeyAuthUserSessions(userID int64) string {
	return fmt.Sprintf(AuthUserSessions, userID)
}

func GetKeyAuthRbacPermissions(realm string, userID int64) string {
	return fmt.Sprintf(AuthRbacPermissions, realm, userID)
}

func GetPatternAuthUserRbacPermissions(userID int64) string {
	return fmt.Sprintf("Auth:Rbac:*:{%d}", userID)
}

func GetPatternAuthRealmRbacPermissions(realm string) string {
	return fmt.Sprintf("Auth:Rbac:{%s}:*", realm)
}

func GetKeyTotpSecret(name string) string {
	return fmt.Sprintf(TotpSecret, name)
}

func GetKeyOutboxPublisherLock(service string) string {
	return fmt.Sprintf(OutboxPublisherLock, service)
}

func GetKeyDeadLetterAlert(service string, source string, eventID string) string {
	return fmt.Sprintf(DeadLetterAlert, service, source, eventID)
}
