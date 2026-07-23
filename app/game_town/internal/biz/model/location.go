package model

import (
	"time"

	"game_town/internal/enum"
)

type Location struct {
	ID                   int64
	WorldID              int64
	Code                 string
	Name                 string
	Description          string
	Status               enum.LocationStatus
	ControllingFactionID *int64
	EnvironmentTags      []string
	Attributes           map[string]any
	Accessible           bool
	Version              int64
	Sort                 int32
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
}
