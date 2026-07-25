package main

import (
	v1enum "common/proto/gen/game_town/v1/enum"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"common/pkg/client/rpc"
	v1 "common/proto/gen/game_town/v1"

	"github.com/samber/lo"
)

type suggestedChoice struct {
	label   string
	content string
	targets []*v1.SubmitGameTownAction_Request_EntityRef
}

type commandResult struct {
	lines            []string
	playerID         int64
	worldID          int64
	dialogNpcID      int64
	clearDialog      bool
	clearSuggestions bool
	err              error
}

func executeCommand(parent context.Context, client *rpc.GameTownClient, playerID int64, worldID int64, dialogNpcID int64, suggestions []suggestedChoice, raw string) commandResult {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return commandResult{}
	}
	if raw == "/help" {
		return commandResult{
			lines: helpLines(),
		}
	}
	if raw == "/status" {
		return commandResult{
			lines: []string{fmt.Sprintf("player=%d world=%d dialog_npc=%d suggestions=%d", playerID, worldID, dialogNpcID, len(suggestions))},
		}
	}
	if raw == "/back" {
		return commandResult{
			lines:            []string{"已退出当前对话。"},
			clearDialog:      true,
			clearSuggestions: true,
		}
	}

	ctx, cancel := context.WithTimeout(parent, 15*time.Second)
	defer cancel()

	if !strings.HasPrefix(raw, "/") {
		if result, ok := submitSuggestedChoice(ctx, client, playerID, worldID, dialogNpcID, suggestions, raw); ok {
			result.clearSuggestions = true
			return result
		}
		if isChoiceIndex(raw) {
			return commandResult{
				lines: []string{"当前没有可选回答。你可以直接输入完整行动，或等待 NPC 回复后再输入数字。"},
			}
		}
		if dialogNpcID > 0 {
			return submitAction(ctx, client, playerID, worldID, raw, []*v1.SubmitGameTownAction_Request_EntityRef{npcTarget(dialogNpcID)})
		}
		if playerID > 0 && worldID > 0 {
			return submitAction(ctx, client, playerID, worldID, raw, nil)
		}
		return commandResult{
			err: requireWorldContext(playerID, worldID),
		}
	}

	parts := strings.Fields(raw)
	if len(parts) == 0 {
		return commandResult{}
	}
	switch {
	case parts[0] == "/register" && len(parts) >= 2:
		return registerPlayer(ctx, client, strings.Join(parts[1:], " "))
	case parts[0] == "/player" && len(parts) == 3 && parts[1] == "use":
		return usePlayer(ctx, client, parts[2])
	case raw == "/config list":
		return listAgentConfigs(ctx, client)
	case parts[0] == "/config" && len(parts) >= 6 && parts[1] == "add":
		return createAgentConfig(ctx, client, parts)
	case raw == "/world list":
		return listWorlds(ctx, client)
	case parts[0] == "/world" && len(parts) == 3 && parts[1] == "use":
		return useWorld(ctx, client, parts[2])
	case parts[0] == "/world" && len(parts) >= 4 && parts[1] == "create":
		return createWorld(ctx, client, playerID, parts)
	case parts[0] == "/world" && len(parts) >= 3 && parts[1] == "join":
		return joinWorld(ctx, client, playerID, parts[2:])
	case raw == "/look":
		return lookWorld(ctx, client, playerID, worldID)
	case raw == "/targets" || raw == "/nearby":
		return listTargets(ctx, client, playerID, worldID)
	case raw == "/who":
		return showPlayerCharacter(ctx, client, playerID, worldID)
	case raw == "/npcs":
		return listKnownNpcs(ctx, client, playerID, worldID)
	case raw == "/factions":
		return listKnownFactions(ctx, client, playerID, worldID)
	case parts[0] == "/move" && len(parts) >= 2:
		return movePlayer(ctx, client, playerID, worldID, strings.Join(parts[1:], " "))
	case parts[0] == "/talk" && len(parts) == 2:
		return enterTalk(ctx, client, playerID, worldID, parts[1])
	case parts[0] == "/talk" && len(parts) >= 3:
		return talkNpc(ctx, client, playerID, worldID, parts)
	case parts[0] == "/act" && len(parts) >= 2:
		return actInWorld(ctx, client, playerID, worldID, parts[1:])
	case raw == "/events":
		return listEvents(ctx, client, playerID, worldID)
	case parts[0] == "/register":
		return commandUsage("/register <name>")
	case parts[0] == "/player":
		return commandUsage("/player use <player_id>")
	case parts[0] == "/config":
		return commandUsage("/config list", "/config add <name> <ollama|openai> <base_url> <model> [SECRET_ENV]")
	case parts[0] == "/world":
		return commandUsage("/world list", "/world use <world_id>", "/world create <config_id> <description>", "/world join <world_code> [character_preference]")
	case parts[0] == "/move":
		return commandUsage("/move <location_id|location_code|location_name>")
	case parts[0] == "/talk":
		return commandUsage("/talk <npc_id>", "/talk <npc_id> <content>")
	case parts[0] == "/act":
		return commandUsage("/act <content>")
	default:
		return commandResult{
			lines: []string{"未知命令，输入 /help 查看帮助。自由行动请先 /world join 后直接输入，或使用 /act <content>。"},
		}
	}
}

