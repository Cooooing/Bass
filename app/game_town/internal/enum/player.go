package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// PlayerStatus 是玩家账号状态的内部持久化枚举。
type PlayerStatus string

const (
	PlayerStatusActive   PlayerStatus = "active"
	PlayerStatusDisabled PlayerStatus = "disabled"
)

// PlayerStatusMap 维护内部持久化值与 proto 枚举之间的映射。
var PlayerStatusMap = commonenum.NewMapping[PlayerStatus, v1.GameTownPlayerStatus](map[PlayerStatus]commonenum.Entry[PlayerStatus, v1.GameTownPlayerStatus]{
	PlayerStatusActive:   {Proto: v1.GameTownPlayerStatus_GAME_TOWN_PLAYER_STATUS_ACTIVE},
	PlayerStatusDisabled: {Proto: v1.GameTownPlayerStatus_GAME_TOWN_PLAYER_STATUS_DISABLED},
})
