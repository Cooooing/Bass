package constant

import (
	"fmt"
)

// RedisKey 定义 Redis 键模板。
var (
	RequestNonce            = "Auth:RequestNonce:{%s}"         // 请求防重放。
	TokenVerifyCode         = "Auth:TokenVerifyCode:{%s}:{%s}" // 验证码令牌。
	Token                   = "Auth:Token:{%s}"                // 登录令牌。
	NotificationTemplateMap = "Notify:NotificationTemplateMap" // 通知模板。
	TwoFactorAuth           = "Auth:TwoFactorAuth:{%s}"        // 2FA 验证码，首次启用 2FA 时使用。
	AsynqTaskVersion        = "Asynq:TaskVersion"              // Asynq 任务版本号映射，任务名 -> 版本。

	SignalNode            = "Signal:Node:{%s}"            // 信令服务 ws 节点信息映射。
	SignalNodeRank        = "Signal:NodeRank"             // 信令服务 ws 节点评分有序集合，用于负载均衡。
	SignalTicket          = "Signal:Ticket:{%s}"          // 信令服务 ws 连接一次性认证凭据，凭据 -> 用户 ID。
	SignalSession         = "Signal:Session:{%d}"         // 信令服务 ws 连接会话映射，用户 ID 和会话 ID -> 节点键。
	SignalSessionUser     = "Signal:SessionUser:{%s}"     // 信令服务 ws 连接会话字符串，会话 ID -> 用户 ID。
	SignalNodeKeySessions = "Signal:NodeKeySessions:{%s}" // 信令服务 ws 节点标识集合，节点键 -> 会话 ID。

	ConnectorSession = "Connector:Session" // 连接服务会话集合，会话 ID。
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

func GetKeyNotificationTemplateMap() string {
	return NotificationTemplateMap
}

func GetKeyTwoFactorAuth(name string) string {
	return fmt.Sprintf(TwoFactorAuth, name)
}

const (
	SignalNodeData               = "data"                // 基础信息。
	SignalNodeCurrentConnections = "current_connections" // 当前连接数。
	SignalNodePingMs             = "ping_ms"             // 节点 ping 耗时。
	SignalNodePowCostMs          = "pow_cost_ms"         // 节点 PoW 耗时。
	SignalNodeLastPingTime       = "last_ping_time"      // 最后一次 ping 时间。
)

func GetKeySignalNode(nodeName string) string {
	return fmt.Sprintf(SignalNode, nodeName)
}

func GetKeySignalTicket(ticket string) string {
	return fmt.Sprintf(SignalTicket, ticket)
}

func GetKeySignalSession(userId int64) string {
	return fmt.Sprintf(SignalSession, userId)
}

func GetKeySignalSessionUser(sessionId string) string {
	return fmt.Sprintf(SignalSessionUser, sessionId)
}

func GetKeySignalNodeKeySessions(nodeKey string) string {
	return fmt.Sprintf(SignalNodeKeySessions, nodeKey)
}
