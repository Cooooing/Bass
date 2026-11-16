package event

import (
	event "common/proto/gen/common/enums/event"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
)

type ChatMessageHandler struct {
}

func NewChatMessageHandler() *ChatMessageHandler {
	return &ChatMessageHandler{}
}

func (h *ChatMessageHandler) Type() event.EventSubject {
	return event.EventSubject_EVENT_SUBJECT_GAME_IDLE_CHAT_MESSAGE
}

func (h *ChatMessageHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeChatMessageReceived,
		Payload:           req.Event.ChatMessage,
		TargetCharacterID: req.Event.ChatMessage.ReceiverCharacterID,
		Broadcast:         req.Event.ChatMessage.ReceiverCharacterID == 0,
	}, nil
}
