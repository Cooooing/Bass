package usecase

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"common/pkg/constant"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

func (r *WorldAgentRunner) callAgent(
	ctx context.Context,
	result *agentResult,
) error {
	var err error
	switch result.job.Type {
	case enum.AgentJobTypeWorldGenerate:
		result.draft, err = r.agentClient.GenerateWorld(ctx, &repo.GenerateWorldReq{
			Config:        result.config,
			World:         result.world,
			NpcCount:      uint32Value(result.source.Payload, "npc_count"),
			LocationCount: uint32Value(result.source.Payload, "location_count"),
		})
		if err != nil {
			r.log.WarnContext(ctx, "world generation agent failed; use fallback draft", "world_id", result.world.ID, constant.LogFieldErr, err)
			result.draft = fallbackWorldDraft(result)
			err = nil
		}
	case enum.AgentJobTypePlayerCharacterGenerate:
		err = r.loadActionContext(ctx, result)
		if err != nil {
			return err
		}
		recentEvents, err := r.playerVisibleEvents(ctx, result.world.ID, result.player.ID)
		if err != nil {
			return err
		}
		preference, _ := result.source.Payload["character_preference"].(string)
		result.character, err = r.agentClient.GenerateCharacter(ctx, &repo.GenerateCharacterReq{
			Config:       result.config,
			World:        result.world,
			State:        result.state,
			Player:       result.player,
			Member:       result.member,
			Location:     result.location,
			RecentEvents: recentEvents,
			Preference:   preference,
		})
	case enum.AgentJobTypeNpcTalk:
		err = r.loadActionContext(ctx, result)
		if err != nil {
			return err
		}
		result.reply, err = r.agentClient.Talk(ctx, &repo.TalkReq{
			Config:       result.config,
			World:        result.world,
			State:        result.state,
			Player:       result.player,
			Member:       result.member,
			Location:     result.location,
			Npc:          result.npc,
			RecentEvents: resultEvents(result),
			Memories:     result.memories,
			Content:      result.source.Content,
		})
		if err != nil {
			r.log.WarnContext(ctx, "npc talk agent failed; use fallback reply", "world_id", result.world.ID, constant.LogFieldErr, err)
			result.reply = fallbackNpcReply(result, err)
			err = nil
		}
	case enum.AgentJobTypePlayerAct, enum.AgentJobTypePlayerActionInterpret:
		err = r.loadActionContext(ctx, result)
		if err != nil {
			return err
		}
		if result.npc != nil {
			result.reply, err = r.agentClient.Talk(ctx, &repo.TalkReq{
				Config:       result.config,
				World:        result.world,
				State:        result.state,
				Player:       result.player,
				Member:       result.member,
				Location:     result.location,
				Npc:          result.npc,
				RecentEvents: resultEvents(result),
				Memories:     result.memories,
				Content:      result.source.Content,
			})
			if err != nil {
				r.log.WarnContext(ctx, "target npc action agent failed; use fallback reply", "world_id", result.world.ID, constant.LogFieldErr, err)
				result.reply = fallbackNpcReply(result, err)
				err = nil
			}
			break
		}
		result.resolution, err = r.agentClient.Act(ctx, &repo.ActReq{
			Config:       result.config,
			World:        result.world,
			State:        result.state,
			Player:       result.player,
			Member:       result.member,
			Location:     result.location,
			RecentEvents: resultEvents(result),
			Content:      result.source.Content,
			Targets:      actionTargets(result.source.Payload),
		})
		if err != nil {
			r.log.WarnContext(ctx, "player action agent failed; use fallback resolution", "world_id", result.world.ID, constant.LogFieldErr, err)
			result.resolution = fallbackPlayerActionResolution(result, err)
			err = nil
		}
	case enum.AgentJobTypeNpcPlan:
		if result.source.NpcID == nil {
			return nil
		}
		result.npc, err = r.npcRepo.Get(ctx, &repo.NpcQuery{
			ID: result.source.NpcID,
		})
		if err != nil {
			return err
		}
		result.location, err = r.locationRepo.Get(ctx, &repo.LocationQuery{
			ID: new(result.npc.CurrentLocationID),
		})
		if err != nil {
			return err
		}
		result.locations, err = r.locationRepo.List(ctx, &repo.LocationQuery{
			WorldID: new(result.world.ID),
		})
		if err != nil {
			return err
		}
		resultEvents, err := r.npcVisibleEvents(ctx, result.world.ID, result.npc.ID)
		if err != nil {
			return err
		}
		result.memories, err = r.npcMemories(ctx, result.npc, result.npc.Goal+" "+result.npc.ContextSummary)
		if err != nil {
			return err
		}
		result.plan, err = r.agentClient.PlanNpc(ctx, &repo.PlanNpcReq{
			Config:       result.config,
			World:        result.world,
			State:        result.state,
			Location:     result.location,
			Npc:          result.npc,
			RecentEvents: resultEvents,
			Memories:     result.memories,
		})
		if err != nil {
			r.log.WarnContext(ctx, "npc plan agent failed; use fallback plan", "world_id", result.world.ID, "npc_id", result.npc.ID, constant.LogFieldErr, err)
			result.plan = fallbackNpcPlan(result, err)
			err = nil
		}
	case enum.AgentJobTypeWorldTick:
		recentEvents, loadErr := r.recentEvents(ctx, result.world.ID)
		if loadErr != nil {
			return loadErr
		}
		result.npcs, loadErr = r.npcRepo.List(ctx, &repo.NpcQuery{
			WorldID: new(result.world.ID),
		})
		if loadErr != nil {
			return loadErr
		}
		result.locations, loadErr = r.locationRepo.List(ctx, &repo.LocationQuery{
			WorldID: new(result.world.ID),
		})
		if loadErr != nil {
			return loadErr
		}
		result.factions, loadErr = r.factionRepo.List(ctx, &repo.FactionQuery{
			WorldID: new(result.world.ID),
		})
		if loadErr != nil {
			return loadErr
		}
		if result.source.Payload == nil {
			result.source.Payload = map[string]any{}
		}
		result.source.Payload["scoped_events"] = recentEvents
		result.resolution, err = r.agentClient.Tick(ctx, &repo.TickReq{
			Config:       result.config,
			World:        result.world,
			State:        result.state,
			RecentEvents: recentEvents,
			Npcs:         result.npcs,
			Locations:    result.locations,
			Factions:     result.factions,
		})
		if err != nil {
			r.log.WarnContext(ctx, "world tick agent failed; use fallback resolution", "world_id", result.world.ID, constant.LogFieldErr, err)
			result.resolution = fallbackWorldTickResolution(result, err)
			err = nil
		}
	case enum.AgentJobTypeMemoryEmbed:
		if result.job.NpcID == nil {
			return fmt.Errorf("memory embed job missing npc id")
		}
		status := enum.EmbeddingStatusPending
		rows, loadErr := r.npcMemoryRepo.List(ctx, &repo.NpcMemoryQuery{
			WorldID:       new(result.job.WorldID),
			NpcID:         result.job.NpcID,
			SourceEventID: new(result.job.SourceEventID),
			Status:        new(status),
			RecentLimit:   1,
		})
		if loadErr != nil {
			return loadErr
		}
		if len(rows) == 0 {
			return nil
		}
		vectors, embedErr := r.embeddingClient.Embed(ctx, []string{rows[0].Content})
		if embedErr != nil {
			return embedErr
		}
		if len(vectors) != 1 {
			return fmt.Errorf("embedding count mismatch: got %d", len(vectors))
		}
		result.memory = rows[0]
		result.embedding = vectors[0]
	}
	return err
}

