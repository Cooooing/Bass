package repo

import (
	"game_town/internal/biz/model"
	"game_town/internal/data/gen"
)

func (r *baseRepo) world(row *gen.World) *model.World {
	return &model.World{ID: row.ID, Code: row.Code, Name: row.Name, Description: row.Description, Scale: row.Scale, Status: row.Status, CreatorPlayerID: row.CreatorPlayerID, DefaultLocationID: row.DefaultLocationID, AgentConfigID: row.AgentConfigID, Seed: row.Seed, GenerationParams: row.GenerationParams, GenerationSummary: row.GenerationSummary, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) player(row *gen.Player) *model.Player {
	return &model.Player{ID: row.ID, Name: row.Name, DisplayName: row.DisplayName, Status: row.Status, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) agentConfig(row *gen.AgentConfig) *model.AgentConfig {
	return &model.AgentConfig{ID: row.ID, PlayerID: row.PlayerID, Name: row.Name, Provider: row.Provider, Model: row.Model, BaseURL: row.BaseURL, APIKey: row.APIKey, TimeoutSeconds: row.TimeoutSeconds, IsDefault: row.IsDefault, Status: row.Status, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) member(row *gen.WorldMember) *model.WorldMember {
	return &model.WorldMember{ID: row.ID, WorldID: row.WorldID, PlayerID: row.PlayerID, CurrentLocationID: row.CurrentLocationID, Role: row.Role, JoinedAt: row.JoinedAt, LastSeenAt: row.LastSeenAt, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) location(row *gen.Location) *model.Location {
	return &model.Location{ID: row.ID, WorldID: row.WorldID, Code: row.Code, Name: row.Name, Description: row.Description, Tags: row.Tags, Sort: row.Sort, Enabled: row.Enabled, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) npc(row *gen.Npc) *model.Npc {
	return &model.Npc{ID: row.ID, WorldID: row.WorldID, Code: row.Code, Name: row.Name, Role: row.Role, Personality: row.Personality, Goal: row.Goal, Background: row.Background, CurrentLocationID: row.CurrentLocationID, State: row.State, SystemPrompt: row.SystemPrompt, GeneratedProfile: row.GeneratedProfile, Enabled: row.Enabled, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) relationship(row *gen.Relationship) *model.Relationship {
	return &model.Relationship{ID: row.ID, WorldID: row.WorldID, PlayerID: row.PlayerID, NpcID: row.NpcID, Affinity: row.Affinity, Trust: row.Trust, Tension: row.Tension, CustomMetrics: row.CustomMetrics, LastInteractionAt: row.LastInteractionAt, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) memory(row *gen.Memory) *model.Memory {
	return &model.Memory{ID: row.ID, WorldID: row.WorldID, PlayerID: row.PlayerID, NpcID: row.NpcID, Type: row.Type, Content: row.Content, Importance: row.Importance, SourceEventID: row.SourceEventID, LastRecalledAt: row.LastRecalledAt, ExpiresAt: row.ExpiresAt, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) metric(row *gen.WorldMetricDefinition) *model.WorldMetricDefinition {
	return &model.WorldMetricDefinition{ID: row.ID, WorldID: row.WorldID, Key: row.Key, Name: row.Name, Description: row.Description, MinValue: row.MinValue, MaxValue: row.MaxValue, InitialValue: row.InitialValue, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) state(row *gen.WorldStateSnapshot) *model.WorldStateSnapshot {
	return &model.WorldStateSnapshot{ID: row.ID, WorldID: row.WorldID, TickCount: row.TickCount, CurrentArc: row.CurrentArc, Metrics: row.Metrics, Summary: row.Summary, ReasonEventID: row.ReasonEventID, CreatedAt: row.CreatedAt}
}

func (r *baseRepo) event(row *gen.Event) *model.Event {
	return &model.Event{ID: row.ID, WorldID: row.WorldID, Type: row.Type, ActorPlayerID: row.ActorPlayerID, TargetNpcID: row.TargetNpcID, LocationID: row.LocationID, CommandID: row.CommandID, Summary: row.Summary, Content: row.Content, Effects: row.Effects, Metadata: row.Metadata, OccurredAt: row.OccurredAt, CreatedAt: row.CreatedAt}
}

func (r *baseRepo) command(row *gen.Command) *model.Command {
	return &model.Command{ID: row.ID, WorldID: row.WorldID, SessionID: row.SessionID, PlayerID: row.PlayerID, RawText: row.RawText, Type: row.Type, ParsedPayload: row.ParsedPayload, Status: row.Status, ErrorCode: row.ErrorCode, ResultSummary: row.ResultSummary, CreatedAt: row.CreatedAt, HandledAt: row.HandledAt}
}

func (r *baseRepo) session(row *gen.Session) *model.Session {
	return &model.Session{ID: row.ID, PlayerID: row.PlayerID, CurrentWorldID: row.CurrentWorldID, ClientType: row.ClientType, StartedAt: row.StartedAt, LastSeenAt: row.LastSeenAt, EndedAt: row.EndedAt, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func (r *baseRepo) agentRun(row *gen.AgentRun) *model.AgentRun {
	return &model.AgentRun{ID: row.ID, WorldID: row.WorldID, AgentConfigID: row.AgentConfigID, RunType: row.RunType, CommandID: row.CommandID, EventID: row.EventID, NpcID: row.NpcID, Model: row.Model, InputJSON: row.InputJSON, OutputJSON: row.OutputJSON, Status: row.Status, ErrorSummary: row.ErrorSummary, LatencyMS: row.LatencyMs, CreatedAt: row.CreatedAt}
}

func optionalNumber(v *int64) any {
	if v == nil {
		return nil
	}
	return float64(*v)
}
