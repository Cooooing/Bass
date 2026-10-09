package event

import (
	commonenum "common/pkg/enum"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
)

type ActionCompletedHandler struct {
}

func NewActionCompletedHandler() *ActionCompletedHandler {
	return &ActionCompletedHandler{}
}

func (h *ActionCompletedHandler) Type() commonenum.EventType {
	return commonenum.EventTypeGameIdleActionCompleted
}

func (h *ActionCompletedHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeActionCompleted,
		Payload:           req.Event.ActionCompleted,
		TargetCharacterID: req.Event.ActionCompleted.CharacterID,
	}, nil
}
