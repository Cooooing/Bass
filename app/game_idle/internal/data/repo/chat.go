package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	chatmessageent "game_idle/internal/data/gen/chatmessage"
	"game_idle/internal/enum"
)

var _ bizrepo.ChatMessageRepo = (*ChatMessageRepo)(nil)

type ChatMessageRepo struct {
	db *gen.Client
}

func NewChatMessageRepo(db *gen.Client) bizrepo.ChatMessageRepo {
	return &ChatMessageRepo{db: db}
}

func (r *ChatMessageRepo) Create(ctx context.Context, message *model.ChatMessage) (*model.ChatMessage, error) {
	create := r.db.ChatMessage.Create().
		SetChannelType(chatmessageent.ChannelType(message.ChannelType)).
		SetChannelID(message.ChannelID).
		SetSenderCharacterID(message.SenderCharacterID).
		SetContent(message.Content).
		SetStatus(chatmessageent.Status(message.Status))
	if message.ReceiverCharacterID != nil {
		create.SetReceiverCharacterID(*message.ReceiverCharacterID)
	}
	row, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatMessage{
		ID:                  row.ID,
		ChannelType:         enum.ChatChannelType(row.ChannelType),
		ChannelID:           row.ChannelID,
		SenderCharacterID:   row.SenderCharacterID,
		ReceiverCharacterID: row.ReceiverCharacterID,
		Content:             row.Content,
		Status:              enum.ChatMessageStatus(row.Status),
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
		DeletedAt:           row.DeletedAt,
	}, nil
}

func (r *ChatMessageRepo) List(ctx context.Context, req *bizrepo.ChatMessageListReq) ([]*model.ChatMessage, error) {
	query := r.db.ChatMessage.Query().
		Where(
			chatmessageent.ChannelTypeEQ(chatmessageent.ChannelType(req.ChannelType)),
			chatmessageent.ChannelIDEQ(req.ChannelID),
			chatmessageent.StatusEQ(chatmessageent.Status(enum.ChatMessageStatusNormal)),
			chatmessageent.DeletedAtIsNil(),
		)
	if req.BeforeID > 0 {
		query = query.Where(chatmessageent.IDLT(req.BeforeID))
	}
	rows, err := query.Order(gen.Desc(chatmessageent.FieldID)).Limit(req.Size).All(ctx)
	if err != nil {
		return nil, err
	}
	messages := make([]*model.ChatMessage, 0, len(rows))
	for _, row := range rows {
		messages = append(messages, &model.ChatMessage{
			ID:                  row.ID,
			ChannelType:         enum.ChatChannelType(row.ChannelType),
			ChannelID:           row.ChannelID,
			SenderCharacterID:   row.SenderCharacterID,
			ReceiverCharacterID: row.ReceiverCharacterID,
			Content:             row.Content,
			Status:              enum.ChatMessageStatus(row.Status),
			CreatedAt:           row.CreatedAt,
			UpdatedAt:           row.UpdatedAt,
			DeletedAt:           row.DeletedAt,
		})
	}
	return messages, nil
}
