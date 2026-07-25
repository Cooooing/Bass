package main

import (
	v1enum "common/proto/gen/game_town/v1/enum"
	"context"
	"fmt"
	"strings"

	"common/pkg/client/rpc"
	"common/proto/gen/common"
	v1 "common/proto/gen/game_town/v1"

	"github.com/samber/lo"
)

func listAgentConfigs(ctx context.Context, client *rpc.GameTownClient) commandResult {
	reply, err := client.AgentConfig.List(ctx, &v1.ListGameTownAgentConfigs_Request{
		Page: &common.PageReq{
			Page: 1,
			Size: 100,
		},
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	if len(reply.GetRows()) == 0 {
		return commandResult{
			lines: []string{"暂无 Agent 配置"},
		}
	}
	return commandResult{
		lines: lo.Map(reply.GetRows(), func(row *v1.ListGameTownAgentConfigs_Resp_Row, _ int) string {
			return fmt.Sprintf("config %d %s provider=%s model=%s base_url=%s", row.GetId(), row.GetName(), row.GetProvider(), row.GetModel(), row.GetBaseUrl())
		}),
	}
}

func createAgentConfig(ctx context.Context, client *rpc.GameTownClient, parts []string) commandResult {
	var provider v1enum.GameTownAgentProvider
	switch strings.ToLower(parts[3]) {
	case "ollama":
		provider = v1enum.GameTownAgentProvider_GAME_TOWN_AGENT_PROVIDER_OLLAMA
	case "openai":
		provider = v1enum.GameTownAgentProvider_GAME_TOWN_AGENT_PROVIDER_OPENAI_COMPATIBLE
	default:
		return commandResult{
			err: fmt.Errorf("provider 只支持 ollama 或 openai"),
		}
	}
	secretEnv := ""
	if len(parts) > 6 {
		secretEnv = parts[6]
	}
	reply, err := client.AgentConfig.Create(ctx, &v1.CreateGameTownAgentConfig_Request{
		Name:           parts[2],
		Provider:       provider,
		BaseUrl:        parts[4],
		Model:          parts[5],
		SecretEnv:      secretEnv,
		TimeoutSeconds: 60,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "agent_configs_name") {
			return commandResult{
				lines: []string{"配置名称已存在，没有重复创建。请使用 /config list 查看已有 config_id，或换一个名称。"},
			}
		}
		return commandResult{
			err: err,
		}
	}
	return commandResult{
		lines: []string{fmt.Sprintf("config created: %d", reply.GetRow().GetId())},
	}
}
