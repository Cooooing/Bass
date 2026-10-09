package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type AgentJobType string

const (
	AgentJobTypeWorldGenerate           AgentJobType = "world_generate"            // world_generate
	AgentJobTypeNpcTalk                 AgentJobType = "npc_talk"                  // npc_talk
	AgentJobTypePlayerAct               AgentJobType = "player_act"                // player_act
	AgentJobTypeWorldTick               AgentJobType = "world_tick"                // world_tick
	AgentJobTypePlayerActionInterpret   AgentJobType = "player_action_interpret"   // player_action_interpret
	AgentJobTypeNpcPlan                 AgentJobType = "npc_plan"                  // npc_plan
	AgentJobTypeMemoryEmbed             AgentJobType = "memory_embed"              // memory_embed
	AgentJobTypePlayerCharacterGenerate AgentJobType = "player_character_generate" // player_character_generate
)

var AgentJobTypeMap = commonenum.NewMapping[AgentJobType, v1.GameTownAgentJobType](map[AgentJobType]commonenum.Entry[AgentJobType, v1.GameTownAgentJobType]{
	AgentJobTypeWorldGenerate: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_WORLD_GENERATE}, AgentJobTypeNpcTalk: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_NPC_TALK}, AgentJobTypePlayerAct: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_PLAYER_ACT}, AgentJobTypeWorldTick: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_WORLD_TICK}, AgentJobTypePlayerActionInterpret: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_PLAYER_ACTION_INTERPRET}, AgentJobTypeNpcPlan: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_NPC_PLAN}, AgentJobTypeMemoryEmbed: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_MEMORY_EMBED}, AgentJobTypePlayerCharacterGenerate: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_PLAYER_CHARACTER_GENERATE},
})

func (e AgentJobType) String() string { return string(e) }
