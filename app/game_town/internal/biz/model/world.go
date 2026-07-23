package model

import (
	"time"

	"game_town/internal/enum"
)

type World struct {
	ID                int64
	Code              string
	Name              string
	Description       string
	Status            enum.WorldStatus
	CreatorPlayerID   int64
	DefaultLocationID *int64
	AgentConfigID     int64
	Seed              int64
	GenerationSummary string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

type WorldState struct {
	ID              int64
	WorldID         int64
	Version         int64
	EventSequence   uint64
	AgentCursor     uint64
	Summary         string
	CurrentArc      string
	PublicChronicle string
	CurrentEra      string
	WorldTime       time.Time
	TimeAnchor      time.Time
	TimeScale       uint32
	RuleVersion     int64
	NextTickAt      *time.Time
	NextDueAt       *time.Time
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type WorldRule struct {
	ID        int64
	WorldID   int64
	Version   int64
	Rules     map[string]any
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type WorldDraft struct {
	Name       string               `json:"name"`
	Summary    string               `json:"summary"`
	CurrentArc string               `json:"current_arc"`
	CurrentEra string               `json:"current_era"`
	Rules      map[string]any       `json:"rules"`
	Locations  []WorldDraftLocation `json:"locations"`
	Factions   []WorldDraftFaction  `json:"factions"`
	Npcs       []WorldDraftNpc      `json:"npcs"`
}

type WorldDraftLocation struct {
	Code            string   `json:"code"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	EnvironmentTags []string `json:"environment_tags"`
}

type WorldDraftFaction struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PublicGoal  string `json:"public_goal"`
}

type WorldDraftNpc struct {
	Code         string         `json:"code"`
	Name         string         `json:"name"`
	Role         string         `json:"role"`
	Species      string         `json:"species"`
	Personality  string         `json:"personality"`
	Goal         string         `json:"goal"`
	Background   string         `json:"background"`
	LocationCode string         `json:"location_code"`
	FactionCode  string         `json:"faction_code"`
	SystemPrompt string         `json:"system_prompt"`
	Attributes   map[string]any `json:"attributes"`
}
