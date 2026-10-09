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

func (h *ActionRemoveHandler) Payload() proto.Message {
	return &v1.ActionRemoveWebSocketPayload{}
}

func (h *ActionRemoveHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload, ok := req.Payload.(*v1.ActionRemoveWebSocketPayload)
	if !ok {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return h.actionQueueUsecase.Remove(ctx, &usecase.RemoveActionReq{
		CharacterID: req.CharacterID,
		Position:    payload.GetPosition(),
	})
}
