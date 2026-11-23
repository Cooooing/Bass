package repo

import (
	"context"
	"time"
)

// ArticleQARepo 维护问答文章的不可撤销采纳事实。
type ArticleQARepo interface {
	GetAcceptedCommentID(ctx context.Context, articleID int64) (*int64, error)
	Accept(ctx context.Context, articleID, commentID int64, now time.Time) error
}
