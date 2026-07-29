package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type AgentProvider string

const (
	AgentProviderOllama           AgentProvider = "ollama"
	AgentProviderOpenAICompatible AgentProvider = "openai_compatible"
)

var AgentProviderMap = commonenum.NewMapping[AgentProvider, v1.GameTownAgentProvider](
	map[AgentProvider]commonenum.Entry[AgentProvider, v1.GameTownAgentProvider]{
		AgentProviderOllama: {
			Proto: v1.GameTownAgentProvider_GAME_TOWN_AGENT_PROVIDER_OLLAMA,
		},
		AgentProviderOpenAICompatible: {
			Proto: v1.GameTownAgentProvider_GAME_TOWN_AGENT_PROVIDER_OPENAI_COMPATIBLE,
		},
	},
)

func (e AgentProvider) String() string {
	return string(e)
}
