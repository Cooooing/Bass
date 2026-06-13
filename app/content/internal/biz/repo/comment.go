package repo

import (
	"common/proto/gen/common"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, comment *model.Comment) (*model.Comment, error)

	UpdateRestriction(ctx context.Context, commentId int64, restriction enum.ContentRestriction, updatedBy int64) error
	AddStats(ctx context.Context, commentId int64, stats CommentStatUpdate, updatedBy *int64) error

	Exist(ctx context.Context, req *CommentGetReq) (bool, error)
	Get(ctx context.Context, req *CommentGetReq) (*model.Comment, error)
	List(ctx context.Context, req *CommentGetReq) ([]*model.Comment, error)
	Map(ctx context.Context, req *CommentGetReq) (map[int64]*model.Comment, error)
	Count(ctx context.Context, req *CommentGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *CommentGetReq) ([]*model.Comment, *common.PageReply, error)
	ListReplyPreviews(ctx context.Context, req *CommentReplyPreviewReq) ([]*CommentReplyPreview, error)
	GetArticleLastComment(ctx context.Context, req *CommentGetReq) (*model.Comment, error)
	MapArticleLastComments(ctx context.Context, req *CommentGetReq) (map[int64]*model.Comment, error)
}

type CommentStatUpdate struct {
	ThankCount int32
	LikeCount  int32
	ReplyCount int32
}

type CommentGetReq struct {
	CommentId    *int64
	CommentIds   []int64
	ParentId     *int64
	ReplyId      *int64
	ArticleId    *int64
	ArticleIds   []int64
	CreatedBy    *int64
	Restriction  *enum.ContentRestriction
	Restrictions []enum.ContentRestriction
	Level        *int32
	Order        *enum.CommentOrder
}

type CommentReplyPreviewReq struct {
	ArticleId      int64
	ParentIds      []int64
	LimitPerParent int32
	Restriction    *enum.ContentRestriction
	Restrictions   []enum.ContentRestriction
	Order          *enum.CommentOrder
}

type CommentReplyPreview struct {
	ParentId int64
	Rows     []*model.Comment
}

type CommentActionRecordRepo interface {
	// Save 保存行为记录，返回值表示是否实际写入了新记录；唯一约束冲突表示记录已存在。
	Save(ctx context.Context, record *model.CommentActionRecord) (bool, error)
	Delete(ctx context.Context, commentId int64, userId int64, action enum.CommentAction) (int, error)
	Exist(ctx context.Context, req *CommentActionRecordReq) (bool, error)
	Get(ctx context.Context, req *CommentActionRecordReq) (*model.CommentActionRecord, error)
	List(ctx context.Context, req *CommentActionRecordReq) ([]*model.CommentActionRecord, error)
	Map(ctx context.Context, req *CommentActionRecordReq) (map[int64]*model.CommentActionRecord, error)
	Count(ctx context.Context, req *CommentActionRecordReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *CommentActionRecordReq) ([]*model.CommentActionRecord, *common.PageReply, error)
}

type CommentActionRecordReq struct {
	ID         *int64
	IDs        []int64
	CommentId  *int64
	CommentIds []int64
	UserId     *int64
	UserIds    []int64
	Type       *enum.CommentAction
	Types      []enum.CommentAction
}
