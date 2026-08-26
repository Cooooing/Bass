package command

import (
	"context"
	"encoding/json"
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

func (h *BackpackGetHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := struct {
		ItemIDs []string `json:"item_ids,omitempty"`
	}{}
	if len(req.Command.Payload) > 0 {
		if err := json.Unmarshal(req.Command.Payload, &payload); err != nil {
			return err
		}
	}
	items, err := h.backpackUsecase.Map(ctx, &usecase.BackpackMapReq{
		CharacterID: req.CharacterID,
		ItemIDs:     payload.ItemIDs,
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
