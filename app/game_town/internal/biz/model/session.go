package model

import "time"

type Session struct {
	ID             int64
	PlayerID       *int64
	CurrentWorldID *int64
	ClientType     string
	StartedAt      time.Time
	LastSeenAt     time.Time
	EndedAt        *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
