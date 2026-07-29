package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type WorldStatus string

const (
	WorldStatusGenerating WorldStatus = "generating"
	WorldStatusActive     WorldStatus = "active"
	WorldStatusFailed     WorldStatus = "failed"
	WorldStatusArchived   WorldStatus = "archived"
)

var WorldStatusMap = commonenum.NewMapping[WorldStatus, v1.GameTownWorldStatus](
	map[WorldStatus]commonenum.Entry[WorldStatus, v1.GameTownWorldStatus]{
		WorldStatusGenerating: {
			Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_GENERATING,
		},
		WorldStatusActive: {
			Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_ACTIVE,
		},
		WorldStatusFailed: {
			Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_FAILED,
		},
		WorldStatusArchived: {
			Proto: v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_ARCHIVED,
		},
	},
)

func (e WorldStatus) String() string {
	return string(e)
}