func (r *WorldAgentRunner) loadActionContext(
	ctx context.Context,
	result *agentResult,
) error {
	var err error
	result.player, err = r.playerRepo.Get(ctx, &repo.PlayerQuery{
		ID: result.source.ActorPlayerID,
	})
	if err != nil {
		return err
	}
	result.member, err = r.worldMemberRepo.Get(ctx, &repo.WorldMemberQuery{
		WorldID:  new(result.world.ID),
		PlayerID: result.source.ActorPlayerID,
	})
	if err != nil {
		return err
	}
	result.location, err = r.locationRepo.Get(ctx, &repo.LocationQuery{
		ID: new(result.member.CurrentLocationID),
	})
	if err != nil {
		return err
	}
	if result.source.NpcID != nil {
		result.npc, err = r.npcRepo.Get(ctx, &repo.NpcQuery{
			ID: result.source.NpcID,
		})
		if err != nil {
			return err
		}
		resultEvents, err := r.npcVisibleEvents(ctx, result.world.ID, result.npc.ID)
		if err != nil {
			return err
		}
		if result.source.Payload == nil {
			result.source.Payload = map[string]any{}
		}
		result.source.Payload["scoped_events"] = resultEvents
		result.memories, err = r.npcMemories(ctx, result.npc, result.source.Content)
		return err
	}
	resultEvents, err := r.playerVisibleEvents(ctx, result.world.ID, result.player.ID)
	if err != nil {
		return err
	}
	if result.source.Payload == nil {
		result.source.Payload = map[string]any{}
	}
	result.source.Payload["scoped_events"] = resultEvents
	return nil
}

