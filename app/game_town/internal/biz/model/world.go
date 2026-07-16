package model

import "time"

type World struct {
	ID                int64
	Code              string
	Name              string
	Description       string
	Scale             string
	Status            string
	CreatorPlayerID   int64
	DefaultLocationID *int64
	AgentConfigID     *int64
	Seed              int64
	GenerationParams  map[string]any
	GenerationSummary string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
