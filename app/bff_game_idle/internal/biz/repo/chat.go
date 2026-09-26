package repo

import (
	"bff_game_idle/internal/biz/model"
	"context"
)

type ChatRepo interface {
	Send(ctx context.Context, req *SendChatMessageReq) (*model.WebSocketChatMessage, error)
	List(ctx context.Context, req *ListChatMessagesReq) ([]*model.WebSocketChatMessage, error)
}

type SendChatMessageReq struct {
	CharacterID         int64
	ChannelType         string
	ChannelID           string
	ReceiverCharacterID int64
	Content             string
}

type ListChatMessagesReq struct {
	ChannelType string
	ChannelID   string
	BeforeID    int64
	Size        int32
}
