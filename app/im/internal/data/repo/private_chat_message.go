package repo

import (
	utilent "common/pkg/util/ent"
	"context"
	"time"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/privatechatmessage"
	"im/internal/enum"
)

type PrivateChatMessageRepo struct{ db *gen.Client }

func NewPrivateChatMessageRepo(db *gen.Client) bizrepo.PrivateChatMessageRepo {
	return &PrivateChatMessageRepo{db: db}
}
func (r *PrivateChatMessageRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *PrivateChatMessageRepo) Get(ctx context.Context, id int64) (*model.PrivateChatMessage, error) {
	row, err := r.client(ctx).PrivateChatMessage.Get(ctx, id)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *PrivateChatMessageRepo) Map(ctx context.Context, ids []int64) (map[int64]*model.PrivateChatMessage, error) {
	if len(ids) == 0 {
		return map[int64]*model.PrivateChatMessage{}, nil
	}
	rows, err := r.client(ctx).PrivateChatMessage.Query().Where(privatechatmessage.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.PrivateChatMessage, len(rows))
	for _, row := range rows {
		result[row.ID] = r.toModel(row)
	}
	return result, nil
}

func (r *PrivateChatMessageRepo) GetByClientMessageID(ctx context.Context, senderID int64, clientMessageID string) (*model.PrivateChatMessage, error) {
	row, err := r.client(ctx).PrivateChatMessage.Query().Where(privatechatmessage.SenderIDEQ(senderID), privatechatmessage.ClientMessageIDEQ(clientMessageID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *PrivateChatMessageRepo) Save(ctx context.Context, item *model.PrivateChatMessage) (*model.PrivateChatMessage, error) {
	row, err := r.client(ctx).PrivateChatMessage.Create().SetSenderID(item.SenderID).SetReceiverID(item.ReceiverID).SetClientMessageID(item.ClientMessageID).SetContent(item.Content).SetMessageType(privatechatmessage.MessageType(item.Type)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *PrivateChatMessageRepo) List(ctx context.Context, userID, peerUserID, beforeID int64, limit, offset int) ([]*model.PrivateChatMessage, int64, error) {
	query := r.client(ctx).PrivateChatMessage.Query().Where(privatechatmessage.Or(privatechatmessage.And(privatechatmessage.SenderIDEQ(userID), privatechatmessage.ReceiverIDEQ(peerUserID)), privatechatmessage.And(privatechatmessage.SenderIDEQ(peerUserID), privatechatmessage.ReceiverIDEQ(userID))))
	if beforeID > 0 {
		query = query.Where(privatechatmessage.IDLT(beforeID))
	}
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	rows, err := query.Order(gen.Desc(privatechatmessage.FieldID)).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*model.PrivateChatMessage, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.toModel(row))
	}
	return result, int64(total), nil
}

func (r *PrivateChatMessageRepo) Revoke(ctx context.Context, id, revokedBy int64) (*model.PrivateChatMessage, error) {
	now := time.Now()
	row, err := r.client(ctx).PrivateChatMessage.UpdateOneID(id).SetIsRevoked(true).SetRevokedAt(now).SetRevokedBy(revokedBy).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (*PrivateChatMessageRepo) toModel(row *gen.PrivateChatMessage) *model.PrivateChatMessage {
	return &model.PrivateChatMessage{ID: row.ID, SenderID: row.SenderID, ReceiverID: row.ReceiverID, ClientMessageID: row.ClientMessageID, Content: row.Content, Type: enum.MessageType(row.MessageType), Revoked: row.IsRevoked, RevokedAt: row.RevokedAt, RevokedBy: row.RevokedBy, CreatedAt: row.CreatedAt}
}
