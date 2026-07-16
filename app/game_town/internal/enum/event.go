package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// EventType 是世界事件类型的内部持久化枚举。
type EventType string

const (
	EventTypeWorldCreated      EventType = "world_created"
	EventTypePlayerJoined      EventType = "player_joined"
	EventTypePlayerMoved       EventType = "player_moved"
	EventTypePlayerTalkedToNpc EventType = "player_talked_to_npc"
	EventTypePlayerAction      EventType = "player_action"
	EventTypeNpcDialogue       EventType = "npc_dialogue"
	EventTypeWorldTick         EventType = "world_tick"
)

// EventTypeMap 维护内部持久化值与 proto 枚举之间的映射。
var EventTypeMap = commonenum.NewMapping[EventType, v1.GameTownEventType](map[EventType]commonenum.Entry[EventType, v1.GameTownEventType]{
	EventTypeWorldCreated:      {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_CREATED},
	EventTypePlayerJoined:      {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_JOINED},
	EventTypePlayerMoved:       {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_MOVED},
	EventTypePlayerTalkedToNpc: {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_TALKED_TO_NPC},
	EventTypePlayerAction:      {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_ACTION},
	EventTypeNpcDialogue:       {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_DIALOGUE},
	EventTypeWorldTick:         {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_TICK},
})
