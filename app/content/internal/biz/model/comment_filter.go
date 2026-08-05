package model

import "content/internal/enum"

type CommentFilter struct {
	CommentID    *int64
	CommentIDs   []int64
	ArticleID    *int64
	ArticleIDs   []int64
	ParentID     *int64
	ReplyID      *int64
	CreatedBy    *int64
	Restriction  *enum.ContentRestriction
	Restrictions []enum.ContentRestriction
	Level        *int32
	Order        *enum.CommentOrder
}
