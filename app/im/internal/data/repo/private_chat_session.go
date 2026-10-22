package repo

import (
	utilent "common/pkg/util/ent"
	"context"
	"time"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/privatechatsession"
)

type PrivateChatSessionRepo struct{ db *gen.Client }

func NewPrivateChatSessionRepo(db *gen.Client) bizrepo.PrivateChatSessionRepo {
	return &PrivateChatSessionRepo{db: db}
}

func (r *PrivateChatSessionRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *PrivateChatSessionRepo) Get(ctx context.Context, userID, peerUserID int64) (*model.PrivateChatSession, error) {
	row, err := r.client(ctx).PrivateChatSession.Query().Where(privatechatsession.UserIDEQ(userID), privatechatsession.PeerUserIDEQ(peerUserID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *PrivateChatSessionRepo) List(ctx context.Context, userID int64, limit, offset int) ([]*model.PrivateChatSession, int64, error) {
	query := r.client(ctx).PrivateChatSession.Query().Where(privatechatsession.UserIDEQ(userID)).Order(gen.Desc(privatechatsession.FieldIsPinned), gen.Desc(privatechatsession.FieldLastMessageAt))
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	rows, err := query.Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*model.PrivateChatSession, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.toModel(row))
	}
	return result, int64(total), nil
}

func (r *PrivateChatSessionRepo) CountUnread(ctx context.Context, userID int64) (uint32, error) {
	rows, err := r.client(ctx).PrivateChatSession.Query().Where(privatechatsession.UserIDEQ(userID)).Select(privatechatsession.FieldUnreadCount).All(ctx)
	if err != nil {
		return 0, err
	}
	var total uint32
	for _, row := range rows {
		total += row.UnreadCount
	}
	return total, nil
}

func (r *PrivateChatSessionRepo) Save(ctx context.Context, item *model.PrivateChatSession) (*model.PrivateChatSession, error) {
	row, err := r.client(ctx).PrivateChatSession.Create().SetUserID(item.UserID).SetPeerUserID(item.PeerUserID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(row), nil
}

func (r *PrivateChatSessionRepo) UpdateAfterSend(ctx context.Context, userID, peerUserID, messageID int64, incrementUnread bool) error {
	update := r.client(ctx).PrivateChatSession.Update().Where(privatechatsession.UserIDEQ(userID), privatechatsession.PeerUserIDEQ(peerUserID)).SetLastMessageID(messageID).SetLastMessageAt(time.Now())
	if incrementUnread {
		update.AddUnreadCount(1)
	}
	return update.Exec(ctx)
}

func (r *PrivateChatSessionRepo) MarkRead(ctx context.Context, userID, peerUserID, messageID int64) error {
	return r.client(ctx).PrivateChatSession.Update().Where(privatechatsession.UserIDEQ(userID), privatechatsession.PeerUserIDEQ(peerUserID)).SetLastReadMessageID(messageID).SetUnreadCount(0).Exec(ctx)
}

func (r *PrivateChatSessionRepo) SetMuted(ctx context.Context, userID, peerUserID int64, muted bool) error {
	return r.client(ctx).PrivateChatSession.Update().Where(privatechatsession.UserIDEQ(userID), privatechatsession.PeerUserIDEQ(peerUserID)).SetIsMuted(muted).Exec(ctx)
}

func (r *PrivateChatSessionRepo) SetPinned(ctx context.Context, userID, peerUserID int64, pinned bool) error {
	return r.client(ctx).PrivateChatSession.Update().Where(privatechatsession.UserIDEQ(userID), privatechatsession.PeerUserIDEQ(peerUserID)).SetIsPinned(pinned).Exec(ctx)
}

func (*PrivateChatSessionRepo) toModel(row *gen.PrivateChatSession) *model.PrivateChatSession {
	return &model.PrivateChatSession{ID: row.ID, UserID: row.UserID, PeerUserID: row.PeerUserID, LastMessageID: row.LastMessageID, LastMessageAt: row.LastMessageAt, LastReadMessageID: row.LastReadMessageID, UnreadCount: row.UnreadCount, Muted: row.IsMuted, Pinned: row.IsPinned}
}