func resultEvents(
	result *agentResult,
) []*model.Event {
	if result == nil || result.source == nil || result.source.Payload == nil {
		return nil
	}
	events, _ := result.source.Payload["scoped_events"].([]*model.Event)
	return events
}

func (r *WorldAgentRunner) npcVisibleEvents(
	ctx context.Context,
	worldID, npcID int64,
) ([]*model.Event, error) {
	return r.visibleEvents(ctx, &repo.ObservationQuery{
		WorldID: new(worldID),
		NpcID:   new(npcID),
	})
}

func (r *WorldAgentRunner) playerVisibleEvents(
	ctx context.Context,
	worldID, playerID int64,
) ([]*model.Event, error) {
	return r.visibleEvents(ctx, &repo.ObservationQuery{
		WorldID:  new(worldID),
		PlayerID: new(playerID),
	})
}

func (r *WorldAgentRunner) visibleEvents(
	ctx context.Context,
	query *repo.ObservationQuery,
) ([]*model.Event, error) {
	observations, err := r.observationRepo.List(ctx, query)
	if err != nil {
		return nil, err
	}
	limit := int(r.conf.GetGameTown().GetAgent().GetRecentEventLimit())
	if limit <= 0 {
		limit = 20
	}
	if len(observations) > limit {
		observations = observations[len(observations)-limit:]
	}
	eventIDs := make([]int64, 0, len(observations))
	for _, observation := range observations {
		eventIDs = append(eventIDs, observation.EventID)
	}
	if len(eventIDs) == 0 {
		return nil, nil
	}
	return r.eventRepo.List(ctx, &repo.EventQuery{
		IDs: eventIDs,
	})
}

func (r *WorldAgentRunner) npcMemories(
	ctx context.Context,
	npc *model.Npc,
	prompt string,
) ([]*model.NpcMemory, error) {
	if npc == nil {
		return nil, nil
	}
	if r.embeddingEnabled() && strings.TrimSpace(prompt) != "" {
		vectors, err := r.embeddingClient.Embed(ctx, []string{prompt})
		if err == nil && len(vectors) == 1 {
			memories, searchErr := r.npcMemoryRepo.Search(ctx, &repo.NpcMemorySearchReq{
				WorldID:        npc.WorldID,
				NpcID:          npc.ID,
				Vector:         vectors[0],
				CandidateLimit: int(r.conf.GetGameTown().GetMemory().GetCandidateLimit()),
				ResultLimit:    int(r.conf.GetGameTown().GetMemory().GetResultLimit()),
				Now:            time.Now(),
			})
			if searchErr == nil {
				return memories, nil
			}
			r.log.WarnContext(ctx, "search npc memory failed; continue without vector memories", "world_id", npc.WorldID, "npc_id", npc.ID, constant.LogFieldErr, searchErr)
			return nil, nil
		} else if err != nil {
			r.log.WarnContext(ctx, "embed npc memory prompt failed; fallback to recent memories", "world_id", npc.WorldID, "npc_id", npc.ID, constant.LogFieldErr, err)
		}
	}
	memories, err := r.npcMemoryRepo.List(ctx, &repo.NpcMemoryQuery{
		WorldID:     new(npc.WorldID),
		NpcID:       new(npc.ID),
		RecentLimit: 8,
	})
	if err == nil {
		return memories, nil
	}
	r.log.WarnContext(ctx, "load ready npc memories failed; fallback to no memories", "world_id", npc.WorldID, "npc_id", npc.ID, constant.LogFieldErr, err)
	return nil, nil
}

