package repo

import (
	"context"
	"time"
)

type CommentViewerActionState struct {
	Liked   bool
	Thanked bool
}

type CommentListItem struct {
	ID                int64
	ArticleID         int64
	Content           string
	ContentRender     string
	Level             int32
	ParentID          *int64
	ReplyID           *int64
	Restriction       int32
	DeletedAt         *time.Time
	ThankCount        int32
	LikeCount         int32
	ReplyCount        int32
	ViewerActionState *CommentViewerActionState
	User              *AccountProfile
	ReplyUser         *AccountProfile
	CreatedBy         *int64
	UpdatedBy         *int64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

type CommentDetail struct {
	ID                int64
	ArticleID         int64
	Content           string
	ContentRender     string
	Level             int32
	ParentID          *int64
	ReplyID           *int64
	Restriction       int32
	DeletedAt         *time.Time
	ThankCount        int32
	LikeCount         int32
	ReplyCount        int32
	ViewerActionState *CommentViewerActionState
	User              *AccountProfile
	ReplyUser         *AccountProfile
	CreatedBy         *int64
	UpdatedBy         *int64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

type CommentThread struct {
	Root           *CommentListItem
	PreviewReplies []*CommentListItem
	ReplyCount     int32
	HasMoreReplies bool
}

type CommentQuery struct {
	CommentID    *int64
	ArticleID    *int64
	ParentID     *int64
	ReplyID      *int64
	Level        *int32
	UserID       *int64
	Restriction  *int32
	Restrictions []int32
	Order        *int32
}

type CreateCommentReq struct {
	UserID    int64
	ArticleID int64
	Content   string
	ReplyID   int64
}

type ListCommentsReq struct {
	UserID int64
	Page   *PageReq
	Query  *CommentQuery
}

type ListCommentsResp struct {
	Page *PageResp
	Rows []*CommentListItem
}

type ListCommentThreadsReq struct {
	UserID            int64
	Page              *PageReq
	ArticleID         int64
	Order             *int32
	ReplyPreviewLimit *int32
}

type ListCommentThreadsResp struct {
	Page *PageResp
	Rows []*CommentThread
}

type ListCommentRepliesReq struct {
	UserID    int64
	Page      *PageReq
	ArticleID int64
	ParentID  int64
	Order     *int32
}

type ListCommentRepliesResp struct {
	Page *PageResp
	Rows []*CommentListItem
}

type ListCommentTimelineReq struct {
	UserID    int64
	Page      *PageReq
	ArticleID int64
	Order     *int32
}

type ListCommentTimelineResp struct {
	Page *PageResp
	Rows []*CommentListItem
}

type LikeCommentReq struct {
	UserID int64
	ID     int64
	Active bool
}

type ThankCommentReq struct {
	UserID int64
	ID     int64
	Active bool
}

type ContentCommentClient interface {
	CreateComment(ctx context.Context, req *CreateCommentReq) (*CommentDetail, error)
	ListComments(ctx context.Context, req *ListCommentsReq) (*ListCommentsResp, error)
	ListCommentThreads(ctx context.Context, req *ListCommentThreadsReq) (*ListCommentThreadsResp, error)
	ListCommentReplies(ctx context.Context, req *ListCommentRepliesReq) (*ListCommentRepliesResp, error)
	ListCommentTimeline(ctx context.Context, req *ListCommentTimelineReq) (*ListCommentTimelineResp, error)
	LikeComment(ctx context.Context, req *LikeCommentReq) (bool, error)
	ThankComment(ctx context.Context, req *ThankCommentReq) (bool, error)
}
