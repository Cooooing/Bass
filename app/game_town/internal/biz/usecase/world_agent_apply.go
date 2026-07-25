package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"common/pkg/constant"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

func (r *WorldAgentRunner) saveNpcConversationMemory(ctx context.Context, npc *model.Npc, event *model.Event, content string) error {
	embeddingStatus := enum.EmbeddingStatusFailed
	embeddingError := "embedding disabled"
	if r.embeddingEnabled() {
		embeddingStatus = enum.EmbeddingStatusPending
		embeddingError = ""
	}

	memory, err := r.npcMemoryRepo.Save(ctx, &model.NpcMemory{
		WorldID:           npc.WorldID,
		NpcID:             npc.ID,
		SourceEventID:     new(event.ID),
		Kind:              enum.MemoryKindConversation,
		Content:           npc.Name + " 与玩家交互：" + strings.TrimSpace(content),
		Importance:        0.7,
		OccurredWorldTime: event.WorldTime,
		EmbeddingStatus:   embeddingStatus,
		EmbeddingError:    embeddingError,
	})
	if err != nil {
		return err
	}
	if !r.embeddingEnabled() {
		return nil
	}
	return r.enqueueMemoryEmbedding(ctx, event, memory)
}

func (r *WorldAgentRunner) enqueueMemoryEmbedding(ctx context.Context, event *model.Event, memory *model.NpcMemory) error {
	jobType := enum.AgentJobTypeMemoryEmbed
	count, err := r.agentJobRepo.Count(ctx, &repo.AgentJobQuery{
		SourceEventID: new(event.ID),
		Type:          new(jobType),
	})
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	_, err = r.agentJobRepo.Save(ctx, &model.AgentJob{
		WorldID:       event.WorldID,
		SourceEventID: event.ID,
		Type:          enum.AgentJobTypeMemoryEmbed,
		Priority:      enum.AgentJobPriorityLow,
		LaneKey:       fmt.Sprintf("memory:npc:%d", memory.NpcID),
		Status:        enum.AgentJobStatusQueued,
		WorldVersion:  0,
		NpcID:         new(memory.NpcID),
		AvailableAt:   time.Now(),
	})
	return err
}

func (r *WorldAgentRunner) applyMemoryEmbedding(ctx context.Context, result *agentResult) error {
	if result.memory == nil {
		return nil
	}
	modelName := r.conf.GetGameTown().GetMemory().GetModel()
	return r.npcMemoryRepo.SetEmbedding(ctx, &repo.NpcMemoryEmbeddingReq{
		ID:           result.memory.ID,
		Vector:       result.embedding,
		Model:        modelName,
		Status:       enum.EmbeddingStatusReady,
		ErrorSummary: "",
	})
}

