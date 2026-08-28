package event

import (
	commonenum "common/pkg/enum"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ActionQueueUpdatedHandler struct {
}

func NewActionQueueUpdatedHandler() *ActionQueueUpdatedHandler {
	return &ActionQueueUpdatedHandler{}
}

func (h *ActionQueueUpdatedHandler) Type() commonenum.EventType {
	return commonenum.EventTypeGameIdleActionQueueUpdated
}

func (h *ActionQueueUpdatedHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeActionQueueUpdated,
		Payload:           req.Event.ActionQueueUpdated,
		TargetCharacterID: req.Event.ActionQueueUpdated.CharacterID,
	}, nil
}
