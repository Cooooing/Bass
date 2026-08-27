package model

import (
	commonenum "common/pkg/enum"
	"time"
)

type WebSocketSession struct {
	CharacterID       int64
	SessionID         string
	RemainingDuration time.Duration
}

type WebSocketChatMessage struct {
	MessageID           int64  `json:"message_id"`
	ChannelType         string `json:"channel_type"`
	ChannelID           string `json:"channel_id"`
	SenderCharacterID   int64  `json:"sender_character_id"`
	ReceiverCharacterID int64  `json:"receiver_character_id,omitempty"`
	Content             string `json:"content"`
}

type WebSocketEvent struct {
	Type             commonenum.EventType
	ChatMessage      *WebSocketChatMessage
	CloseSession     *WebSocketCloseSession
	ActionCompleted  *WebSocketActionCompleted
	AbilityLeveledUp *WebSocketAbilityLeveledUp
}

type WebSocketCloseSession struct {
	SessionID       string `json:"session_id"`
	Reason          string `json:"reason"`
	Message         string `json:"message,omitempty"`
	ShouldReconnect bool   `json:"should_reconnect"`
}

type WebSocketActionCompleted struct {
	CharacterID    int64                     `json:"character_id"`
	Action         *WebSocketCompletedAction `json:"action"`
	ItemChanges    []*WebSocketItemChange    `json:"item_changes"`
	AbilityChanges []*WebSocketAbilityChange `json:"ability_changes"`
}

type WebSocketCompletedAction struct {
	ActionID       string     `json:"action_id"`
	TimesFinished  int64      `json:"times_finished"`
	TimesRemaining int64      `json:"times_remaining"`
	StartedAt      *time.Time `json:"started_at,omitempty"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
}

type WebSocketItemChange struct {
	ItemID        string `json:"item_id"`
	QuantityDelta int64  `json:"quantity_delta"`
	QuantityAfter int64  `json:"quantity_after"`
}

type WebSocketAbilityChange struct {
	AbilityID string `json:"ability_id"`
	ExpDelta  int64  `json:"exp_delta"`
	ExpAfter  int64  `json:"exp_after"`
}

type WebSocketAbilityLeveledUp struct {
	CharacterID  int64  `json:"character_id"`
	AbilityID    string `json:"ability_id"`
	Level        int32  `json:"level"`
	Exp          int64  `json:"exp"`
	NextLevelExp int64  `json:"next_level_exp"`
}