func actionTargets(
	payload map[string]any,
) []model.EntityRef {
	if payload == nil {
		return nil
	}
	values, _ := payload["targets"].([]any)
	result := make([]model.EntityRef, 0, len(values))
	for _, value := range values {
		item, ok := value.(map[string]any)
		if !ok {
			continue
		}
		kind, ok := item["type"].(string)
		if !ok {
			continue
		}
		id := entityIDValue(item["id"])
		if id <= 0 {
			continue
		}
		result = append(result, model.EntityRef{
			Type: enum.EntityType(kind),
			ID:   id,
		})
	}
	slices.SortFunc(result, func(a, b model.EntityRef) int {
		if a.Type != b.Type {
			return strings.Compare(string(a.Type), string(b.Type))
		}
		return int(a.ID - b.ID)
	})
	return result
}

func entityIDValue(
	value any,
) int64 {
	switch typedValue := value.(type) {
	case int:
		return int64(typedValue)
	case int8:
		return int64(typedValue)
	case int16:
		return int64(typedValue)
	case int32:
		return int64(typedValue)
	case int64:
		return typedValue
	case uint:
		return int64(typedValue)
	case uint8:
		return int64(typedValue)
	case uint16:
		return int64(typedValue)
	case uint32:
		return int64(typedValue)
	case uint64:
		if typedValue > uint64(^uint64(0)>>1) {
			return 0
		}
		return int64(typedValue)
	case float32:
		return int64(typedValue)
	case float64:
		return int64(typedValue)
	case string:
		return parseEntityIDString(typedValue)
	case fmt.Stringer:
		return parseEntityIDString(typedValue.String())
	default:
		return 0
	}
}

func parseEntityIDString(
	value string,
) int64 {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}
	result, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		return result
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return int64(floatValue)
}

func fallbackNpcReply(
	result *agentResult,
	err error,
) *model.NpcReply {
	reply := "我需要一点时间整理线索。你可以先说明更具体的目标，或换一个方向调查。"
	if result != nil && result.npc != nil {
		reply = result.npc.Name + "短暂沉思后说：" + reply
	}
	return &model.NpcReply{
		Reply:          reply,
		ContextSummary: fallbackContextSummary(result, err),
		SuggestedActions: []model.SuggestedAction{
			{Label: "追问线索", Content: "请告诉我你目前最确定的一条线索。"},
			{Label: "换个问题", Content: "我想从另一个角度调查这件事。"},
		},
	}
}

