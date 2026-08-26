package command

import (
	"context"
	"encoding/json"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type QueueListHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewQueueListHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *QueueListHandler {
	return &QueueListHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *QueueListHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeQueueList
}

func (h *QueueListHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	queue, err := h.actionQueueUsecase.List(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeQueueList,
		Payload: queue,
	})
	return nil
}

type QueueAddHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewQueueAddHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *QueueAddHandler {
	return &QueueAddHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *QueueAddHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeQueueAdd
}

func (h *QueueAddHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := struct {
		ActionID string `json:"action_id"`
		Times    int64  `json:"times"`
		Position *int32 `json:"position,omitempty"`
	}{}
	if err := json.Unmarshal(req.Command.Payload, &payload); err != nil {
		return err
	}
	return h.actionQueueUsecase.Add(ctx, &usecase.AddActionReq{
		CharacterID: req.CharacterID,
		ActionID:    payload.ActionID,
		Times:       payload.Times,
		Position:    payload.Position,
	})
}

type QueueMoveHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewQueueMoveHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *QueueMoveHandler {
	return &QueueMoveHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *QueueMoveHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeQueueMove
}

func (h *QueueMoveHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := struct {
		CurrentPosition int32 `json:"current_position"`
		TargetPosition  int32 `json:"target_position"`
	}{}
	if err := json.Unmarshal(req.Command.Payload, &payload); err != nil {
		return err
	}
	return h.actionQueueUsecase.Move(ctx, &usecase.MoveActionReq{
		CharacterID:     req.CharacterID,
		CurrentPosition: payload.CurrentPosition,
		TargetPosition:  payload.TargetPosition,
	})
}

type QueueRemoveHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewQueueRemoveHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *QueueRemoveHandler {
	return &QueueRemoveHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *QueueRemoveHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeQueueRemove
}

func (h *QueueRemoveHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := struct {
		Position int32 `json:"position"`
	}{}
	if err := json.Unmarshal(req.Command.Payload, &payload); err != nil {
		return err
	}
	return h.actionQueueUsecase.Remove(ctx, &usecase.RemoveActionReq{
		CharacterID: req.CharacterID,
		Position:    payload.Position,
	})
}

type QueueClearHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewQueueClearHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *QueueClearHandler {
	return &QueueClearHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *QueueClearHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeQueueClear
}

func (h *QueueClearHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	return h.actionQueueUsecase.Clear(ctx, req.CharacterID)
}
