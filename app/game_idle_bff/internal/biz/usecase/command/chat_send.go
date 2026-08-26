package command

import (
	"context"
	"encoding/json"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ChatSendHandler struct {
	chatUsecase *usecase.ChatUsecase
}

func NewChatSendHandler(chatUsecase *usecase.ChatUsecase) *ChatSendHandler {
	return &ChatSendHandler{
		chatUsecase: chatUsecase,
	}
}

func (h *ChatSendHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeChatSend
}

func (h *ChatSendHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := struct {
		ChannelType         string `json:"channel_type"`
		ChannelID           string `json:"channel_id"`
		ReceiverCharacterID int64  `json:"receiver_character_id,omitempty"`
		Content             string `json:"content"`
	}{}
	if err := json.Unmarshal(req.Command.Payload, &payload); err != nil {
		return err
	}
	_, err := h.chatUsecase.Send(ctx, &usecase.SendChatMessageReq{
		CharacterID:         req.CharacterID,
		ChannelType:         payload.ChannelType,
		ChannelID:           payload.ChannelID,
		ReceiverCharacterID: payload.ReceiverCharacterID,
		Content:             payload.Content,
	})
	return err
}

type ChatListHandler struct {
	chatUsecase *usecase.ChatUsecase
}

func NewChatListHandler(chatUsecase *usecase.ChatUsecase) *ChatListHandler {
	return &ChatListHandler{
		chatUsecase: chatUsecase,
	}
}

func (h *ChatListHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeChatList
}

func (h *ChatListHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := struct {
		ChannelType string `json:"channel_type"`
		ChannelID   string `json:"channel_id"`
		BeforeID    int64  `json:"before_id,omitempty"`
		Size        int32  `json:"size,omitempty"`
	}{}
	if err := json.Unmarshal(req.Command.Payload, &payload); err != nil {
		return err
	}
	rows, err := h.chatUsecase.List(ctx, &usecase.ListChatMessagesReq{
		ChannelType: payload.ChannelType,
		ChannelID:   payload.ChannelID,
		BeforeID:    payload.BeforeID,
		Size:        payload.Size,
	})
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeChatList,
		Payload: rows,
	})
	return nil
}
