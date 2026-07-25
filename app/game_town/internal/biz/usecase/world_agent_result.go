package usecase

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
	"time"

	"common/pkg/constant"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

const (
	maxWorldNameRunes    = 128
	maxLocationNameRunes = 128
	maxNpcNameRunes      = 128
	maxNpcRoleRunes      = 128
	maxCurrentArcRunes   = 256
	maxEventSummaryRunes = 512
)

func (r *WorldAgentRunner) applyResult(
	ctx context.Context,
	result *agentResult,
) {
	if result.err != nil {
		r.log.WarnContext(ctx, "apply agent result failed before transaction", constant.LogFieldTaskID, result.job.ID, "world_id", result.job.WorldID, "type", result.job.Type, constant.LogFieldErr, result.err)
		r.handleFailure(ctx, result)
		return
	}

	var published *model.Event
	err := r.tx(ctx, func(ctx context.Context) error {
		var err error
		switch result.job.Type {
		case enum.AgentJobTypeWorldGenerate:
			published, err = r.applyWorldDraft(ctx, result)
		case enum.AgentJobTypeNpcTalk:
			published, err = r.applyNpcReply(ctx, result)
		case enum.AgentJobTypePlayerCharacterGenerate:
			published, err = r.applyPlayerCharacter(ctx, result)
		case enum.AgentJobTypePlayerAct, enum.AgentJobTypePlayerActionInterpret:
			if result.reply != nil {
				published, err = r.applyNpcReply(ctx, result)
			} else {
				published, err = r.applyPlayerActionResult(ctx, result)
			}
		case enum.AgentJobTypeNpcPlan:
			published, err = r.applyNpcPlan(ctx, result)
		case enum.AgentJobTypeWorldTick:
			published, err = r.applyResolution(ctx, result, enum.EventTypeWorldEvolved)
		case enum.AgentJobTypeMemoryEmbed:
			err = r.applyMemoryEmbedding(ctx, result)
		}
		if err != nil {
			return err
		}
		_, err = r.agentJobRepo.MarkSucceeded(ctx, result.job.ID, time.Now())
		return err
	})
	if err != nil {
		result.err = err
		r.log.WarnContext(ctx, "apply agent result transaction failed", constant.LogFieldTaskID, result.job.ID, "world_id", result.job.WorldID, "type", result.job.Type, constant.LogFieldErr, err)
		r.handleFailure(ctx, result)
		return
	}
	if result.config != nil {
		r.recordSuccess(result.config.ID)
	}
	r.log.InfoContext(ctx, "agent result applied", constant.LogFieldTaskID, result.job.ID, "world_id", result.job.WorldID, "type", result.job.Type)
	for _, event := range result.sideEvents {
		r.eventUsecase.Publish(event)
	}
	r.eventUsecase.Publish(published)
}

func (r *WorldAgentRunner) applyPlayerCharacter(
	ctx context.Context,
	result *agentResult,
) (*model.Event, error) {
	if result.character == nil || result.member == nil || result.player == nil {
		return nil, fmt.Errorf("player character context is incomplete")
	}
	name := normalizeModelText(result.character.Name, maxNpcNameRunes)
	if name == "" {
		name = "无名旅人"
	}
	background := normalizeModelText(result.character.Background, 1024)
	goal := normalizeModelText(result.character.Goal, maxEventSummaryRunes)
	traits := result.character.Traits
	if len(traits) > 6 {
		traits = traits[:6]
	}

	member, err := r.worldMemberRepo.UpdateCharacter(ctx, &repo.WorldMemberCharacterReq{
		MemberID:   result.member.ID,
		Name:       name,
		Background: background,
		Goal:       goal,
		Traits:     traits,
		Ready:      true,
	})
	if err != nil {
		return nil, err
	}

	return r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypePlayerCharacterReady,
		ActorPlayerID:    new(member.PlayerID),
		LocationID:       new(member.CurrentLocationID),
		CausationEventID: new(result.source.ID),
		Summary:          "玩家角色生成完成：" + name,
		Content:          background,
		Payload: map[string]any{
			"character_name":       name,
			"character_background": background,
			"character_goal":       goal,
			"character_traits":     traits,
		},
	})
}

