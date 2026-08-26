package enum

// CharacterCloseSessionReason 表示角色在线会话关闭原因。
type CharacterCloseSessionReason string

const (
	CharacterCloseSessionReasonOccupied CharacterCloseSessionReason = "occupied" // 被其他连接挤占
	CharacterCloseSessionReasonTimeout  CharacterCloseSessionReason = "timeout"  // 心跳超时
)