func fallbackWorldDraft(
	result *agentResult,
) *model.WorldDraft {
	description := "一个正在自行演化的文字世界"
	if result != nil && result.world != nil && strings.TrimSpace(result.world.Description) != "" {
		description = result.world.Description
	}
	return &model.WorldDraft{
		Name:       "未命名边境",
		Summary:    normalizeModelText(description, 120),
		CurrentArc: "新局势正在形成",
		CurrentEra: "初始纪元",
		Rules: map[string]any{
			"time_scale": 24,
		},
		Locations: []model.WorldDraftLocation{
			{Code: "center", Name: "中心据点", Description: "人们交换消息、寻找机会和应对危机的地方。", EnvironmentTags: []string{"public", "hub"}},
			{Code: "wilds", Name: "荒野边缘", Description: "危险与资源并存，局势会随灾害和行动改变。", EnvironmentTags: []string{"danger", "resource"}},
			{Code: "ruins", Name: "旧日遗迹", Description: "埋藏秘密、失踪线索和势力争夺目标。", EnvironmentTags: []string{"secret", "ancient"}},
		},
		Factions: []model.WorldDraftFaction{
			{Code: "council", Name: "地方议会", Description: "维持秩序并控制公开资源。", PublicGoal: "稳定局势"},
			{Code: "seekers", Name: "寻迹者同盟", Description: "追查异常与失踪事件。", PublicGoal: "揭开真相"},
		},
		Npcs: []model.WorldDraftNpc{
			{Code: "warden", Name: "守望者", Role: "秩序维护者", Species: "人类", Personality: "谨慎", Goal: "防止危机扩散", Background: "长期守护中心据点。", LocationCode: "center", FactionCode: "council", SystemPrompt: "你是谨慎的守望者，只依据已知事实回应。"},
			{Code: "scout", Name: "斥候", Role: "情报收集者", Species: "人类", Personality: "敏锐", Goal: "调查异常源头", Background: "经常往返荒野与据点。", LocationCode: "wilds", FactionCode: "seekers", SystemPrompt: "你是敏锐的斥候，只说自己观察到的内容。"},
			{Code: "keeper", Name: "遗迹看守", Role: "秘密保管者", Species: "未知", Personality: "沉默", Goal: "守住旧日秘密", Background: "与旧日遗迹存在长期联系。", LocationCode: "ruins", FactionCode: "seekers", SystemPrompt: "你是沉默的遗迹看守，不泄漏未知秘密。"},
			{Code: "trader", Name: "行商", Role: "资源中介", Species: "人类", Personality: "精明", Goal: "在动荡中寻找交易机会", Background: "掌握多方传闻但真假混杂。", LocationCode: "center", FactionCode: "council", SystemPrompt: "你是精明的行商，区分亲眼所见和听来的传闻。"},
		},
	}
}

func fallbackPlayerActionResolution(
	result *agentResult,
	err error,
) *model.ActionResolution {
	summary := "玩家行动已被记录，世界暂时需要更多信息来裁决细节"
	clarification := "请补充你的具体目标、对象或愿意承担的代价。"
	if result != nil && result.source != nil && strings.TrimSpace(result.source.Content) != "" {
		summary = normalizeModelText("玩家尝试："+result.source.Content, 120)
	}
	return &model.ActionResolution{
		Status:        "clarification",
		Summary:       summary,
		Clarification: clarification,
		WorldSummary:  fallbackWorldSummary(result),
		CurrentArc:    fallbackCurrentArc(result),
		SuggestedActions: []model.SuggestedAction{
			{Label: "说明目标", Content: "我的目标是获得一个明确线索，并愿意付出合理代价。"},
			{Label: "谨慎行动", Content: "我先观察周围反应，不贸然推进。"},
		},
	}
}

func deterministicPlayerActionResolution(
	result *agentResult,
) *model.ActionResolution {
	content := ""
	if result != nil && result.source != nil {
		content = strings.TrimSpace(result.source.Content)
	}
	if content == "" {
		content = "玩家采取了一个未说明细节的行动"
	}
	summary := normalizeModelText("玩家行动产生影响："+content, 120)
	actions := make([]model.ActionStep, 0, 2)
	if result != nil && result.source != nil {
		for _, target := range actionTargets(result.source.Payload) {
			if target.Type == enum.EntityTypeLocation {
				actions = append(actions, model.ActionStep{
					Type: "move_player",
					Target: &model.EntityRef{
						Type: enum.EntityTypeLocation,
						ID:   target.ID,
					},
				})
				break
			}
		}
	}
	if len(actions) == 0 && result != nil && result.location != nil {
		actions = append(actions, model.ActionStep{
			Type: "change_location",
			Target: &model.EntityRef{
				Type: enum.EntityTypeLocation,
				ID:   result.location.ID,
			},
			Parameters: map[string]any{
				"description":      normalizeModelText(result.location.Description+"玩家行动让这里的局势出现新的变化。", 180),
				"environment_tags": []string{"player_influenced"},
			},
			DurationMinutes: 30,
		})
	}
	return &model.ActionResolution{
		Status:       "resolved",
		Summary:      summary,
		WorldSummary: fallbackWorldSummary(result),
		CurrentArc:   fallbackCurrentArc(result),
		Actions:      actions,
		SuggestedActions: []model.SuggestedAction{
			{Label: "继续推进", Content: "我继续沿着这个方向扩大影响。"},
			{Label: "观察后果", Content: "我先观察各方对这件事的反应。"},
		},
	}
}