func (r *WorldAgentRunner) applyWorldDraft(
	ctx context.Context,
	result *agentResult,
) (*model.Event, error) {
	draft := normalizeWorldDraft(result)
	locationCount := int(uint32Value(result.source.Payload, "location_count"))
	npcCount := int(uint32Value(result.source.Payload, "npc_count"))
	if locationCount <= 0 {
		locationCount = len(draft.Locations)
	}
	if npcCount <= 0 {
		npcCount = len(draft.Npcs)
	}
	ensureWorldDraftCounts(draft, locationCount, npcCount)
	if len(draft.Locations) > locationCount {
		draft.Locations = draft.Locations[:locationCount]
	}
	if len(draft.Npcs) > npcCount {
		draft.Npcs = draft.Npcs[:npcCount]
	}

	draft.Name = normalizeModelText(draft.Name, maxWorldNameRunes)
	draft.CurrentArc = normalizeModelText(draft.CurrentArc, maxCurrentArcRunes)
	if draft.Name == "" {
		return nil, fmt.Errorf("world name is empty")
	}

	locations := make(map[string]*model.Location, len(draft.Locations))
	var firstLocation *model.Location
	for index, item := range draft.Locations {
		code := normalizedDraftCode(item.Code, fmt.Sprintf("location_%d", index+1), locations)
		name := normalizeModelText(item.Name, maxLocationNameRunes)
		if name == "" {
			name = fmt.Sprintf("地点%d", index+1)
		}
		location, err := r.locationRepo.Save(ctx, &model.Location{
			WorldID:         result.world.ID,
			Code:            code,
			Name:            name,
			Description:     strings.TrimSpace(item.Description),
			Status:          enum.LocationStatusActive,
			EnvironmentTags: item.EnvironmentTags,
			Attributes:      map[string]any{},
			Accessible:      true,
			Sort:            int32(index),
		})
		if err != nil {
			return nil, err
		}
		locations[code] = location
		if firstLocation == nil {
			firstLocation = location
		}
	}
	if firstLocation == nil {
		return nil, fmt.Errorf("world has no location")
	}

	factions, err := r.saveDraftFactions(ctx, result.world.ID, draft)
	if err != nil {
		return nil, err
	}

	npcCodes := make(map[string]struct{}, len(draft.Npcs))
	for index, item := range draft.Npcs {
		code := normalizedNpcCode(item.Code, fmt.Sprintf("npc_%d", index+1), npcCodes)
		name := normalizeModelText(item.Name, maxNpcNameRunes)
		role := normalizeModelText(item.Role, maxNpcRoleRunes)
		location := locations[strings.TrimSpace(item.LocationCode)]
		if name == "" {
			name = fmt.Sprintf("角色%d", index+1)
		}
		if role == "" {
			role = "世界居民"
		}
		if location == nil {
			location = firstLocation
		}
		npcCodes[code] = struct{}{}
		saved, err := r.npcRepo.Save(ctx, &model.Npc{
			WorldID:           result.world.ID,
			Code:              code,
			Name:              name,
			Role:              role,
			Personality:       strings.TrimSpace(item.Personality),
			Goal:              strings.TrimSpace(item.Goal),
			Background:        strings.TrimSpace(item.Background),
			CurrentLocationID: location.ID,
			SystemPrompt:      strings.TrimSpace(item.SystemPrompt),
			Species:           strings.TrimSpace(item.Species),
			LifeStatus:        enum.NpcLifeStatusAlive,
			Attributes:        item.Attributes,
			NextDecisionAt:    new(result.state.WorldTime),
			StateTags:         []string{"active"},
		})
		if err != nil {
			return nil, err
		}
		if faction := factions[strings.TrimSpace(item.FactionCode)]; faction != nil {
			_, err = r.factionMembershipRepo.Save(ctx, &model.FactionMembership{
				WorldID:    result.world.ID,
				FactionID:  faction.ID,
				MemberType: enum.EntityTypeNpc,
				MemberID:   saved.ID,
				Role:       role,
				Reputation: map[string]float64{},
				Tags:       []string{"founding_member"},
				JoinedAt:   result.state.WorldTime,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	if _, err = r.worldRuleRepo.Save(ctx, &model.WorldRule{
		WorldID: result.world.ID,
		Version: 1,
		Rules:   defaultWorldRules(draft.Rules),
	}); err != nil {
		return nil, err
	}

	result.world.Name = draft.Name
	result.world.GenerationSummary = strings.TrimSpace(draft.Summary)
	result.world.Status = enum.WorldStatusActive
	result.world.DefaultLocationID = new(firstLocation.ID)
	if _, err = r.worldRepo.Update(ctx, result.world); err != nil {
		return nil, err
	}

	nextTickAt := time.Now().Add(r.tickInterval())
	if _, err = r.worldStateRepo.UpdateNarrative(ctx, &repo.WorldStateUpdateNarrativeReq{
		WorldID:         result.world.ID,
		Version:         result.state.Version,
		Summary:         strings.TrimSpace(draft.Summary),
		CurrentArc:      draft.CurrentArc,
		NextTickAt:      new(nextTickAt),
		PublicChronicle: strings.TrimSpace(draft.Summary),
		CurrentEra:      normalizeModelText(draft.CurrentEra, 128),
		NextDueAt:       new(nextTickAt),
	}); err != nil {
		return nil, err
	}

	return r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypeWorldReady,
		ActorPlayerID:    new(result.world.CreatorPlayerID),
		LocationID:       new(firstLocation.ID),
		CausationEventID: new(result.source.ID),
		Summary:          "世界生成完成",
		Content:          strings.TrimSpace(draft.Summary),
		Payload:          map[string]any{"public": true},
	})
}

func (r *WorldAgentRunner) applyNpcReply(
	ctx context.Context,
	result *agentResult,
) (*model.Event, error) {
	if result.reply == nil || result.npc == nil || result.location == nil {
		return nil, fmt.Errorf("npc reply context is incomplete")
	}

	contextSummary := strings.TrimSpace(result.reply.ContextSummary)
	if contextSummary == "" {
		contextSummary = result.npc.ContextSummary
	}
	if _, err := r.npcRepo.UpdateContext(ctx, result.npc.ID, result.npc.Version, contextSummary); err != nil {
		if !isNpcVersionConflict(err) {
			return nil, err
		}
	}

	event, err := r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypeNpcReplied,
		ActorPlayerID:    result.source.ActorPlayerID,
		NpcID:            new(result.npc.ID),
		LocationID:       new(result.location.ID),
		CausationEventID: new(result.source.ID),
		Summary:          result.npc.Name + " 回复",
		Content:          strings.TrimSpace(result.reply.Reply),
		Payload: map[string]any{
			"suggested_actions": suggestedActionPayload(result.reply.SuggestedActions),
			"claims":            claimDraftPayload(result.reply.Claims),
		},
	})
	if err != nil {
		return nil, err
	}
	if err = r.saveClaims(ctx, event, result.npc, result.reply.Claims); err != nil {
		return nil, err
	}
	if strings.TrimSpace(result.reply.Reply) != "" {
		if err = r.saveNpcConversationMemory(ctx, result.npc, event, result.reply.Reply); err != nil {
			return nil, err
		}
	}
	return event, nil
}

