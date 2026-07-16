package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// SessionClientType 是会话客户端类型的内部持久化枚举。
type SessionClientType string

const (
	SessionClientTypeTextClient SessionClientType = "text_client"
)

// SessionClientTypeMap 维护内部持久化值与 proto 枚举之间的映射。
var SessionClientTypeMap = commonenum.NewMapping[SessionClientType, v1.GameTownSessionClientType](map[SessionClientType]commonenum.Entry[SessionClientType, v1.GameTownSessionClientType]{
	SessionClientTypeTextClient: {Proto: v1.GameTownSessionClientType_GAME_TOWN_SESSION_CLIENT_TYPE_TEXT_CLIENT},
})
