package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	UpdateRestriction(ctx context.Context, req *CommentUpdateRestrictionReq) error
	AddStats(ctx context.Context, req *CommentAddStatsReq) error
	Exist(ctx context.Context, req *CommentGetReq) (bool, error)
	Get(ctx context.Context, req *CommentGetReq) (*model.Comment, error)
	List(ctx context.Context, req *CommentGetReq) ([]*model.Comment, error)
	Map(ctx context.Context, commentGetReq *CommentGetReq) (map[int64]*model.Comment, error)
	Count(ctx context.Context, req *CommentGetReq) (int, error)
	Page(ctx context.Context, req *CommentGetReq) (*CommentPageResp, error)
	ListReplyPreviews(ctx context.Context, req *CommentReplyPreviewReq) ([]*CommentReplyPreview, error)
	GetArticleLastComment(ctx context.Context, req *CommentGetReq) (*model.Comment, error)
	MapArticleLastComments(ctx context.Context, req *CommentGetReq) (map[int64]*model.Comment, error)
}

type CommentUpdateRestrictionReq struct {
	CommentID   int64
	Restriction enum.ContentRestriction
	UpdatedBy   int64
}

type CommentAddStatsReq struct {
	CommentID int64
	Stats     CommentStatUpdate
	UpdatedBy *int64
}

type CommentStatUpdate struct {
	ThankCount int32
	LikeCount  int32
	ReplyCount int32
}

type CommentPageResp struct {
	Rows []*model.Comment
	Page *base.PageResp
}

type CommentGetReq struct {
	Page   *base.PageRequest
	Filter *model.CommentFilter
	Scope  *model.CommentScopeFilter
}

type CommentReplyPreviewReq struct {
	Filter         *model.CommentFilter
	Scope          *model.CommentScopeFilter
	ParentIDs      []int64
	LimitPerParent int32
}

type CommentReplyPreview struct {
	ParentId int64
	Rows     []*model.Comment
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, record *model.CommentActionRecord) (bool, error)
	Delete(ctx context.Context, req *CommentActionRecordDeleteReq) (int, error)
	Exist(ctx context.Context, req *CommentActionRecordReq) (bool, error)
	Get(ctx context.Context, req *CommentActionRecordReq) (*model.CommentActionRecord, error)
	List(ctx context.Context, req *CommentActionRecordReq) ([]*model.CommentActionRecord, error)
	Map(ctx context.Context, req *CommentActionRecordReq) (map[int64]*model.CommentActionRecord, error)
	Count(ctx context.Context, req *CommentActionRecordReq) (int, error)
	Page(ctx context.Context, req *CommentActionRecordReq) (*CommentActionRecordPageResp, error)
}

type CommentActionRecordDeleteReq struct {
	CommentID int64
	UserID    int64
	Action    enum.CommentAction
}

type CommentActionRecordPageResp struct {
	Rows []*model.CommentActionRecord
	Page *base.PageResp
}

type CommentActionRecordReq struct {
	Page       *base.PageRequest
	ID         *int64
	IDs        []int64
	CommentId  *int64
	CommentIds []int64
	UserId     *int64
	UserIds    []int64
	Type       *enum.CommentAction
	Types      []enum.CommentAction
}
