package model

import (
	"time"

	"game_town/internal/enum"
)

type Npc struct {
	ID                int64
	WorldID           int64
	Code              string
	Name              string
	Role              string
	Species           string
	Personality       string
	Goal              string
	Background        string
	CurrentLocationID int64
	SystemPrompt      string
	ContextSummary    string
	LifeStatus        enum.NpcLifeStatus
	StateTags         []string
	Attributes        map[string]any
	BirthWorldTime    *time.Time
	DeathWorldTime    *time.Time
	NextDecisionAt    *time.Time
	LastPlannedAt     *time.Time
	Version           int64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

type SuggestedAction struct {
	Label   string      `json:"label"`
	Content string      `json:"content"`
	Targets []EntityRef `json:"targets"`
}

type NpcReply struct {
	Reply            string            `json:"reply"`
	ContextSummary   string            `json:"context_summary"`
	SuggestedActions []SuggestedAction `json:"suggested_actions"`
	Claims           []ClaimDraft      `json:"claims"`
}

type NpcPlan struct {
	Summary        string       `json:"summary"`
	Goal           string       `json:"goal"`
	NextDecisionIn int64        `json:"next_decision_minutes"`
	Actions        []ActionStep `json:"actions"`
}
