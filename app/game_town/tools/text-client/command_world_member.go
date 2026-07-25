package main

import (
	v1enum "common/proto/gen/game_town/v1/enum"
	"context"
	"fmt"
	"strconv"
	"strings"

	"common/pkg/client/rpc"
	v1 "common/proto/gen/game_town/v1"
)

func lookWorld(ctx context.Context, client *rpc.GameTownClient, playerID, worldID int64) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	member, err := client.WorldMember.Get(ctx, &v1.GetGameTownWorldMember_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	locations, err := client.Location.List(ctx, &v1.ListGameTownLocations_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	npcs, err := client.Npc.List(ctx, &v1.ListGameTownNpcs_Request{
		WorldId:    worldID,
		PlayerId:   playerID,
		LocationId: new(member.GetCurrentLocationId()),
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}

	lines := []string{fmt.Sprintf("current location=%d world_time=%s", member.GetCurrentLocationId(), member.GetWorldTime().AsTime().Format("2006-01-02 15:04"))}
	if member.GetCharacterReady() {
		lines = append(lines, fmt.Sprintf("character=%s goal=%s traits=%s", member.GetCharacterName(), member.GetCharacterGoal(), strings.Join(member.GetCharacterTraits(), ",")))
		if member.GetCharacterBackground() != "" {
			lines = append(lines, "background: "+member.GetCharacterBackground())
		}
	} else {
		lines = append(lines, "character=生成中，等待 player_character_ready 事件")
	}
	lines = append(lines, "Known locations:")
	for _, location := range locations.GetRows() {
		marker := "-"
		if location.GetCurrent() {
			marker = "*"
		}
		lines = append(lines, fmt.Sprintf("%s %s [%s] id=%d accessible=%t", marker, location.GetName(), location.GetCode(), location.GetId(), location.GetAccessible()))
	}
	lines = append(lines, "NPCs at current location:")
	if len(npcs.GetRows()) == 0 {
		lines = append(lines, "- 当前地点没有你已知的 NPC")
	}
	for _, npc := range npcs.GetRows() {
		lines = append(lines, fmt.Sprintf("- npc %d %s (%s, %s)", npc.GetId(), npc.GetName(), npc.GetRole(), npc.GetSpecies()))
	}
	lines = append(lines, "可用命令：/talk <npc_id>、/move <地点>、/targets、/who、/npcs、/factions")
	return commandResult{
		lines: lines,
	}
}

func showPlayerCharacter(ctx context.Context, client *rpc.GameTownClient, playerID int64, worldID int64) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	member, err := client.WorldMember.Get(ctx, &v1.GetGameTownWorldMember_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	lines := []string{fmt.Sprintf("player=%d world=%d location=%d", playerID, worldID, member.GetCurrentLocationId())}
	if !member.GetCharacterReady() {
		lines = append(lines, "character=生成中，等待 player_character_ready 事件")
		return commandResult{
			lines: lines,
		}
	}
	lines = append(lines, fmt.Sprintf("character=%s", member.GetCharacterName()))
	if member.GetCharacterGoal() != "" {
		lines = append(lines, "goal: "+member.GetCharacterGoal())
	}
	if len(member.GetCharacterTraits()) > 0 {
		lines = append(lines, "traits: "+strings.Join(member.GetCharacterTraits(), ","))
	}
	if member.GetCharacterBackground() != "" {
		lines = append(lines, "background: "+member.GetCharacterBackground())
	}
	return commandResult{
		lines: lines,
	}
}

func listTargets(ctx context.Context, client *rpc.GameTownClient, playerID int64, worldID int64) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	member, err := client.WorldMember.Get(ctx, &v1.GetGameTownWorldMember_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	locations, err := client.Location.List(ctx, &v1.ListGameTownLocations_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	npcs, err := client.Npc.List(ctx, &v1.ListGameTownNpcs_Request{
		WorldId:    worldID,
		PlayerId:   playerID,
		LocationId: new(member.GetCurrentLocationId()),
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	lines := []string{"可交互目标："}
	lines = append(lines, "NPC:")
	if len(npcs.GetRows()) == 0 {
		lines = append(lines, "- 当前地点没有可见 NPC")
	}
	for _, npc := range npcs.GetRows() {
		lines = append(lines, fmt.Sprintf("- /talk %d  # %s（%s，%s）", npc.GetId(), npc.GetName(), npc.GetRole(), npc.GetSpecies()))
	}
	lines = append(lines, "地点:")
	for _, location := range locations.GetRows() {
		marker := "-"
		if location.GetCurrent() {
			marker = "*"
		}
		lines = append(lines, fmt.Sprintf("%s /move %d  # %s [%s] accessible=%t", marker, location.GetId(), location.GetName(), location.GetCode(), location.GetAccessible()))
	}
	return commandResult{
		lines: lines,
	}
}

func listKnownNpcs(ctx context.Context, client *rpc.GameTownClient, playerID int64, worldID int64) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	reply, err := client.Npc.List(ctx, &v1.ListGameTownNpcs_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	if len(reply.GetRows()) == 0 {
		return commandResult{
			lines: []string{"暂无已知 NPC"},
		}
	}
	lines := []string{"已知 NPC:"}
	for _, npc := range reply.GetRows() {
		location := "未知地点"
		if npc.LastKnownLocationId != nil {
			location = fmt.Sprintf("location=%d", npc.GetLastKnownLocationId())
		}
		lines = append(lines, fmt.Sprintf("- npc %d %s (%s, %s) status=%s %s tags=%s", npc.GetId(), npc.GetName(), npc.GetRole(), npc.GetSpecies(), npc.GetLifeStatus().String(), location, strings.Join(npc.GetStateTags(), ",")))
	}
	return commandResult{
		lines: lines,
	}
}

func listKnownFactions(ctx context.Context, client *rpc.GameTownClient, playerID int64, worldID int64) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	reply, err := client.Faction.List(ctx, &v1.ListGameTownFactions_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	if len(reply.GetRows()) == 0 {
		return commandResult{
			lines: []string{"暂无已知阵营"},
		}
	}
	lines := []string{"已知阵营:"}
	for _, faction := range reply.GetRows() {
		lines = append(lines, fmt.Sprintf("- faction %d %s [%s] status=%s attitude=%s reputation=%s", faction.GetId(), faction.GetName(), faction.GetCode(), faction.GetStatus().String(), faction.GetAttitude().String(), strings.Join(faction.GetReputationTags(), ",")))
	}
	return commandResult{
		lines: lines,
	}
}

func movePlayer(ctx context.Context, client *rpc.GameTownClient, playerID, worldID int64, locationQuery string) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	locationQuery = strings.TrimSpace(locationQuery)
	if locationQuery == "" {
		return commandUsage("/move <location_id|location_code|location_name>")
	}
	location, alternatives, err := resolveLocation(ctx, client, playerID, worldID, locationQuery)
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	if location == nil {
		lines := []string{fmt.Sprintf("未找到地点 %q，可用地点：", locationQuery)}
		lines = append(lines, alternatives...)
		return commandResult{
			lines: lines,
		}
	}
	content := fmt.Sprintf("移动到 %s", location.GetName())
	targets := []*v1.SubmitGameTownAction_Request_EntityRef{{Type: v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_LOCATION, Id: location.GetId()}}
	return submitAction(ctx, client, playerID, worldID, content, targets)
}

func resolveLocation(ctx context.Context, client *rpc.GameTownClient, playerID, worldID int64, query string) (*v1.ListGameTownLocations_Resp_Row, []string, error) {
	reply, err := client.Location.List(ctx, &v1.ListGameTownLocations_Request{
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return nil, nil, err
	}
	query = strings.TrimSpace(query)
	lowerQuery := strings.ToLower(query)
	alternatives := make([]string, 0, len(reply.GetRows()))
	for _, location := range reply.GetRows() {
		alternatives = append(alternatives, fmt.Sprintf("- %s [%s] id=%d accessible=%t", location.GetName(), location.GetCode(), location.GetId(), location.GetAccessible()))
		if strings.EqualFold(location.GetCode(), query) || strings.EqualFold(location.GetName(), query) || strconv.FormatInt(location.GetId(), 10) == query {
			return location, alternatives, nil
		}
	}
	for _, location := range reply.GetRows() {
		if strings.Contains(strings.ToLower(location.GetName()), lowerQuery) || strings.Contains(strings.ToLower(location.GetCode()), lowerQuery) {
			return location, alternatives, nil
		}
	}
	return nil, alternatives, nil
}

func talkNpc(ctx context.Context, client *rpc.GameTownClient, playerID, worldID int64, parts []string) commandResult {
	if len(parts) < 3 {
		return commandResult{
			err: fmt.Errorf("usage: /talk <npc_id> <content>"),
		}
	}
	npcID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return commandResult{
			err: fmt.Errorf("invalid npc id: %w", err),
		}
	}
	result := submitAction(ctx, client, playerID, worldID, strings.Join(parts[2:], " "), []*v1.SubmitGameTownAction_Request_EntityRef{npcTarget(npcID)})
	if result.err == nil {
		result.dialogNpcID = npcID
	}
	return result
}

func actInWorld(ctx context.Context, client *rpc.GameTownClient, playerID, worldID int64, content []string) commandResult {
	return submitAction(ctx, client, playerID, worldID, strings.Join(content, " "), nil)
}

func submitAction(ctx context.Context, client *rpc.GameTownClient, playerID, worldID int64, content string, targets []*v1.SubmitGameTownAction_Request_EntityRef) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	content = strings.TrimSpace(content)
	if content == "" {
		return commandResult{
			err: fmt.Errorf("行动内容不能为空"),
		}
	}
	targets = validSubmitTargets(targets)
	reply, err := client.WorldMember.SubmitAction(ctx, &v1.SubmitGameTownAction_Request{
		WorldId:  worldID,
		PlayerId: playerID,
		Content:  content,
		Targets:  targets,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	return commandResult{
		lines: []string{fmt.Sprintf("action accepted event=%d", reply.GetEventId())},
	}
}
