package model

import "time"

type Npc struct {
	ID                int64
	WorldID           int64
	Code              string
	Name              string
	Role              string
	Personality       string
	Goal              string
	Background        string
	CurrentLocationID int64
	State             string
	SystemPrompt      string
	GeneratedProfile  map[string]any
	Enabled           bool
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
