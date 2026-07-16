package model

import "time"

type AgentConfig struct {
	ID             int64
	PlayerID       int64
	Name           string
	Provider       string
	Model          string
	BaseURL        string
	APIKey         string
	TimeoutSeconds int32
	IsDefault      bool
	Status         string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
