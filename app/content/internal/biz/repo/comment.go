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

	UpdateStatus(ctx context.Context, tx *gen.Client, commentId int64, status v1.CommentStatus) error
	UpdateStat(ctx context.Context, tx *gen.Client, commentId int64, action v1.CommentAction, num int32) error

	Exist(ctx context.Context, tx *gen.Client, req *CommentGetReq) (bool, error)
	GetOne(ctx context.Context, tx *gen.Client, req *CommentGetReq) (*model.Comment, error)
	GetList(ctx context.Context, tx *gen.Client, req *CommentGetReq) ([]*model.Comment, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *CommentGetReq) ([]*model.Comment, *cv1.PageReply, error)
	GetArticleLastComment(ctx context.Context, tx *gen.Client, req *CommentGetReq) (*model.Comment, error)
	GetArticleLastComments(ctx context.Context, tx *gen.Client, req *CommentGetReq) (dict.Map[int64, *model.Comment], error)
}

type CommentGetReq struct {
	CommentId  *int64
	CommentIds []int64
	ParentId   *int64
	ReplyId    *int64
	ArticleId  *int64
	ArticleIds []int64
	CreatedBy  *int64
	Status     *v1.CommentStatus
	Level      *int32
	Order      *int32

	WithArticle bool
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, record *model.CommentActionRecord) (*model.CommentActionRecord, error)
	Delete(ctx context.Context, tx *gen.Client, commentId int64, userId int64, action v1.CommentAction) error
}
