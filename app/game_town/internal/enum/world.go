package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// WorldScale 是世界规模的内部持久化枚举。
type WorldScale string

const (
	WorldScaleSmall  WorldScale = "small"
	WorldScaleMedium WorldScale = "medium"
	WorldScaleLarge  WorldScale = "large"
)

// WorldScaleMap 维护内部持久化值与 proto 枚举之间的映射。
var WorldScaleMap = commonenum.NewMapping[WorldScale, v1.GameTownWorldScale](map[WorldScale]commonenum.Entry[WorldScale, v1.GameTownWorldScale]{
	WorldScaleSmall:  {Proto: v1.GameTownWorldScale_GAME_TOWN_WORLD_SCALE_SMALL},
	WorldScaleMedium: {Proto: v1.GameTownWorldScale_GAME_TOWN_WORLD_SCALE_MEDIUM},
	WorldScaleLarge:  {Proto: v1.GameTownWorldScale_GAME_TOWN_WORLD_SCALE_LARGE},
})

// WorldStatus 是世界生命周期状态的内部持久化枚举。
type WorldStatus string

const (
	WorldStatusGenerating WorldStatus = "generating"
	WorldStatusActive     WorldStatus = "active"
	WorldStatusFailed     WorldStatus = "failed"
	WorldStatusArchived   WorldStatus = "archived"
)

// WorldStatusMap 维护内部持久化值与 proto 枚举之间的映射。
var WorldStatusMap = commonenum.NewMapping[WorldStatus, v1.GameTownWorldStatus](map[WorldStatus]commonenum.Entry[WorldStatus, v1.GameTownWorldStatus]{
	WorldStatusGenerating: {Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_GENERATING},
	WorldStatusActive:     {Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_ACTIVE},
	WorldStatusFailed:     {Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_FAILED},
	WorldStatusArchived:   {Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_ARCHIVED},
})
