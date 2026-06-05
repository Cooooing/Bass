package model

import "content/internal/enum"

type CommentActionRecord struct {
	ID        int64
	CommentID int64
	UserID    int64
	Type      enum.CommentAction
}
