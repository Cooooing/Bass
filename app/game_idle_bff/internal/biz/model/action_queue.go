package model

import "time"

type ActionQueue struct {
	CharacterID int64
	Items       []*ActionQueueItem
}

type ActionQueueItem struct {
	ActionID  string
	Times     int64
	CreatedAt time.Time
}
