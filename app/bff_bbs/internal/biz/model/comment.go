package model

import (
	"time"

	"bff_bbs/internal/enum"
)

type CommentViewerActionState struct {
	Liked   bool
	Thanked bool
}

type ArticleBrief struct {
	ID    int64
	Title string
}

type CommentListItem struct {
	ID                int64
	ArticleID         int64
	Article           *ArticleBrief
	Content           string
	ContentRender     string
	Level             int32
	ParentID          *int64
	ReplyID           *int64
	Restriction       enum.ContentRestriction
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
