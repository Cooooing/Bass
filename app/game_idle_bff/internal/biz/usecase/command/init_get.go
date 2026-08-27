package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
	"time"
)

type InitGetHandler struct {
	actionQueueUsecase      *usecase.ActionQueueUsecase
	characterAbilityUsecase *usecase.CharacterAbilityUsecase
	backpackUsecase         *usecase.BackpackUsecase
	chatUsecase             *usecase.ChatUsecase
}

func NewInitGetHandler(
	actionQueueUsecase *usecase.ActionQueueUsecase,
	characterAbilityUsecase *usecase.CharacterAbilityUsecase,
	backpackUsecase *usecase.BackpackUsecase,
	chatUsecase *usecase.ChatUsecase,
) *InitGetHandler {
	return &InitGetHandler{
		actionQueueUsecase:      actionQueueUsecase,
		characterAbilityUsecase: characterAbilityUsecase,
		backpackUsecase:         backpackUsecase,
		chatUsecase:             chatUsecase,
	}
}

func (h *InitGetHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeInitGet
}

func (h *InitGetHandler) Validate(command *v1.WebSocketCommand) bool {
	return true
}

func (h *InitGetHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	actionQueue, err := h.actionQueueUsecase.List(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	abilities, err := h.characterAbilityUsecase.Map(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	backpackItems, err := h.backpackUsecase.Map(ctx, &usecase.BackpackMapReq{
		CharacterID: req.CharacterID,
	})
	if err != nil {
		return err
	}
	chatMessages, err := h.chatUsecase.List(ctx, &usecase.ListChatMessagesReq{
		ChannelType: "world",
		ChannelID:   "world",
		Size:        50,
	})
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type: enum.WebSocketMessageTypeInitCompleted,
		Payload: &model.WebSocketInit{
			ActionQueue:   actionQueue,
			Abilities:     abilities,
			BackpackItems: backpackItems,
			ChatMessages:  chatMessages,
			ServerTime:    time.Now().Unix(),
		},
	})
	return nil
}
