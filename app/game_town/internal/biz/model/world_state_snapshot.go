package model

import "time"

type WorldStateSnapshot struct {
	ID            int64
	WorldID       int64
	TickCount     int64
	CurrentArc    string
	Metrics       map[string]any
	Summary       string
	ReasonEventID *int64
	CreatedAt     time.Time
}
