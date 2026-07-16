package model

import "time"

type Memory struct {
	ID             int64
	WorldID        int64
	PlayerID       int64
	NpcID          int64
	Type           string
	Content        string
	Importance     int32
	SourceEventID  *int64
	LastRecalledAt *time.Time
	ExpiresAt      *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
