package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type AgentJobPriority string

const (
	AgentJobPriorityLow  AgentJobPriority = "low"  // low
	AgentJobPriorityHigh AgentJobPriority = "high" // high
)

var AgentJobPriorityMap = commonenum.NewMapping[AgentJobPriority, v1.GameTownAgentJobPriority](map[AgentJobPriority]commonenum.Entry[AgentJobPriority, v1.GameTownAgentJobPriority]{
	AgentJobPriorityLow: {Proto: v1.GameTownAgentJobPriority_GAME_TOWN_AGENT_JOB_PRIORITY_LOW}, AgentJobPriorityHigh: {Proto: v1.GameTownAgentJobPriority_GAME_TOWN_AGENT_JOB_PRIORITY_HIGH},
})

func (e AgentJobPriority) String() string { return string(e) }
