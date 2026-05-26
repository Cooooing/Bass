package repo

import (
	"context"

	v1 "common/api/gen/content/v1"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/commentactionrecord"
	"content/internal/enum"
)

var _ repo.CommentActionRecordRepo = (*CommentActionRecordRepo)(nil)

type CommentActionRecordRepo struct {
	db *gen.Client
}

func NewCommentActionRecordRepo(db *gen.Client) repo.CommentActionRecordRepo {
	return &CommentActionRecordRepo{db: db}
}

func (r *CommentActionRecordRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *CommentActionRecordRepo) Save(ctx context.Context, record *model.CommentActionRecord) (*model.CommentActionRecord, error) {
	save, err := r.getClient(ctx).CommentActionRecord.Create().
		SetCommentID(record.CommentID).
		SetUserID(record.UserID).
		SetType(commentactionrecord.Type(record.Type)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CommentActionRecord{
		ID:        save.ID,
		CommentID: save.CommentID,
		UserID:    save.UserID,
		Type:      enum.CommentAction(save.Type),
	}, nil
}

func (r *CommentActionRecordRepo) Delete(ctx context.Context, commentId int64, userId int64, action v1.CommentAction) (int, error) {
	dbType, _ := enum.CommentActionMap.ToEnum(action)
	return r.getClient(ctx).CommentActionRecord.Delete().
		Where(commentactionrecord.CommentIDEQ(commentId)).
		Where(commentactionrecord.UserIDEQ(userId)).
		Where(commentactionrecord.TypeEQ(commentactionrecord.Type(dbType))).
		Exec(ctx)
}

func (r *CommentActionRecordRepo) Exist(ctx context.Context, commentId int64, userId int64, action v1.CommentAction) (bool, error) {
	dbType, _ := enum.CommentActionMap.ToEnum(action)
	return r.getClient(ctx).CommentActionRecord.Query().
		Where(commentactionrecord.CommentIDEQ(commentId)).
		Where(commentactionrecord.UserIDEQ(userId)).
		Where(commentactionrecord.TypeEQ(commentactionrecord.Type(dbType))).
		Exist(ctx)
}
