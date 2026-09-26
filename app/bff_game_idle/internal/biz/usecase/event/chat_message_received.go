package event

import (
	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
	commonenum "common/pkg/enum"
	"context"
)

type ChatMessageHandler struct {
}

func NewChatMessageHandler() *ChatMessageHandler {
	return &ChatMessageHandler{}
}

func (h *ChatMessageHandler) Type() commonenum.EventType {
	return commonenum.EventTypeGameIdleChatMessage
}

func (h *ChatMessageHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeChatMessageReceived,
		Payload:           req.Event.ChatMessage,
		TargetCharacterID: req.Event.ChatMessage.ReceiverCharacterID,
		Broadcast:         req.Event.ChatMessage.ReceiverCharacterID == 0,
	}, nil
}
