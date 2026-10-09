package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"    // active
	LocationStatusDamaged   LocationStatus = "damaged"   // damaged
	LocationStatusBlocked   LocationStatus = "blocked"   // blocked
	LocationStatusDestroyed LocationStatus = "destroyed" // destroyed
	LocationStatusRuins     LocationStatus = "ruins"     // ruins
)

var LocationStatusMap = commonenum.NewMapping[LocationStatus, v1.GameTownLocationStatus](map[LocationStatus]commonenum.Entry[LocationStatus, v1.GameTownLocationStatus]{LocationStatusActive: {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_ACTIVE}, LocationStatusDamaged: {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_DAMAGED}, LocationStatusBlocked: {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_BLOCKED}, LocationStatusDestroyed: {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_DESTROYED}, LocationStatusRuins: {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_RUINS}})

func (e LocationStatus) String() string { return string(e) }