func fallbackNpcPlan(
	result *agentResult,
	err error,
) *model.NpcPlan {
	summary := "NPC 根据当前局势维持原计划，并继续观察变化。"
	goal := "观察局势"
	actions := make([]model.ActionStep, 0, 1)
	if result != nil && result.npc != nil {
		if strings.TrimSpace(result.npc.Goal) != "" {
			goal = result.npc.Goal
		}
		summary = result.npc.Name + "继续围绕当前目标行动。"
		for _, location := range result.locations {
			if location == nil || location.ID == result.npc.CurrentLocationID || !location.Accessible {
				continue
			}
			actions = append(actions, model.ActionStep{
				Type: "move_npc",
				Target: &model.EntityRef{
					Type: enum.EntityTypeLocation,
					ID:   location.ID,
				},
				Parameters: map[string]any{
					"npc_id":      result.npc.ID,
					"location_id": location.ID,
				},
				DurationMinutes: 60,
			})
			break
		}
	}
	return &model.NpcPlan{
		Summary:        summary,
		Goal:           goal,
		NextDecisionIn: 24 * 60,
		Actions:        actions,
	}
}

func fallbackWorldTickResolution(
	result *agentResult,
	err error,
) *model.ActionResolution {
	summary := "世界在没有玩家直接推动时继续缓慢演化。"
	if result != nil && result.state != nil && strings.TrimSpace(result.state.CurrentArc) != "" {
		summary = normalizeModelText(result.state.CurrentArc+"出现新的余波。", 120)
	}
	actions := make([]model.ActionStep, 0, 2)
	if result != nil && len(result.locations) > 0 {
		location := result.locations[0]
		actions = append(actions, model.ActionStep{
			Type: "change_location",
			Target: &model.EntityRef{
				Type: enum.EntityTypeLocation,
				ID:   location.ID,
			},
			Parameters: map[string]any{
				"description":      normalizeModelText(location.Description+"局势的余波让这里出现新的变化。", 180),
				"environment_tags": []string{"evolving"},
			},
			DurationMinutes: 60,
		})
	}
	if result != nil && len(result.factions) > 0 {
		faction := result.factions[0]
		actions = append(actions, model.ActionStep{
			Type: "change_faction",
			Target: &model.EntityRef{
				Type: enum.EntityTypeFaction,
				ID:   faction.ID,
			},
			Parameters: map[string]any{
				"public_goal": normalizeModelText(faction.PublicGoal+"，并重新评估近期风险。", 120),
			},
			DurationMinutes: 60,
		})
	}
	return &model.ActionResolution{
		Status:       "resolved",
		Summary:      summary,
		WorldSummary: fallbackWorldSummary(result),
		CurrentArc:   fallbackCurrentArc(result),
		Actions:      actions,
	}
}

func fallbackWorldSummary(
	result *agentResult,
) string {
	if result != nil && result.state != nil && strings.TrimSpace(result.state.Summary) != "" {
		return result.state.Summary
	}
	if result != nil && result.world != nil && strings.TrimSpace(result.world.GenerationSummary) != "" {
		return result.world.GenerationSummary
	}
	return "世界局势仍在持续变化。"
}

func fallbackCurrentArc(
	result *agentResult,
) string {
	if result != nil && result.state != nil && strings.TrimSpace(result.state.CurrentArc) != "" {
		return normalizeModelText(result.state.CurrentArc, 80)
	}
	return "暗流涌动"
}

func fallbackContextSummary(
	result *agentResult,
	err error,
) string {
	if result != nil && result.npc != nil && strings.TrimSpace(result.npc.ContextSummary) != "" {
		return normalizeModelText(result.npc.ContextSummary, 120)
	}
	return "模型响应不稳定，NPC 保留当前认知并等待更多可观察事实。"
}
