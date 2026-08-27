package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type BackpackGetHandler struct {
	backpackUsecase *usecase.BackpackUsecase
}

func NewBackpackGetHandler(backpackUsecase *usecase.BackpackUsecase) *BackpackGetHandler {
	return &BackpackGetHandler{
		backpackUsecase: backpackUsecase,
	}
}

func (h *BackpackGetHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeBackpackGet
}

func (h *BackpackGetHandler) Validate(command *v1.WebSocketCommand) bool {
	return command.GetPayload().GetBackpackGet() != nil
}

func (h *BackpackGetHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Command.GetPayload().GetBackpackGet()
	items, err := h.backpackUsecase.Map(ctx, &usecase.BackpackMapReq{
		CharacterID: req.CharacterID,
		ItemIDs:     payload.GetItemIds(),
	})
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeBackpackItemsListed,
		Payload: items,
	})
	return nil
}
