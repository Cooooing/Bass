package command

import (
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"

	"google.golang.org/protobuf/proto"
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

func (h *ActionClearHandler) Payload() proto.Message {
	return nil
}

func (h *ActionClearHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	return h.actionQueueUsecase.Clear(ctx, req.CharacterID)
}
