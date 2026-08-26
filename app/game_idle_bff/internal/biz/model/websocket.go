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
	Type         commonenum.EventType
	ChatMessage  *WebSocketChatMessage
	CloseSession *WebSocketCloseSession
}

type WebSocketCloseSession struct {
	SessionID       string `json:"session_id"`
	Reason          string `json:"reason"`
	Message         string `json:"message,omitempty"`
	ShouldReconnect bool   `json:"should_reconnect"`
}
