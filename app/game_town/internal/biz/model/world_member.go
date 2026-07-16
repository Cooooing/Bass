package model

import "time"

type WorldMember struct {
	ID                int64
	WorldID           int64
	PlayerID          int64
	CurrentLocationID int64
	Role              string
	JoinedAt          time.Time
	LastSeenAt        time.Time
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
