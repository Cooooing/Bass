package model

import "time"

type Relationship struct {
	ID                int64
	WorldID           int64
	PlayerID          int64
	NpcID             int64
	Affinity          int32
	Trust             int32
	Tension           int32
	CustomMetrics     map[string]any
	LastInteractionAt *time.Time
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
