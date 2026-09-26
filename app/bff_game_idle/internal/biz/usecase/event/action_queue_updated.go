package event

import (
	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
	commonenum "common/pkg/enum"
	"context"
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
