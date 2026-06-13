package model

import "time"

// PushEvent 推送事件。
type PushEvent struct {
	ID        int64     `json:"id"`
	Type      int32     `json:"type"`
	UserID    int64     `json:"user_id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}