func (r *WorldAgentRunner) applyResolution(
	ctx context.Context,
	result *agentResult,
	eventType enum.EventType,
) (*model.Event, error) {
	if result.resolution == nil {
		return nil, fmt.Errorf("action resolution is nil")
	}

	worldSummary := strings.TrimSpace(result.resolution.WorldSummary)
	if worldSummary == "" {
		worldSummary = result.state.Summary
	}
	currentArc := normalizeModelText(result.resolution.CurrentArc, maxCurrentArcRunes)
	if currentArc == "" {
		currentArc = normalizeModelText(result.state.CurrentArc, maxCurrentArcRunes)
	}
	eventSummary := normalizeModelText(result.resolution.Summary, maxEventSummaryRunes)
	if eventSummary == "" {
		eventSummary = "世界继续演进"
		if eventType == enum.EventTypeActionResolved {
			eventSummary = "玩家行动已处理"
		}
	}
	eventContent := strings.TrimSpace(result.resolution.Summary)

	if eventType == enum.EventTypeWorldEvolved {
		duplicated, err := r.isDuplicateWorldEvolution(ctx, result.world.ID, eventSummary, eventContent)
		if err != nil {
			return nil, err
		}
		if duplicated {
			nextTickAt := time.Now().Add(r.tickInterval())
			if err = r.worldStateRepo.UpdateNextTick(ctx, result.world.ID, nextTickAt); err != nil {
				return nil, err
			}
			r.log.InfoContext(ctx, "skip duplicate world evolution", "world_id", result.world.ID, "summary", eventSummary)
			return nil, nil
		}
	}

	if err := r.applyActionSteps(ctx, result, result.resolution.Actions); err != nil {
		return nil, err
	}

	nextTickAt := time.Now().Add(r.tickInterval())
	if _, err := r.worldStateRepo.UpdateNarrative(ctx, &repo.WorldStateUpdateNarrativeReq{
		WorldID:         result.world.ID,
		Version:         result.state.Version,
		Summary:         worldSummary,
		CurrentArc:      currentArc,
		PublicChronicle: result.state.PublicChronicle,
		CurrentEra:      result.state.CurrentEra,
		NextTickAt:      new(nextTickAt),
		NextDueAt:       new(nextTickAt),
	}); err != nil {
		if !isWorldVersionConflict(err) {
			return nil, err
		}
		if err = r.worldStateRepo.UpdateNextTick(ctx, result.world.ID, nextTickAt); err != nil {
			return nil, err
		}
	}

	event, err := r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             eventType,
		ActorPlayerID:    result.source.ActorPlayerID,
		LocationID:       result.source.LocationID,
		CausationEventID: new(result.source.ID),
		Summary:          eventSummary,
		Content:          eventContent,
		Payload: map[string]any{
			"actions": actionStepPayload(result.resolution.Actions),
			"claims":  claimDraftPayload(result.resolution.Claims),
			"public":  eventType == enum.EventTypeWorldEvolved,
		},
	})
	if err != nil {
		return nil, err
	}
	if err = r.saveClaims(ctx, event, result.npc, result.resolution.Claims); err != nil {
		return nil, err
	}
	return event, nil
}

