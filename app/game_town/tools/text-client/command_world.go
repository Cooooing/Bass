package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"common/pkg/client/rpc"
	"common/proto/gen/common"
	v1 "common/proto/gen/game_town/v1"

	"github.com/samber/lo"
)

func useWorld(ctx context.Context, client *rpc.GameTownClient, value string) commandResult {
	worldID, err := strconv.ParseInt(value, 10, 64)
	if err != nil || worldID <= 0 {
		return commandUsage("/world use <world_id>")
	}
	reply, err := client.World.Get(ctx, &v1.GetGameTownWorld_Request{
		Id: worldID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	return commandResult{
		worldID: worldID,
		lines: []string{fmt.Sprintf(
			"using world %s (%s, id=%d, status=%s)",
			reply.GetRow().GetName(),
			reply.GetRow().GetCode(),
			worldID,
			reply.GetRow().GetStatus(),
		)},
	}
}

func listWorlds(ctx context.Context, client *rpc.GameTownClient) commandResult {
	reply, err := client.World.Page(ctx, &v1.PageGameTownWorlds_Request{
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
			lines: []string{"暂无世界"},
		}
	}
	return commandResult{
		lines: lo.Map(reply.GetRows(), func(row *v1.PageGameTownWorlds_Resp_Row, _ int) string {
			return fmt.Sprintf("world id=%d %s (%s) status=%s", row.GetId(), row.GetName(), row.GetCode(), row.GetStatus())
		}),
	}
}

func createWorld(ctx context.Context, client *rpc.GameTownClient, playerID int64, parts []string) commandResult {
	if playerID == 0 {
		return commandResult{
			err: fmt.Errorf("请先使用 /register <name> 注册玩家，或 /player use <player_id>"),
		}
	}
	configID, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil || configID <= 0 {
		return commandResult{
			err: fmt.Errorf("invalid config id"),
		}
	}
	description := strings.TrimSpace(strings.Join(parts[3:], " "))
	if description == "" {
		return commandUsage("/world create <config_id> <description>")
	}
	reply, err := client.World.Create(ctx, &v1.CreateGameTownWorld_Request{
		CreatorPlayerId: playerID,
		Description:     description,
		AgentConfigId:   configID,
		NpcCount:        4,
		LocationCount:   4,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	return commandResult{
		worldID: reply.GetRow().GetId(),
		lines: []string{
			fmt.Sprintf("world queued: %s event=%d", reply.GetRow().GetCode(), reply.GetEventId()),
			"请等待 world_ready 事件后执行 /world join <world_code> [character_preference]。创建者也需要加入世界后才会生成角色。",
		},
	}
}

func joinWorld(ctx context.Context, client *rpc.GameTownClient, playerID int64, parts []string) commandResult {
	if playerID == 0 {
		return commandResult{
			err: fmt.Errorf("请先使用 /register <name> 注册玩家，或 /player use <player_id>"),
		}
	}
	if len(parts) == 0 || strings.TrimSpace(parts[0]) == "" {
		return commandUsage("/world join <world_code> [character_preference]")
	}
	worldCode := strings.TrimSpace(parts[0])
	if strings.HasPrefix(worldCode, "<") && strings.HasSuffix(worldCode, ">") {
		return commandResult{
			err: fmt.Errorf("请把 %s 替换成 /world list 中显示的真实世界 code，例如 /world join wdkxxxx 我的角色倾向", worldCode),
		}
	}
	preference := strings.TrimSpace(strings.Join(parts[1:], " "))
	reply, err := client.WorldMember.Join(ctx, &v1.JoinGameTownWorld_Request{
		PlayerId:            playerID,
		WorldCode:           worldCode,
		CharacterPreference: &preference,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	return commandResult{
		worldID: reply.GetWorldId(),
		lines: []string{
			"joined world",
			"世界正在为你生成角色，等待 player_character_ready 事件。",
		},
	}
}
