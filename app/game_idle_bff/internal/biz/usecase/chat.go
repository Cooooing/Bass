package usecase

import (
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

type ChatUsecase struct {
	chatRepo repo.ChatRepo
}

func NewChatUsecase(chatRepo repo.ChatRepo) *ChatUsecase {
	return &ChatUsecase{
		chatRepo: chatRepo,
	}
}

type SendChatMessageReq struct {
	CharacterID         int64
	ChannelType         string
	ChannelID           string
	ReceiverCharacterID int64
	Content             string
}

func (u *ChatUsecase) Send(ctx context.Context, req *SendChatMessageReq) (*model.WebSocketChatMessage, error) {
	return u.chatRepo.Send(ctx, &repo.SendChatMessageReq{
		CharacterID:         req.CharacterID,
		ChannelType:         req.ChannelType,
		ChannelID:           req.ChannelID,
		ReceiverCharacterID: req.ReceiverCharacterID,
		Content:             req.Content,
	})
}

type ListChatMessagesReq struct {
	ChannelType string
	ChannelID   string
	BeforeID    int64
	Size        int32
}

func (u *ChatUsecase) List(ctx context.Context, req *ListChatMessagesReq) ([]*model.WebSocketChatMessage, error) {
	return u.chatRepo.List(ctx, &repo.ListChatMessagesReq{
		ChannelType: req.ChannelType,
		ChannelID:   req.ChannelID,
		BeforeID:    req.BeforeID,
		Size:        req.Size,
	})
}
