package data

import (
	v1 "common/api/content/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type CommentActionRecordRepo struct {
	*BaseData
	client *gen.Client
}

func NewCommentActionRecordRepo(BaseData *BaseData, client *gen.Client) repo.CommentActionRecordRepo {
	return &CommentActionRecordRepo{
		BaseData: BaseData,
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
