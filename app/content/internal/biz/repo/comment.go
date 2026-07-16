package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type CommentRepo interface {
	Save(ctx context.Context, req *CommentSaveReq) (*CommentSaveResponse, error)
	UpdateRestriction(ctx context.Context, req *CommentUpdateRestrictionReq) (*CommentUpdateRestrictionResponse, error)
	AddStats(ctx context.Context, req *CommentAddStatsReq) (*CommentAddStatsResponse, error)
	Exist(ctx context.Context, req *CommentGetReq) (*CommentExistResponse, error)
	Get(ctx context.Context, req *CommentGetReq) (*CommentGetResponse, error)
	List(ctx context.Context, req *CommentGetReq) (*CommentListResponse, error)
	Map(ctx context.Context, req *CommentMapReq) (*CommentMapResponse, error)
	Count(ctx context.Context, req *CommentGetReq) (*CommentCountResponse, error)
	Page(ctx context.Context, req *CommentGetReq) (*CommentPageResponse, error)
	ListReplyPreviews(ctx context.Context, req *CommentReplyPreviewReq) (*CommentReplyPreviewResponse, error)
	GetArticleLastComment(ctx context.Context, req *CommentGetReq) (*CommentGetResponse, error)
	MapArticleLastComments(ctx context.Context, req *CommentGetReq) (*CommentArticleLastCommentsResponse, error)
}

type CommentSaveReq struct {
	Comment *model.Comment
}

type CommentSaveResponse struct {
	Comment *model.Comment
}

type CommentUpdateRestrictionReq struct {
	CommentID   int64
	Restriction enum.ContentRestriction
	UpdatedBy   int64
}

type CommentUpdateRestrictionResponse struct{}

type CommentAddStatsReq struct {
	CommentID int64
	Stats     CommentStatUpdate
	UpdatedBy *int64
}

type CommentAddStatsResponse struct{}

type CommentStatUpdate struct {
	ThankCount int32
	LikeCount  int32
	ReplyCount int32
}

type CommentExistResponse struct {
	Exist bool
}

type CommentGetResponse struct {
	Comment *model.Comment
}

type CommentListResponse struct {
	Rows []*model.Comment
}

type CommentMapReq struct {
	*CommentGetReq
}

type CommentMapResponse struct {
	Rows map[int64]*model.Comment
}

type CommentCountResponse struct {
	Count int
}

type CommentPageResponse struct {
	Rows []*model.Comment
	Page *base.PageResponse
}

type CommentGetReq struct {
	Page         *base.PageRequest
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

type CommentReplyPreviewResponse struct {
	Rows []*CommentReplyPreview
}

type CommentReplyPreview struct {
	ParentId int64
	Rows     []*model.Comment
}

type CommentArticleLastCommentsResponse struct {
	Rows map[int64]*model.Comment
}

type CommentActionRecordRepo interface {
	Save(ctx context.Context, req *CommentActionRecordSaveReq) (*CommentActionRecordSaveResponse, error)
	Delete(ctx context.Context, req *CommentActionRecordDeleteReq) (*CommentActionRecordDeleteResponse, error)
	Exist(ctx context.Context, req *CommentActionRecordReq) (*CommentActionRecordExistResponse, error)
	Get(ctx context.Context, req *CommentActionRecordReq) (*CommentActionRecordGetResponse, error)
	List(ctx context.Context, req *CommentActionRecordReq) (*CommentActionRecordListResponse, error)
	Map(ctx context.Context, req *CommentActionRecordReq) (*CommentActionRecordMapResponse, error)
	Count(ctx context.Context, req *CommentActionRecordReq) (*CommentActionRecordCountResponse, error)
	Page(ctx context.Context, req *CommentActionRecordReq) (*CommentActionRecordPageResponse, error)
}

type CommentActionRecordSaveReq struct {
	Record *model.CommentActionRecord
}

type CommentActionRecordSaveResponse struct {
	Created bool
}

type CommentActionRecordDeleteReq struct {
	CommentID int64
	UserID    int64
	Action    enum.CommentAction
}

type CommentActionRecordDeleteResponse struct {
	Deleted int
}

type CommentActionRecordExistResponse struct {
	Exist bool
}

type CommentActionRecordGetResponse struct {
	Record *model.CommentActionRecord
}

type CommentActionRecordListResponse struct {
	Rows []*model.CommentActionRecord
}

type CommentActionRecordMapResponse struct {
	Rows map[int64]*model.CommentActionRecord
}

type CommentActionRecordCountResponse struct {
	Count int
}

type CommentActionRecordPageResponse struct {
	Rows []*model.CommentActionRecord
	Page *base.PageResponse
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
