package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

func (r *WorldAgentRunner) applyPlayerActionResult(ctx context.Context, result *agentResult) (*model.Event, error) {
	if result.resolution == nil {
		return nil, fmt.Errorf("player action resolution is nil")
	}

	status := strings.ToLower(strings.TrimSpace(result.resolution.Status))
	if status == "" {
		status = "resolved"
	}
	eventType := enum.EventTypeActionResolved
	if status == "rejected" {
		eventType = enum.EventTypeActionRejected
	}
	if status == "clarification" {
		eventType = enum.EventTypeActionClarificationRequired
	}
	if status == "resolved" {
		if err := r.applyActionSteps(ctx, result, result.resolution.Actions); err != nil {
			return nil, err
		}
	}

	worldSummary := strings.TrimSpace(result.resolution.WorldSummary)
	if worldSummary == "" {
		worldSummary = result.state.Summary
	}
	currentArc := r.normalizeModelText(result.resolution.CurrentArc, 256)
	if currentArc == "" {
		currentArc = result.state.CurrentArc
	}
	if _, err := r.worldStateRepo.UpdateNarrative(ctx, &repo.WorldStateUpdateNarrativeReq{
		WorldID:         result.world.ID,
		Version:         result.state.Version,
		Summary:         worldSummary,
		CurrentArc:      currentArc,
		PublicChronicle: result.state.PublicChronicle,
		CurrentEra:      result.state.CurrentEra,
	}); err != nil {
		if !r.isWorldVersionConflict(err) {
			return nil, err
		}
	}

	summary := r.normalizeModelText(result.resolution.Summary, 512)
	if summary == "" {
		summary = "玩家行动已处。"
	}
	content := summary
	if status == "clarification" && strings.TrimSpace(result.resolution.Clarification) != "" {
		content = strings.TrimSpace(result.resolution.Clarification)
	}

	event, err := r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             eventType,
		ActorPlayerID:    result.source.ActorPlayerID,
		NpcID:            result.source.NpcID,
		LocationID:       result.source.LocationID,
		CausationEventID: new(result.source.ID),
		Summary:          summary,
		Content:          content,
		Payload: map[string]any{
			"status":            status,
			"actions":           r.actionStepPayload(result.resolution.Actions),
			"claims":            r.claimDraftPayload(result.resolution.Claims),
			"suggested_actions": r.suggestedActionPayload(result.resolution.SuggestedActions),
		},
	})
	if err != nil {
		return nil, err
	}
	if err = r.saveClaims(ctx, event, result.npc, result.resolution.Claims); err != nil {
		return nil, err
	}
	if result.npc != nil && strings.TrimSpace(content) != "" {
		if err = r.saveNpcConversationMemory(ctx, result.npc, event, content); err != nil {
			return nil, err
		}
	}
	return event, nil
}

func (r *WorldAgentRunner) isWorldVersionConflict(err error) bool {
	return err != nil && strings.Contains(err.Error(), "world version conflict")
}

func (r *WorldAgentRunner) isNpcVersionConflict(err error) bool {
	return err != nil && strings.Contains(err.Error(), "npc version conflict")
}

func (r *WorldAgentRunner) applyNpcPlan(ctx context.Context, result *agentResult) (*model.Event, error) {
	if result.npc == nil || result.plan == nil {
		return nil, fmt.Errorf("npc plan context is incomplete")
	}

	minutes := result.plan.NextDecisionIn
	if minutes <= 0 {
		minutes = 24 * 60
	}
	next := result.state.WorldTime.Add(time.Duration(minutes) * time.Minute)
	updated, err := r.npcRepo.UpdateAutonomy(ctx, &repo.NpcAutonomyUpdateReq{
		NpcID:          result.npc.ID,
		Version:        result.npc.Version,
		Goal:           r.normalizeModelText(result.plan.Goal, 512),
		ContextSummary: r.normalizeModelText(result.plan.Summary, 512),
		NextDecisionAt: new(next),
		LastPlannedAt:  new(result.state.WorldTime),
	})
	if err != nil {
		if !r.isNpcVersionConflict(err) {
			return nil, err
		}
		updated, err = r.npcRepo.Get(ctx, &repo.NpcQuery{
			ID: new(result.npc.ID),
		})
		if err != nil {
			return nil, err
		}
	}
	result.npc = updated
	if err = r.applyActionSteps(ctx, result, result.plan.Actions); err != nil {
		return nil, err
	}
	refreshed, err := r.npcRepo.Get(ctx, &repo.NpcQuery{
		ID: new(updated.ID),
	})
	if err == nil {
		updated = refreshed
	}

	payload := map[string]any{
		"goal":    updated.Goal,
		"actions": r.actionStepPayload(result.plan.Actions),
	}
	return r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
		WorldID:          result.world.ID,
		Type:             enum.EventTypeNpcPlanned,
		NpcID:            new(updated.ID),
		LocationID:       new(updated.CurrentLocationID),
		CausationEventID: new(result.source.ID),
		Summary:          updated.Name + " 制定了新的计。",
		Content:          r.normalizeModelText(result.plan.Summary, 512),
		Payload:          payload,
	})
}

func (r *WorldAgentRunner) suggestedActionPayload(values []model.SuggestedAction) []any {
	result := make([]any, 0, len(values))
	for _, value := range values {
		targets := make([]any, 0, len(value.Targets))
		for _, target := range value.Targets {
			targets = append(targets, map[string]any{"type": target.Type.String(), "id": target.ID})
		}
		result = append(result, map[string]any{"label": value.Label, "content": value.Content, "targets": targets})
	}
	return result
}

func (r *WorldAgentRunner) actionStepPayload(values []model.ActionStep) []any {
	result := make([]any, 0, len(values))
	for _, value := range values {
		item := map[string]any{
			"type":             value.Type,
			"parameters":       value.Parameters,
			"duration_minutes": value.DurationMinutes,
		}
		if value.Target != nil {
			item["target"] = map[string]any{"type": value.Target.Type.String(), "id": value.Target.ID}
		}
		result = append(result, item)
	}
	return result
}
