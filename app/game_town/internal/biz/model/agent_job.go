package model

import (
	"time"

	"game_town/internal/enum"
)

type AgentJob struct {
	ID            int64
	WorldID       int64
	SourceEventID int64
	Type          enum.AgentJobType
	Priority      enum.AgentJobPriority
	LaneKey       string
	Status        enum.AgentJobStatus
	WorldVersion  int64
	NpcID         *int64
	AttemptCount  int32
	AvailableAt   time.Time
	StartedAt     *time.Time
	FinishedAt    *time.Time
	ErrorSummary  string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
