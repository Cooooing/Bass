package model

import (
	"time"

	"game_town/internal/enum"
)

type AgentConfig struct {
	ID             int64
	Name           string
	Provider       enum.AgentProvider
	BaseURL        string
	Model          string
	SecretEnv      string
	TimeoutSeconds int32
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
