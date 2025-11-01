package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, client *gen.Client, comment *model.Comment) (*model.Comment, error)

	UpdateStatus(ctx context.Context, client *gen.Client, commentId int64, status cv1.CommentStatus) error
	UpdateStat(ctx context.Context, client *gen.Client, commentId int64, action cv1.CommentAction, num int32) error

	GetCommentById(ctx context.Context, client *gen.Client, id int64) (*model.Comment, error)
	GetCommentList(ctx context.Context, client *gen.Client, req *v1.GetCommentRequest) (*v1.GetCommentReply, error)
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, client *gen.Client, record *model.CommentActionRecord) (*model.CommentActionRecord, error)
	Delete(ctx context.Context, client *gen.Client, articleId int64, userId int64, action cv1.CommentAction) error
}
