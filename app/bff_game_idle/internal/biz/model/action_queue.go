package model

import "time"

type ActionQueue struct {
	CharacterID int64              `json:"character_id"`
	Items       []*ActionQueueItem `json:"items"`
}

type ActionQueueItem struct {
	ActionID  string    `json:"action_id"`
	Times     int64     `json:"times"`
	CreatedAt time.Time `json:"created_at"`
}
