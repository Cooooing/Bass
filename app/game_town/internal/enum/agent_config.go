package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// AgentConfigStatus 是 Agent 配置状态的内部持久化枚举。
type AgentConfigStatus string

const (
	AgentConfigStatusActive   AgentConfigStatus = "active"
	AgentConfigStatusDisabled AgentConfigStatus = "disabled"
)

// AgentConfigStatusMap 维护内部持久化值与 proto 枚举之间的映射。
var AgentConfigStatusMap = commonenum.NewMapping[AgentConfigStatus, v1.GameTownAgentConfigStatus](map[AgentConfigStatus]commonenum.Entry[AgentConfigStatus, v1.GameTownAgentConfigStatus]{
	AgentConfigStatusActive:   {Proto: v1.GameTownAgentConfigStatus_GAME_TOWN_AGENT_CONFIG_STATUS_ACTIVE},
	AgentConfigStatusDisabled: {Proto: v1.GameTownAgentConfigStatus_GAME_TOWN_AGENT_CONFIG_STATUS_DISABLED},
})