func (r *WorldAgentRunner) isDuplicateWorldEvolution(
	ctx context.Context,
	worldID int64,
	summary string,
	content string,
) (bool, error) {
	eventType := enum.EventTypeWorldEvolved
	events, err := r.eventRepo.List(ctx, &repo.EventQuery{
		WorldID:     new(worldID),
		Type:        new(eventType),
		RecentLimit: 5,
	})
	if err != nil {
		return false, err
	}

	summary = strings.TrimSpace(summary)
	content = strings.TrimSpace(content)
	for _, event := range events {
		if event == nil {
			continue
		}
		if strings.TrimSpace(event.Summary) == summary && strings.TrimSpace(event.Content) == content {
			return true, nil
		}
	}
	return false, nil
}

func normalizeModelText(
	value string,
	maxRunes int,
) string {
	value = strings.TrimSpace(value)
	runes := []rune(value)
	if len(runes) > maxRunes {
		return string(runes[:maxRunes])
	}
	return value
}

func normalizeWorldDraft(
	result *agentResult,
) *model.WorldDraft {
	if result == nil || result.draft == nil {
		return fallbackWorldDraft(result)
	}
	draft := result.draft
	fallback := fallbackWorldDraft(result)
	if strings.TrimSpace(draft.Name) == "" {
		draft.Name = fallback.Name
	}
	if strings.TrimSpace(draft.Summary) == "" {
		draft.Summary = fallback.Summary
	}
	if strings.TrimSpace(draft.CurrentArc) == "" {
		draft.CurrentArc = fallback.CurrentArc
	}
	if strings.TrimSpace(draft.CurrentEra) == "" {
		draft.CurrentEra = fallback.CurrentEra
	}
	if len(draft.Locations) == 0 {
		draft.Locations = fallback.Locations
	}
	if len(draft.Factions) == 0 {
		draft.Factions = fallback.Factions
	}
	if len(draft.Npcs) == 0 {
		draft.Npcs = fallback.Npcs
	}
	return draft
}

