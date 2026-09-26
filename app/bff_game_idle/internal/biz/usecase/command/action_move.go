package command

import (
	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
	v1 "common/proto/gen/bff_game_idle/v1"
	"context"

	"google.golang.org/protobuf/proto"
)

type ActionMoveHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionMoveHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionMoveHandler {
	return &ActionMoveHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *ActionMoveHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionMove
}

func (h *ActionMoveHandler) Payload() proto.Message {
	return &v1.ActionMoveWebSocketPayload{}
}

func (h *ActionMoveHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Payload.(*v1.ActionMoveWebSocketPayload)
	return h.actionQueueUsecase.Move(ctx, &usecase.MoveActionReq{
		CharacterID:     req.CharacterID,
		CurrentPosition: payload.GetCurrentPosition(),
		TargetPosition:  payload.GetTargetPosition(),
	})
}
