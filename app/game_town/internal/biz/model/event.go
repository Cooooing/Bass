package model

import "time"

type Event struct {
	ID            int64
	WorldID       int64
	Type          string
	ActorPlayerID *int64
	TargetNpcID   *int64
	LocationID    *int64
	CommandID     *int64
	Summary       string
	Content       string
	Effects       map[string]any
	Metadata      map[string]any
	OccurredAt    time.Time
	CreatedAt     time.Time
}