func ensureWorldDraftCounts(
	draft *model.WorldDraft,
	locationCount int,
	npcCount int,
) {
	fallback := fallbackWorldDraft(nil)
	if locationCount <= 0 && npcCount > 0 {
		locationCount = 1
	}
	for len(draft.Locations) < locationCount {
		index := len(draft.Locations)
		if index < len(fallback.Locations) {
			draft.Locations = append(draft.Locations, fallback.Locations[index])
			continue
		}
		draft.Locations = append(draft.Locations, model.WorldDraftLocation{
			Code:        fmt.Sprintf("location_%d", index+1),
			Name:        fmt.Sprintf("地点%d", index+1),
			Description: "一个会随世界局势变化的地点。",
		})
	}
	for len(draft.Npcs) < npcCount {
		index := len(draft.Npcs)
		if index < len(fallback.Npcs) {
			draft.Npcs = append(draft.Npcs, fallback.Npcs[index])
			continue
		}
		draft.Npcs = append(draft.Npcs, model.WorldDraftNpc{
			Code:         fmt.Sprintf("npc_%d", index+1),
			Name:         fmt.Sprintf("角色%d", index+1),
			Role:         "世界居民",
			Species:      "人类",
			Personality:  "谨慎",
			Goal:         "适应变化",
			Background:   "在动荡中寻找自己的位置。",
			LocationCode: draft.Locations[0].Code,
			SystemPrompt: "你是世界中的独立居民，只依据自己知道的信息回应。",
		})
	}
}

func normalizedDraftCode(
	value string,
	fallback string,
	existing map[string]*model.Location,
) string {
	code := normalizeASCIIIdentifier(value)
	if code == "" {
		code = fallback
	}
	for index := 2; existing[code] != nil; index++ {
		code = fmt.Sprintf("%s_%d", fallback, index)
	}
	return normalizeModelText(code, 64)
}

func normalizedNpcCode(
	value string,
	fallback string,
	existing map[string]struct{},
) string {
	code := normalizeASCIIIdentifier(value)
	if code == "" {
		code = fallback
	}
	for index := 2; ; index++ {
		if _, ok := existing[code]; !ok {
			return normalizeModelText(code, 64)
		}
		code = fmt.Sprintf("%s_%d", fallback, index)
	}
}

func normalizeASCIIIdentifier(
	value string,
) string {
	value = strings.ToLower(strings.TrimSpace(value))
	var builder strings.Builder
	for _, char := range value {
		switch {
		case char >= 'a' && char <= 'z':
			builder.WriteRune(char)
		case char >= '0' && char <= '9':
			builder.WriteRune(char)
		case char == '_' || char == '-':
			builder.WriteRune('_')
		}
	}
	return strings.Trim(builder.String(), "_")
}