func submitSuggestedChoice(ctx context.Context, client *rpc.GameTownClient, playerID int64, worldID int64, dialogNpcID int64, suggestions []suggestedChoice, raw string) (commandResult, bool) {
	index, err := strconv.Atoi(raw)
	if err != nil || index <= 0 || index > len(suggestions) {
		return commandResult{}, false
	}
	choice := suggestions[index-1]
	targets := validSubmitTargets(choice.targets)
	if len(targets) == 0 && dialogNpcID > 0 {
		targets = []*v1.SubmitGameTownAction_Request_EntityRef{npcTarget(dialogNpcID)}
	}
	return submitAction(ctx, client, playerID, worldID, choice.content, targets), true
}

func isChoiceIndex(raw string) bool {
	index, err := strconv.Atoi(strings.TrimSpace(raw))
	return err == nil && index > 0
}

func npcTarget(npcID int64) *v1.SubmitGameTownAction_Request_EntityRef {
	return &v1.SubmitGameTownAction_Request_EntityRef{
		Type: v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_NPC,
		Id:   npcID,
	}
}

func validSubmitTargets(targets []*v1.SubmitGameTownAction_Request_EntityRef) []*v1.SubmitGameTownAction_Request_EntityRef {
	result := make([]*v1.SubmitGameTownAction_Request_EntityRef, 0, len(targets))
	for _, target := range targets {
		if target == nil {
			continue
		}
		if target.GetId() <= 0 {
			continue
		}
		if target.GetType() == v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_UNSPECIFIED {
			continue
		}
		result = append(result, target)
	}
	return result
}

func enterTalk(ctx context.Context, client *rpc.GameTownClient, playerID int64, worldID int64, rawNpcID string) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	npcID, err := strconv.ParseInt(rawNpcID, 10, 64)
	if err != nil || npcID <= 0 {
		return commandResult{
			err: fmt.Errorf("invalid npc id"),
		}
	}
	_, err = client.Npc.Get(ctx, &v1.GetGameTownNpc_Request{
		WorldId:  worldID,
		PlayerId: playerID,
		Id:       npcID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	return commandResult{
		lines:       []string{fmt.Sprintf("已进入 NPC %d 对话。直接输入文本发送；输入数字选择建议回答；/back 退出。", npcID)},
		dialogNpcID: npcID,
	}
}

func commandUsage(lines ...string) commandResult {
	return commandResult{
		lines: lo.Map(lines, func(line string, _ int) string {
			return "用法: " + line
		}),
	}
}

func requireWorldContext(playerID int64, worldID int64) error {
	if playerID == 0 {
		return fmt.Errorf("请先 /register <name> 或 /player use <player_id>")
	}
	if worldID == 0 {
		return fmt.Errorf("请先 /world create 后等待 world_ready，再 /world join <world_code>；或使用 /world use <world_id>")
	}
	return nil
}

func helpLines() []string {
	return []string{
		"/register <name>",
		"/player use <player_id>",
		"/config add <name> <ollama|openai> <base_url> <model> [SECRET_ENV]",
		"/config list",
		"/world create <config_id> <description>",
		"/world list",
		"/world use <world_id>",
		"/world join <world_code> [character_preference]",
		"/look",
		"/targets 或 /nearby",
		"/who",
		"/npcs",
		"/factions",
		"/move <location_id|location_code|location_name>",
		"/talk <npc_id>",
		"/talk <npc_id> <content>",
		"/back",
		"/act <content>",
		"直接输入任意内容：作为自由行动发送",
		"输入 1/2/3：选择最近一次建议回答",
		"/events",
		"/status",
		"/quit",
	}
}

func eventName(value v1enum.GameTownEventType) string {
	return strings.ToLower(strings.TrimPrefix(value.String(), "GAME_TOWN_EVENT_TYPE_"))
}
