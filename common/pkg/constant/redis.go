package constant

import (
	v1 "common/api/notify/v1"
	"fmt"
)

// Redis key
var (
	RequestNonce            = "Auth:RequestNonce:{%s}"            // 请求防重放
	TokenVerifyCode         = "Auth:TokenVerifyCode:{%s}:{%s}"    // 验证码 Token
	Token                   = "Auth:Token:{%s}"                   // Token
	NotificationTemplateMap = "Notify:NotificationTemplateMap"    // 通知模板
	TwoFactorAuthentication = "Auth:TwoFactorAuthentication:{%s}" // 2FA 验证码，首次启用 2FA 时使用

	SignalNode     = "Signal:Node:{%s}" // 信令服务 ws节点信息 map
	SignalNodeRank = "Signal:NodeRank"  // 信令服务 ws节点评分排名 zset，用于负载均衡
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

func GetKeyNotificationTemplate(notificationType *v1.NotificationType, channel *v1.NotificationChannel) string {
	if notificationType == nil || channel == nil {
		return ""
	}
	return fmt.Sprintf("%s_%s", notificationType.String(), channel.String())
}

func GetKeyTwoFactorAuthentication(name string) string {
	return fmt.Sprintf(TwoFactorAuthentication, name)
}

const (
	SignalNodeData               = "data"                // 基础信息
	SignalNodeCurrentConnections = "current_connections" // 当前连接数
	SignalNodePingMs             = "ping_ms"             // 节点 ping 耗时
	SignalNodePowCostMs          = "pow_cost_ms"         // 节点 PoW 耗时
	SignalNodeLastPingTime       = "last_ping_time"      // 最后一次 ping 时间
)

func GetKeySignalNode(nodeName string) string {
	return fmt.Sprintf(SignalNode, nodeName)
}
