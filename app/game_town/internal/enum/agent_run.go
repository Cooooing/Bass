package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// AgentRunType 是 Agent 执行类型的内部持久化枚举。
type AgentRunType string

const (
	AgentRunTypeWorldGenerate AgentRunType = "world_generate"
	AgentRunTypeNpcTalk       AgentRunType = "npc_talk"
	AgentRunTypeWorldDirector AgentRunType = "world_director"
)

// AgentRunTypeMap 维护内部持久化值与 proto 枚举之间的映射。
var AgentRunTypeMap = commonenum.NewMapping[AgentRunType, v1.GameTownAgentRunType](map[AgentRunType]commonenum.Entry[AgentRunType, v1.GameTownAgentRunType]{
	AgentRunTypeWorldGenerate: {Proto: v1.GameTownAgentRunType_GAME_TOWN_AGENT_RUN_TYPE_WORLD_GENERATE},
	AgentRunTypeNpcTalk:       {Proto: v1.GameTownAgentRunType_GAME_TOWN_AGENT_RUN_TYPE_NPC_TALK},
	AgentRunTypeWorldDirector: {Proto: v1.GameTownAgentRunType_GAME_TOWN_AGENT_RUN_TYPE_WORLD_DIRECTOR},
})

// AgentRunStatus 是 Agent 执行结果状态的内部持久化枚举。
type AgentRunStatus string

const (
	AgentRunStatusSucceeded AgentRunStatus = "succeeded"
	AgentRunStatusFailed    AgentRunStatus = "failed"
)

// AgentRunStatusMap 维护内部持久化值与 proto 枚举之间的映射。
var AgentRunStatusMap = commonenum.NewMapping[AgentRunStatus, v1.GameTownAgentRunStatus](map[AgentRunStatus]commonenum.Entry[AgentRunStatus, v1.GameTownAgentRunStatus]{
	AgentRunStatusSucceeded: {Proto: v1.GameTownAgentRunStatus_GAME_TOWN_AGENT_RUN_STATUS_SUCCEEDED},
	AgentRunStatusFailed:    {Proto: v1.GameTownAgentRunStatus_GAME_TOWN_AGENT_RUN_STATUS_FAILED},
})
