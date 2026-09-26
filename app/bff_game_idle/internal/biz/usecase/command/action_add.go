package command

import (
	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
	v1 "common/proto/gen/bff_game_idle/v1"
	"context"

	"google.golang.org/protobuf/proto"
)

type ActionAddHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionAddHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionAddHandler {
	return &ActionAddHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *ActionAddHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionAdd
}

func (h *ActionAddHandler) Payload() proto.Message {
	return &v1.ActionAddWebSocketPayload{}
}

func (h *ActionAddHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Payload.(*v1.ActionAddWebSocketPayload)
	return h.actionQueueUsecase.Add(ctx, &usecase.AddActionReq{
		CharacterID: req.CharacterID,
		ActionID:    payload.GetActionId(),
		Times:       payload.GetTimes(),
		Position:    payload.Position,
	})
}
