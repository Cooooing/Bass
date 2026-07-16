package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// MemoryType 是 NPC 记忆类型的内部持久化枚举。
type MemoryType string

const (
	MemoryTypeShortTerm MemoryType = "short_term"
	MemoryTypeLongTerm  MemoryType = "long_term"
)

// MemoryTypeMap 维护内部持久化值与 proto 枚举之间的映射。
var MemoryTypeMap = commonenum.NewMapping[MemoryType, v1.GameTownMemoryType](map[MemoryType]commonenum.Entry[MemoryType, v1.GameTownMemoryType]{
	MemoryTypeShortTerm: {Proto: v1.GameTownMemoryType_GAME_TOWN_MEMORY_TYPE_SHORT_TERM},
	MemoryTypeLongTerm:  {Proto: v1.GameTownMemoryType_GAME_TOWN_MEMORY_TYPE_LONG_TERM},
})
