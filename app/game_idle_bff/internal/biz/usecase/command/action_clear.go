package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ActionClearHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionClearHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionClearHandler {
	return &ActionClearHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *ActionClearHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionClear
}

func (h *ActionClearHandler) Validate(command *v1.WebSocketCommand) bool {
	return true
}

func (h *ActionClearHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	return h.actionQueueUsecase.Clear(ctx, req.CharacterID)
}
