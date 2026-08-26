package event

import (
	commonenum "common/pkg/enum"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
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
		Type:              enum.WebSocketMessageTypeChatMessage,
		Payload:           req.Event.ChatMessage,
		TargetCharacterID: req.Event.ChatMessage.ReceiverCharacterID,
		Broadcast:         req.Event.ChatMessage.ReceiverCharacterID == 0,
	}, nil
}
