package event

import (
	commonenum "common/pkg/enum"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
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
