package data

import (
	"common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type CommentActionRecordRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewCommentActionRecordRepo(baseRepo *BaseRepo, client *gen.Client) repo.CommentActionRecordRepo {
	return &CommentActionRecordRepo{
		BaseRepo: baseRepo,
		client:   client,
	}
}

func (a CommentActionRecordRepo) Save(ctx context.Context, client *gen.Client, record *model.CommentActionRecord) (*model.CommentActionRecord, error) {
	// TODO implement me
	panic("implement me")
}

func (a CommentActionRecordRepo) Delete(ctx context.Context, client *gen.Client, commentId int64, userId int64, action v1.CommentAction) error {
	// TODO implement me
	panic("implement me")
}
