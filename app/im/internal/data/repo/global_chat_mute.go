package repo

import (
	"context"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/globalchatmute"
)

type GlobalChatMuteRepo struct{ db *gen.Client }

func NewGlobalChatMuteRepo(db *gen.Client) bizrepo.GlobalChatMuteRepo {
	return &GlobalChatMuteRepo{db: db}
}

func (r *GlobalChatMuteRepo) Get(ctx context.Context, userID int64) (*model.GlobalChatMute, error) {
	row, err := r.db.GlobalChatMute.Query().Where(globalchatmute.UserIDEQ(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.GlobalChatMute{UserID: row.UserID, MutedUntil: row.MutedUntil, Reason: row.Reason, OperatorUserID: row.OperatorUserID}, nil
}

func (r *GlobalChatMuteRepo) Save(ctx context.Context, item *model.GlobalChatMute) error {
	return r.db.GlobalChatMute.Create().SetUserID(item.UserID).SetMutedUntil(item.MutedUntil).SetNillableReason(item.Reason).SetOperatorUserID(item.OperatorUserID).OnConflictColumns(globalchatmute.FieldUserID).UpdateNewValues().Exec(ctx)
}

func (r *GlobalChatMuteRepo) Delete(ctx context.Context, userID int64) error {
	_, err := r.db.GlobalChatMute.Delete().Where(globalchatmute.UserIDEQ(userID)).Exec(ctx)
	return err
}
