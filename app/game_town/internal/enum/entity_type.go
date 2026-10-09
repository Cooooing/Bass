package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type EntityType string

const (
	EntityTypePlayer   EntityType = "player"   // player
	EntityTypeNpc      EntityType = "npc"      // npc
	EntityTypeLocation EntityType = "location" // location
	EntityTypeFaction  EntityType = "faction"  // faction
)

var EntityTypeMap = commonenum.NewMapping[EntityType, v1.GameTownEntityType](map[EntityType]commonenum.Entry[EntityType, v1.GameTownEntityType]{
	EntityTypePlayer: {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_PLAYER}, EntityTypeNpc: {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_NPC}, EntityTypeLocation: {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_LOCATION}, EntityTypeFaction: {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_FACTION},
})

func (e EntityType) String() string { return string(e) }
