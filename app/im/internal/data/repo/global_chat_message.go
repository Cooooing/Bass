package repo

import (
	"context"
	"time"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/globalchatmessage"
	"im/internal/enum"
)

type GlobalChatMessageRepo struct{ db *gen.Client }

func NewGlobalChatMessageRepo(db *gen.Client) bizrepo.GlobalChatMessageRepo {
	return &GlobalChatMessageRepo{db: db}
}

func (r *GlobalChatMessageRepo) Get(ctx context.Context, id int64) (*model.GlobalChatMessage, error) {
	row, err := r.db.GlobalChatMessage.Get(ctx, id)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *GlobalChatMessageRepo) GetByClientMessageID(ctx context.Context, senderID int64, clientMessageID string) (*model.GlobalChatMessage, error) {
	row, err := r.db.GlobalChatMessage.Query().Where(globalchatmessage.SenderIDEQ(senderID), globalchatmessage.ClientMessageIDEQ(clientMessageID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *GlobalChatMessageRepo) Save(ctx context.Context, item *model.GlobalChatMessage) (*model.GlobalChatMessage, error) {
	row, err := r.db.GlobalChatMessage.Create().SetSenderID(item.SenderID).SetClientMessageID(item.ClientMessageID).SetContent(item.Content).SetMessageType(globalchatmessage.MessageType(item.Type)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *GlobalChatMessageRepo) List(ctx context.Context, beforeID int64, limit, offset int) ([]*model.GlobalChatMessage, int64, error) {
	query := r.db.GlobalChatMessage.Query()
	if beforeID > 0 {
		query = query.Where(globalchatmessage.IDLT(beforeID))
	}
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	rows, err := query.Order(gen.Desc(globalchatmessage.FieldID)).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*model.GlobalChatMessage, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.toModel(row))
	}
	return result, int64(total), nil
}

func (r *GlobalChatMessageRepo) Revoke(ctx context.Context, id, revokedBy int64) (*model.GlobalChatMessage, error) {
	now := time.Now()
	row, err := r.db.GlobalChatMessage.UpdateOneID(id).SetIsRevoked(true).SetRevokedAt(now).SetRevokedBy(revokedBy).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (*GlobalChatMessageRepo) toModel(row *gen.GlobalChatMessage) *model.GlobalChatMessage {
	return &model.GlobalChatMessage{ID: row.ID, SenderID: row.SenderID, ClientMessageID: row.ClientMessageID, Content: row.Content, Type: enum.MessageType(row.MessageType), Revoked: row.IsRevoked, RevokedAt: row.RevokedAt, RevokedBy: row.RevokedBy, CreatedAt: row.CreatedAt}
}