func (r *WorldAgentRunner) handleFailure(
	ctx context.Context,
	result *agentResult,
) {
	r.recordFailure(result.config)
	if result.job.AttemptCount <= int32(r.maxRetry()) {
		delay := time.Duration(float64(r.retryDelay()) * math.Pow(2, float64(result.job.AttemptCount-1)))
		jitterRange := delay / 2
		if jitterRange > 0 {
			delay += time.Duration(rand.Int64N(int64(jitterRange)+1)) - delay/4
		}
		r.log.WarnContext(ctx, "retry agent job later", constant.LogFieldTaskID, result.job.ID, "world_id", result.job.WorldID, "type", result.job.Type, "delay", delay.String(), constant.LogFieldErr, result.err)
		if r.persistAgentJobRetry(ctx, result, delay) {
			r.publishRetryEvent(ctx, result, delay)
			time.AfterFunc(delay, r.wakeScheduler)
			return
		}
		r.scheduleRetryPersistence(result, delay, 1)
		return
	}

	var event *model.Event
	if err := r.tx(ctx, func(ctx context.Context) error {
		if _, err := r.agentJobRepo.MarkFailed(ctx, &repo.AgentJobMarkFailedReq{
			JobID:        result.job.ID,
			FinishedAt:   time.Now(),
			ErrorSummary: truncateError(result.err),
		}); err != nil {
			return err
		}
		if result.job.Type == enum.AgentJobTypeWorldTick {
			if err := r.worldStateRepo.UpdateNextTick(ctx, result.job.WorldID, time.Now().Add(r.tickInterval())); err != nil {
				return err
			}
		}

		eventType := enum.EventTypeAgentJobFailed
		summary := "Agent 任务失败"
		if result.job.Type == enum.AgentJobTypePlayerCharacterGenerate {
			eventType = enum.EventTypePlayerCharacterFailed
			summary = "玩家角色生成失败"
		}
		if result.job.Type == enum.AgentJobTypeWorldGenerate && result.world != nil {
			result.world.Status = enum.WorldStatusFailed
			result.world.GenerationSummary = truncateError(result.err)
			if _, err := r.worldRepo.Update(ctx, result.world); err != nil {
				return err
			}
			eventType = enum.EventTypeWorldGenerationFailed
			summary = "世界生成失败"
		}

		var causationEventID *int64
		var actorPlayerID *int64
		var locationID *int64
		if result.source != nil {
			causationEventID = new(result.source.ID)
			actorPlayerID = result.source.ActorPlayerID
			locationID = result.source.LocationID
		}
		if causationEventID != nil {
			count, err := r.eventRepo.Count(ctx, &repo.EventQuery{
				WorldID:          new(result.job.WorldID),
				Type:             new(eventType),
				CausationEventID: causationEventID,
			})
			if err != nil {
				return err
			}
			if count > 0 {
				return nil
			}
		}

		var err error
		event, err = r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
			WorldID:          result.job.WorldID,
			Type:             eventType,
			ActorPlayerID:    actorPlayerID,
			LocationID:       locationID,
			CausationEventID: causationEventID,
			Summary:          summary,
			Content:          truncateError(result.err),
		})
		return err
	}); err != nil {
		r.log.ErrorContext(ctx, "persist agent job failure failed", constant.LogFieldTaskID, result.job.ID, constant.LogFieldErr, err)
		return
	}
	r.eventUsecase.Publish(event)
}

func (r *WorldAgentRunner) persistAgentJobRetry(
	ctx context.Context,
	result *agentResult,
	delay time.Duration,
) bool {
	_, err := r.agentJobRepo.Retry(ctx, &repo.AgentJobRetryReq{
		JobID:        result.job.ID,
		AttemptCount: result.job.AttemptCount,
		AvailableAt:  time.Now().Add(delay),
		ErrorSummary: truncateError(result.err),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "retry agent job failed", constant.LogFieldTaskID, result.job.ID, constant.LogFieldErr, err)
		return false
	}
	return true
}

func (r *WorldAgentRunner) scheduleRetryPersistence(
	result *agentResult,
	delay time.Duration,
	attempt int,
) {
	if result == nil || result.job == nil || attempt > 6 {
		return
	}
	wait := time.Duration(attempt) * r.retryDelay()
	if wait > 30*time.Second {
		wait = 30 * time.Second
	}
	time.AfterFunc(wait, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if !r.persistAgentJobRetry(ctx, result, delay) {
			r.scheduleRetryPersistence(result, delay, attempt+1)
			return
		}
		r.publishRetryEvent(ctx, result, delay)
		time.AfterFunc(delay, r.wakeScheduler)
	})
}

func (r *WorldAgentRunner) publishRetryEvent(
	ctx context.Context,
	result *agentResult,
	delay time.Duration,
) {
	r.log.WarnContext(
		ctx,
		"agent job will retry",
		constant.LogFieldTaskID,
		result.job.ID,
		"world_id",
		result.job.WorldID,
		"type",
		result.job.Type,
		"delay",
		delay.String(),
		constant.LogFieldErr,
		result.err,
	)
}
