package model

import "content/internal/enum"

type ArticleActionRecord struct {
	ID        int64
	ArticleID int64
	UserID    int64
	Type      enum.ArticleAction
}
