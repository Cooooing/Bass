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
	payload, ok := req.Payload.(*v1.ActionAddWebSocketPayload)
	if !ok {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return h.actionQueueUsecase.Add(ctx, &usecase.AddActionReq{
		CharacterID: req.CharacterID,
		ActionID:    payload.GetActionId(),
		Times:       payload.GetTimes(),
		Position:    payload.Position,
	})
}
