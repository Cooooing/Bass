package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// NpcState 是 NPC 当前状态的内部持久化枚举。
type NpcState string

const (
	NpcStateIdle     NpcState = "idle"
	NpcStateTalking  NpcState = "talking"
	NpcStateThinking NpcState = "thinking"
	NpcStateBusy     NpcState = "busy"
)

// NpcStateMap 维护内部持久化值与 proto 枚举之间的映射。
var NpcStateMap = commonenum.NewMapping[NpcState, v1.GameTownNpcState](map[NpcState]commonenum.Entry[NpcState, v1.GameTownNpcState]{
	NpcStateIdle:     {Proto: v1.GameTownNpcState_GAME_TOWN_NPC_STATE_IDLE},
	NpcStateTalking:  {Proto: v1.GameTownNpcState_GAME_TOWN_NPC_STATE_TALKING},
	NpcStateThinking: {Proto: v1.GameTownNpcState_GAME_TOWN_NPC_STATE_THINKING},
	NpcStateBusy:     {Proto: v1.GameTownNpcState_GAME_TOWN_NPC_STATE_BUSY},
})
