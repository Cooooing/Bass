package command

import (
	"common/pkg/apperror"
	v1 "common/proto/gen/bff_game_idle/v1"
	cerrors "common/proto/gen/common/errors"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"

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
	payload, ok := req.Payload.(*v1.ActionMoveWebSocketPayload)
	if !ok {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return h.actionQueueUsecase.Move(ctx, &usecase.MoveActionReq{
		CharacterID:     req.CharacterID,
		CurrentPosition: payload.GetCurrentPosition(),
		TargetPosition:  payload.GetTargetPosition(),
	})
}