func (r *WorldAgentRunner) saveClaims(ctx context.Context, event *model.Event, sourceNpc *model.Npc, drafts []model.ClaimDraft) error {
	for _, draft := range drafts {
		predicate := strings.TrimSpace(draft.Predicate)
		if predicate == "" || draft.SubjectID <= 0 || draft.SubjectType == "" {
			continue
		}
		truth := draft.Truth
		if truth == "" {
			truth = enum.ClaimTruthUnknown
		}
		object := draft.Object
		if object == nil {
			object = map[string]any{}
		}
		claim, err := r.claimRepo.Save(ctx, &model.Claim{
			WorldID:       event.WorldID,
			OriginEventID: new(event.ID),
			SubjectType:   draft.SubjectType,
			SubjectID:     draft.SubjectID,
			Predicate:     predicate,
			Object:        object,
			Truth:         truth,
		})
		if err != nil {
			return err
		}
		if sourceNpc == nil {
			continue
		}
		_, err = r.npcBeliefRepo.Save(ctx, &model.NpcBelief{
			WorldID:        event.WorldID,
			NpcID:          sourceNpc.ID,
			ClaimID:        claim.ID,
			SourceNpcID:    new(sourceNpc.ID),
			SourcePlayerID: event.ActorPlayerID,
			SourceEventID:  new(event.ID),
			Stance:         enum.BeliefStanceBelieves,
			Confidence:     0.75,
			LearnedAt:      event.WorldTime,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WorldAgentRunner) applyActionSteps(ctx context.Context, result *agentResult, steps []model.ActionStep) error {
	for _, step := range steps {
		if err := r.applyActionStep(ctx, result, step); err != nil {
			if ctx.Err() != nil {
				return err
			}
			r.log.WarnContext(
				ctx,
				"proposed action rejected",
				"world_id",
				result.world.ID,
				"action_type",
				step.Type,
				constant.LogFieldErr,
				err,
			)
		}
	}
	return nil
}

func (r *WorldAgentRunner) applyActionStep(ctx context.Context, result *agentResult, step model.ActionStep) error {
	actionType := strings.TrimSpace(step.Type)
	switch actionType {
	case "move", "move_player":
		return r.applyPlayerMove(ctx, result, step)
	case "move_npc":
		return r.applyNpcMove(ctx, result, step)
	case "change_npc_state", "npc_state":
		return r.applyNpcStateChange(ctx, result, step)
	case "change_location", "location_state":
		return r.applyLocationChange(ctx, result, step)
	case "change_faction", "faction_state":
		return r.applyFactionChange(ctx, result, step)
	case "change_relationship", "relationship":
		return r.applyRelationshipChange(ctx, result, step)
	case "share_claim":
		return r.applyClaimShare(ctx, result, step)
	default:
		return nil
	}
}

func (r *WorldAgentRunner) applyPlayerMove(ctx context.Context, result *agentResult, step model.ActionStep) error {
	if result.member == nil || step.Target == nil || step.Target.Type != enum.EntityTypeLocation {
		return nil
	}
	location, err := r.locationRepo.Get(ctx, &repo.LocationQuery{
		ID:      new(step.Target.ID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	if !location.Accessible || location.Status == enum.LocationStatusDestroyed {
		return fmt.Errorf("target location is inaccessible")
	}
	_, err = r.worldMemberRepo.Move(ctx, result.member.ID, location.ID)
	return err
}

func (r *WorldAgentRunner) applyNpcMove(ctx context.Context, result *agentResult, step model.ActionStep) error {
	npcID := int64Param(step.Parameters, "npc_id")
	if npcID == 0 && step.Target != nil && step.Target.Type == enum.EntityTypeNpc {
		npcID = step.Target.ID
	}
	if npcID == 0 && result.npc != nil {
		npcID = result.npc.ID
	}
	locationID := int64Param(step.Parameters, "location_id")
	if locationID == 0 && step.Target != nil && step.Target.Type == enum.EntityTypeLocation {
		locationID = step.Target.ID
	}
	if npcID == 0 || locationID == 0 {
		return nil
	}
	npc, err := r.npcRepo.Get(ctx, &repo.NpcQuery{
		ID:      new(npcID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	location, err := r.locationRepo.Get(ctx, &repo.LocationQuery{
		ID:      new(locationID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	if !location.Accessible || location.Status == enum.LocationStatusDestroyed {
		return fmt.Errorf("target location is inaccessible")
	}
	_, err = r.npcRepo.UpdateState(ctx, &repo.NpcStateUpdateReq{
		NpcID:             npc.ID,
		Version:           npc.Version,
		CurrentLocationID: new(location.ID),
	})
	if err != nil {
		return err
	}
	event, err := r.appendSideEventInTx(ctx, result, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypeNpcMoved,
		NpcID:            new(npc.ID),
		LocationID:       new(location.ID),
		CausationEventID: new(result.source.ID),
		Summary:          npc.Name + " 移动到 " + location.Name,
		Content:          npc.Name + " 前往 " + location.Name + "，这会影响后续可见事件和 NPC 交流。",
		Payload: map[string]any{
			"npc_id":      npc.ID,
			"location_id": location.ID,
		},
	})
	if err != nil {
		return err
	}
	result.sideEvents = append(result.sideEvents, event)
	return nil
}

func (r *WorldAgentRunner) applyNpcStateChange(ctx context.Context, result *agentResult, step model.ActionStep) error {
	npcID := targetID(step, enum.EntityTypeNpc)
	if npcID == 0 && result.npc != nil {
		npcID = result.npc.ID
	}
	if npcID == 0 {
		return nil
	}
	npc, err := r.npcRepo.Get(ctx, &repo.NpcQuery{
		ID:      new(npcID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	var lifeStatus *enum.NpcLifeStatus
	if raw := strings.TrimSpace(stringParam(step.Parameters, "life_status")); raw != "" {
		status := enum.NpcLifeStatus(raw)
		lifeStatus = new(status)
	}
	var deathWorldTime *time.Time
	if lifeStatus != nil && *lifeStatus == enum.NpcLifeStatusDead {
		deathWorldTime = new(result.state.WorldTime)
	}
	_, err = r.npcRepo.UpdateState(ctx, &repo.NpcStateUpdateReq{
		NpcID:          npc.ID,
		Version:        npc.Version,
		LifeStatus:     lifeStatus,
		StateTags:      stringSliceParam(step.Parameters, "state_tags"),
		Attributes:     mapParam(step.Parameters, "attributes"),
		DeathWorldTime: deathWorldTime,
	})
	if err != nil {
		return err
	}
	eventType := enum.EventTypeNpcStateChanged
	summary := npc.Name + " 状态发生变化"
	if lifeStatus != nil && *lifeStatus == enum.NpcLifeStatusDead {
		eventType = enum.EventTypeNpcDied
		summary = npc.Name + " 死亡"
	}
	event, err := r.appendSideEventInTx(ctx, result, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             eventType,
		NpcID:            new(npc.ID),
		LocationID:       new(npc.CurrentLocationID),
		CausationEventID: new(result.source.ID),
		Summary:          summary,
		Content:          summary,
		Payload: map[string]any{
			"npc_id":      npc.ID,
			"life_status": stringValue(lifeStatus),
			"state_tags":  stringSliceParam(step.Parameters, "state_tags"),
		},
	})
	if err != nil {
		return err
	}
	result.sideEvents = append(result.sideEvents, event)
	return nil
}

func (r *WorldAgentRunner) applyLocationChange(ctx context.Context, result *agentResult, step model.ActionStep) error {
	locationID := targetID(step, enum.EntityTypeLocation)
	if locationID == 0 {
		return nil
	}
	location, err := r.locationRepo.Get(ctx, &repo.LocationQuery{
		ID:      new(locationID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	var status *enum.LocationStatus
	if raw := strings.TrimSpace(stringParam(step.Parameters, "status")); raw != "" {
		value := enum.LocationStatus(raw)
		status = new(value)
	}
	accessible := boolPtrParam(step.Parameters, "accessible")
	_, err = r.locationRepo.UpdateState(ctx, &repo.LocationStateUpdateReq{
		LocationID:      location.ID,
		Version:         location.Version,
		Status:          status,
		Description:     stringParam(step.Parameters, "description"),
		Accessible:      accessible,
		EnvironmentTags: stringSliceParam(step.Parameters, "environment_tags"),
		Attributes:      mapParam(step.Parameters, "attributes"),
	})
	if err != nil {
		return err
	}
	event, err := r.appendSideEventInTx(ctx, result, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypeLocationChanged,
		LocationID:       new(location.ID),
		CausationEventID: new(result.source.ID),
		Summary:          location.Name + " 发生变化",
		Content:          nonEmptyString(stringParam(step.Parameters, "description"), location.Description),
		Payload: map[string]any{
			"public":           true,
			"location_id":      location.ID,
			"status":           stringValue(status),
			"accessible":       boolPtrValue(accessible),
			"environment_tags": stringSliceParam(step.Parameters, "environment_tags"),
		},
	})
	if err != nil {
		return err
	}
	result.sideEvents = append(result.sideEvents, event)
	return nil
}

func (r *WorldAgentRunner) applyFactionChange(ctx context.Context, result *agentResult, step model.ActionStep) error {
	factionID := targetID(step, enum.EntityTypeFaction)
	if factionID == 0 {
		return nil
	}
	faction, err := r.factionRepo.Get(ctx, &repo.FactionQuery{
		ID:      new(factionID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	var status *enum.FactionStatus
	if raw := strings.TrimSpace(stringParam(step.Parameters, "status")); raw != "" {
		value := enum.FactionStatus(raw)
		status = new(value)
	}
	_, err = r.factionRepo.UpdateState(ctx, &repo.FactionStateUpdateReq{
		FactionID:   faction.ID,
		Version:     faction.Version,
		Status:      status,
		Description: stringParam(step.Parameters, "description"),
		PublicGoal:  stringParam(step.Parameters, "public_goal"),
		Attributes:  mapParam(step.Parameters, "attributes"),
	})
	if err != nil {
		return err
	}
	event, err := r.appendSideEventInTx(ctx, result, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypeFactionChanged,
		CausationEventID: new(result.source.ID),
		Summary:          faction.Name + " 发生变化",
		Content:          nonEmptyString(stringParam(step.Parameters, "public_goal"), faction.PublicGoal),
		Payload: map[string]any{
			"public":     true,
			"faction_id": faction.ID,
			"status":     stringValue(status),
		},
	})
	if err != nil {
		return err
	}
	result.sideEvents = append(result.sideEvents, event)
	return nil
}

func (r *WorldAgentRunner) applyRelationshipChange(ctx context.Context, result *agentResult, step model.ActionStep) error {
	sourceType := enum.EntityType(stringParam(step.Parameters, "source_type"))
	sourceID := int64Param(step.Parameters, "source_id")
	if sourceType == "" && result.npc != nil {
		sourceType = enum.EntityTypeNpc
		sourceID = result.npc.ID
	}
	if sourceType == "" && result.member != nil {
		sourceType = enum.EntityTypePlayer
		sourceID = result.member.PlayerID
	}
	targetType := enum.EntityType(stringParam(step.Parameters, "target_type"))
	targetID := int64Param(step.Parameters, "target_id")
	if step.Target != nil {
		targetType = step.Target.Type
		targetID = step.Target.ID
	}
	if sourceType == "" || sourceID == 0 || targetType == "" || targetID == 0 {
		return nil
	}
	_, err := r.relationshipRepo.Upsert(ctx, &repo.RelationshipUpsertReq{
		WorldID:    result.world.ID,
		SourceType: sourceType,
		SourceID:   sourceID,
		TargetType: targetType,
		TargetID:   targetID,
		Metrics:    floatMapParam(step.Parameters, "metrics"),
		Tags:       stringSliceParam(step.Parameters, "tags"),
	})
	return err
}

func (r *WorldAgentRunner) applyClaimShare(ctx context.Context, result *agentResult, step model.ActionStep) error {
	npcID := targetID(step, enum.EntityTypeNpc)
	claimID := int64Param(step.Parameters, "claim_id")
	if npcID == 0 || claimID == 0 {
		return nil
	}
	npc, err := r.npcRepo.Get(ctx, &repo.NpcQuery{
		ID:      new(npcID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	claim, err := r.claimRepo.Get(ctx, &repo.ClaimQuery{
		ID:      new(claimID),
		WorldID: new(result.world.ID),
	})
	if err != nil {
		return err
	}
	_, err = r.npcBeliefRepo.Save(ctx, &model.NpcBelief{
		WorldID:       result.world.ID,
		NpcID:         npc.ID,
		ClaimID:       claim.ID,
		SourceEventID: new(result.source.ID),
		Stance:        enum.BeliefStanceBelieves,
		Confidence:    0.55,
		LearnedAt:     result.state.WorldTime,
	})
	if err != nil {
		return err
	}
	event, err := r.appendSideEventInTx(ctx, result, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypeClaimShared,
		NpcID:            new(npc.ID),
		LocationID:       new(npc.CurrentLocationID),
		CausationEventID: new(result.source.ID),
		Summary:          npc.Name + " 获得一条新消息",
		Content:          claim.Predicate,
		Payload: map[string]any{
			"npc_id":   npc.ID,
			"claim_id": claim.ID,
		},
	})
	if err != nil {
		return err
	}
	result.sideEvents = append(result.sideEvents, event)
	return nil
}

func (r *WorldAgentRunner) appendSideEventInTx(ctx context.Context, result *agentResult, req *AppendEventReq) (*model.Event, error) {
	if result == nil || result.source == nil || req == nil {
		return nil, nil
	}
	return r.eventUsecase.AppendInTx(ctx, req)
}

func claimDraftPayload(values []model.ClaimDraft) []any {
	result := make([]any, 0, len(values))
	for _, value := range values {
		result = append(result, map[string]any{
			"subject_type": string(value.SubjectType),
			"subject_id":   value.SubjectID,
			"predicate":    value.Predicate,
			"object":       value.Object,
			"truth":        string(value.Truth),
		})
	}
	return result
}

func targetID(step model.ActionStep, entityType enum.EntityType) int64 {
	if step.Target == nil || step.Target.Type != entityType {
		return 0
	}
	return step.Target.ID
}

func int64Param(values map[string]any, key string) int64 {
	if values == nil {
		return 0
	}
	switch value := values[key].(type) {
	case int64:
		return value
	case int:
		return int64(value)
	case float64:
		return int64(value)
	case jsonNumber:
		return int64(value)
	default:
		return 0
	}
}

type jsonNumber float64

func stringParam(values map[string]any, key string) string {
	if values == nil {
		return ""
	}
	value, _ := values[key].(string)
	return strings.TrimSpace(value)
}

func boolPtrParam(values map[string]any, key string) *bool {
	if values == nil {
		return nil
	}
	value, ok := values[key].(bool)
	if !ok {
		return nil
	}
	return new(value)
}

func stringSliceParam(values map[string]any, key string) []string {
	if values == nil {
		return nil
	}
	raw, ok := values[key].([]any)
	if !ok {
		if typed, ok := values[key].([]string); ok {
			return typed
		}
		return nil
	}
	result := make([]string, 0, len(raw))
	for _, value := range raw {
		text, ok := value.(string)
		if ok && strings.TrimSpace(text) != "" {
			result = append(result, strings.TrimSpace(text))
		}
	}
	return result
}

func mapParam(values map[string]any, key string) map[string]any {
	if values == nil {
		return nil
	}
	value, _ := values[key].(map[string]any)
	return value
}

func floatMapParam(values map[string]any, key string) map[string]float64 {
	result := make(map[string]float64)
	if values == nil {
		return result
	}
	raw, ok := values[key].(map[string]any)
	if !ok {
		if typed, ok := values[key].(map[string]float64); ok {
			return typed
		}
		return result
	}
	for key, value := range raw {
		switch typed := value.(type) {
		case float64:
			result[key] = typed
		case int:
			result[key] = float64(typed)
		case int64:
			result[key] = float64(typed)
		}
	}
	return result
}

func stringValue[T ~string](value *T) string {
	if value == nil {
		return ""
	}
	return string(*value)
}

func boolPtrValue(value *bool) any {
	if value == nil {
		return nil
	}
	return *value
}

func nonEmptyString(value string, fallback string) string {
	value = strings.TrimSpace(value)
	if value != "" {
		return value
	}
	return strings.TrimSpace(fallback)
}
