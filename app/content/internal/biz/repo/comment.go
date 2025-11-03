package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"common/pkg/util/collections/dict"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, tx *gen.Client, comment *model.Comment) (*model.Comment, error)

	UpdateStatus(ctx context.Context, tx *gen.Client, commentId int64, status cv1.CommentStatus) error
	UpdateStat(ctx context.Context, tx *gen.Client, commentId int64, action cv1.CommentAction, num int32) error

	Exist(ctx context.Context, tx *gen.Client, id int64) (bool, error)
	GetById(ctx context.Context, tx *gen.Client, id int64) (*model.Comment, error)
	GetList(ctx context.Context, tx *gen.Client, req *v1.GetCommentRequest) (*v1.GetCommentReply, error)
	GetArticleLastComment(ctx context.Context, tx *gen.Client, articleId int64) (*model.Comment, error)
	GetArticleLastComments(ctx context.Context, tx *gen.Client, articleIds []int64) (dict.Map[int64, *model.Comment], error)
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, record *model.CommentActionRecord) (*model.CommentActionRecord, error)
	Delete(ctx context.Context, tx *gen.Client, articleId int64, userId int64, action cv1.CommentAction) error
}
