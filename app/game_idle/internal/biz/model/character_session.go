package model

import (
	"game_idle/internal/enum"
	"time"
)

type CharacterSession struct {
	CharacterID int64
	SessionID   string
	ExpiresIn   time.Duration
}

type CharacterCloseSessionEvent struct {
	SessionID       string
	Reason          enum.CharacterCloseSessionReason
	Message         string
	ShouldReconnect bool
}
