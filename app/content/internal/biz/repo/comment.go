package repo

import (
	v1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, client *gen.Client, comment *model.Comment) (*model.Comment, error)

	UpdateStatus(ctx context.Context, client *gen.Client, commentId int, status v1.CommentStatus) error
	UpdateStat(ctx context.Context, client *gen.Client, commentId int, action v1.CommentAction, num int) error
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, client *gen.Client, record *model.CommentActionRecord) (*model.CommentActionRecord, error)
	Delete(ctx context.Context, client *gen.Client, articleId int, userId int, action v1.CommentAction) error
}
