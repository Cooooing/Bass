package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ActionRemoveHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionRemoveHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionRemoveHandler {
	return &ActionRemoveHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *ActionRemoveHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionRemove
}

func (h *ActionRemoveHandler) Validate(command *v1.WebSocketCommand) bool {
	return command.GetPayload().GetActionRemove() != nil
}

func (h *ActionRemoveHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Command.GetPayload().GetActionRemove()
	return h.actionQueueUsecase.Remove(ctx, &usecase.RemoveActionReq{
		CharacterID: req.CharacterID,
		Position:    payload.GetPosition(),
	})
}
