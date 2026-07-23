package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

type AgentJobType string

const (
	AgentJobTypeWorldGenerate           AgentJobType = "world_generate"
	AgentJobTypeNpcTalk                 AgentJobType = "npc_talk"
	AgentJobTypePlayerAct               AgentJobType = "player_act"
	AgentJobTypeWorldTick               AgentJobType = "world_tick"
	AgentJobTypePlayerActionInterpret   AgentJobType = "player_action_interpret"
	AgentJobTypeNpcPlan                 AgentJobType = "npc_plan"
	AgentJobTypeMemoryEmbed             AgentJobType = "memory_embed"
	AgentJobTypePlayerCharacterGenerate AgentJobType = "player_character_generate"
)

var AgentJobTypeMap = commonenum.NewMapping[AgentJobType, v1.GameTownAgentJobType](
	map[AgentJobType]commonenum.Entry[AgentJobType, v1.GameTownAgentJobType]{
		AgentJobTypeWorldGenerate:           {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_WORLD_GENERATE},
		AgentJobTypeNpcTalk:                 {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_NPC_TALK},
		AgentJobTypePlayerAct:               {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_PLAYER_ACT},
		AgentJobTypeWorldTick:               {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_WORLD_TICK},
		AgentJobTypePlayerActionInterpret:   {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_PLAYER_ACTION_INTERPRET},
		AgentJobTypeNpcPlan:                 {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_NPC_PLAN},
		AgentJobTypeMemoryEmbed:             {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_MEMORY_EMBED},
		AgentJobTypePlayerCharacterGenerate: {Proto: v1.GameTownAgentJobType_GAME_TOWN_AGENT_JOB_TYPE_PLAYER_CHARACTER_GENERATE},
	},
)

type AgentJobPriority string

const (
	AgentJobPriorityLow  AgentJobPriority = "low"
	AgentJobPriorityHigh AgentJobPriority = "high"
)

var AgentJobPriorityMap = commonenum.NewMapping[AgentJobPriority, v1.GameTownAgentJobPriority](
	map[AgentJobPriority]commonenum.Entry[AgentJobPriority, v1.GameTownAgentJobPriority]{
		AgentJobPriorityLow:  {Proto: v1.GameTownAgentJobPriority_GAME_TOWN_AGENT_JOB_PRIORITY_LOW},
		AgentJobPriorityHigh: {Proto: v1.GameTownAgentJobPriority_GAME_TOWN_AGENT_JOB_PRIORITY_HIGH},
	},
)

type AgentJobStatus string

const (
	AgentJobStatusQueued    AgentJobStatus = "queued"
	AgentJobStatusRunning   AgentJobStatus = "running"
	AgentJobStatusSucceeded AgentJobStatus = "succeeded"
	AgentJobStatusFailed    AgentJobStatus = "failed"
)

var AgentJobStatusMap = commonenum.NewMapping[AgentJobStatus, v1.GameTownAgentJobStatus](
	map[AgentJobStatus]commonenum.Entry[AgentJobStatus, v1.GameTownAgentJobStatus]{
		AgentJobStatusQueued:    {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_QUEUED},
		AgentJobStatusRunning:   {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_RUNNING},
		AgentJobStatusSucceeded: {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_SUCCEEDED},
		AgentJobStatusFailed:    {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_FAILED},
	},
)
