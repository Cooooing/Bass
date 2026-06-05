package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	"content/internal/biz/model"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, comment *model.Comment) (*model.Comment, error)

	UpdateStatus(ctx context.Context, commentId int64, userId int64, status v1.CommentStatus) error
	UpdateStat(ctx context.Context, commentId int64, userId int64, action v1.CommentAction, num int32) error

	Exist(ctx context.Context, req *CommentGetReq) (bool, error)
	Get(ctx context.Context, req *CommentGetReq) (*model.Comment, error)
	GetList(ctx context.Context, req *CommentGetReq) ([]*model.Comment, error)
	Page(ctx context.Context, page *common.PageRequest, req *CommentGetReq) ([]*model.Comment, *common.PageReply, error)
	GetArticleLastComment(ctx context.Context, req *CommentGetReq) (*model.Comment, error)
	GetArticleLastComments(ctx context.Context, req *CommentGetReq) (map[int64]*model.Comment, error)
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
	Order      *v1.CommentOrder
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, record *model.CommentActionRecord) (*model.CommentActionRecord, error)
	Delete(ctx context.Context, commentId int64, userId int64, action v1.CommentAction) (int, error)
	Exist(ctx context.Context, commentId int64, userId int64, action v1.CommentAction) (bool, error)
}
