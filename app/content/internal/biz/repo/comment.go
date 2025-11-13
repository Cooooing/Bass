package repo

import (
	cv1 "common/api/common/v1"
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
	GetList(ctx context.Context, tx *gen.Client, req *CommentGetReq) ([]*model.Comment, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *CommentGetReq) ([]*model.Comment, *cv1.PageReply, error)
	GetArticleLastComment(ctx context.Context, tx *gen.Client, articleId int64) (*model.Comment, error)
	GetArticleLastComments(ctx context.Context, tx *gen.Client, articleIds []int64) (dict.Map[int64, *model.Comment], error)
}

type CommentGetReq struct {
	CommentId *int64
	ArticleId *int64
	UserId    *int64
	Order     *int32
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, record *model.CommentActionRecord) (*model.CommentActionRecord, error)
	Delete(ctx context.Context, tx *gen.Client, articleId int64, userId int64, action cv1.CommentAction) error
}
