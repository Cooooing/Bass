package model

import "time"

type Command struct {
	ID            int64
	WorldID       *int64
	SessionID     int64
	PlayerID      *int64
	RawText       string
	Type          string
	ParsedPayload map[string]any
	Status        string
	ErrorCode     *int32
	ResultSummary string
	CreatedAt     time.Time
	HandledAt     *time.Time
}
