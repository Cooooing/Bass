package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	"content/internal/biz/model"
	"content/internal/data/gen"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, tx *gen.Client, comment *model.Comment) (*model.Comment, error)

	UpdateStatus(ctx context.Context, tx *gen.Client, commentId int64, status v1.CommentStatus) error
	UpdateStat(ctx context.Context, tx *gen.Client, commentId int64, action v1.CommentAction, num int32) error

	Exist(ctx context.Context, tx *gen.Client, req *CommentGetReq) (bool, error)
	GetOne(ctx context.Context, tx *gen.Client, req *CommentGetReq) (*model.Comment, error)
	GetList(ctx context.Context, tx *gen.Client, req *CommentGetReq) ([]*model.Comment, error)
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *CommentGetReq) ([]*model.Comment, *common.PageReply, error)
	GetArticleLastComment(ctx context.Context, tx *gen.Client, req *CommentGetReq) (*model.Comment, error)
	GetArticleLastComments(ctx context.Context, tx *gen.Client, req *CommentGetReq) (map[int64]*model.Comment, error)
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
