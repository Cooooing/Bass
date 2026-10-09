package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// CharacterCloseSessionReason 表示角色在线会话关闭原因。
type CharacterCloseSessionReason string

const (
	CharacterCloseSessionReasonOccupied CharacterCloseSessionReason = "occupied" // 被其他连接挤占
	CharacterCloseSessionReasonTimeout  CharacterCloseSessionReason = "timeout"  // 心跳超时
)

var CharacterCloseSessionReasonMap = commonenum.NewMapping[CharacterCloseSessionReason, v1.CharacterCloseSessionReason](map[CharacterCloseSessionReason]commonenum.Entry[CharacterCloseSessionReason, v1.CharacterCloseSessionReason]{
	CharacterCloseSessionReasonOccupied: {Proto: v1.CharacterCloseSessionReason_CHARACTER_CLOSE_SESSION_REASON_OCCUPIED}, CharacterCloseSessionReasonTimeout: {Proto: v1.CharacterCloseSessionReason_CHARACTER_CLOSE_SESSION_REASON_TIMEOUT},
})
