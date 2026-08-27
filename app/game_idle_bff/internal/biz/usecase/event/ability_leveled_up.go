package event

import (
	commonenum "common/pkg/enum"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type AbilityLeveledUpHandler struct {
}

func NewAbilityLeveledUpHandler() *AbilityLeveledUpHandler {
	return &AbilityLeveledUpHandler{}
}

func (h *AbilityLeveledUpHandler) Type() commonenum.EventType {
	return commonenum.EventTypeGameIdleAbilityLeveledUp
}

func (h *AbilityLeveledUpHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeAbilityLeveledUp,
		Payload:           req.Event.AbilityLeveledUp,
		TargetCharacterID: req.Event.AbilityLeveledUp.CharacterID,
	}, nil
}
