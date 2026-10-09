package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type AgentJobStatus string

const (
	AgentJobStatusQueued    AgentJobStatus = "queued"    // queued
	AgentJobStatusRunning   AgentJobStatus = "running"   // running
	AgentJobStatusSucceeded AgentJobStatus = "succeeded" // succeeded
	AgentJobStatusFailed    AgentJobStatus = "failed"    // failed
)

var AgentJobStatusMap = commonenum.NewMapping[AgentJobStatus, v1.GameTownAgentJobStatus](map[AgentJobStatus]commonenum.Entry[AgentJobStatus, v1.GameTownAgentJobStatus]{
	AgentJobStatusQueued: {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_QUEUED}, AgentJobStatusRunning: {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_RUNNING}, AgentJobStatusSucceeded: {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_SUCCEEDED}, AgentJobStatusFailed: {Proto: v1.GameTownAgentJobStatus_GAME_TOWN_AGENT_JOB_STATUS_FAILED},
})

func (e AgentJobStatus) String() string { return string(e) }
