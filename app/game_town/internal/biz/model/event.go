package model

import (
	"time"

	"game_town/internal/enum"
)

type Event struct {
	ID               int64
	WorldID          int64
	Sequence         uint64
	Type             enum.EventType
	ActorPlayerID    *int64
	NpcID            *int64
	LocationID       *int64
	CausationEventID *int64
	Summary          string
	Content          string
	Payload          map[string]any
	WorldTime        time.Time
	OccurredAt       time.Time
	CreatedAt        time.Time
}

type EntityRef struct {
	Type enum.EntityType `json:"type"`
	ID   int64           `json:"id"`
}

type ActionStep struct {
	Type            string         `json:"type"`
	Target          *EntityRef     `json:"target"`
	Parameters      map[string]any `json:"parameters"`
	DurationMinutes int64          `json:"duration_minutes"`
}

type ActionResolution struct {
	Status           string            `json:"status"`
	Summary          string            `json:"summary"`
	Clarification    string            `json:"clarification"`
	WorldSummary     string            `json:"world_summary"`
	CurrentArc       string            `json:"current_arc"`
	Actions          []ActionStep      `json:"actions"`
	Claims           []ClaimDraft      `json:"claims"`
	SuggestedActions []SuggestedAction `json:"suggested_actions"`
}
