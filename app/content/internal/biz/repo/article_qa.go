package repo

import (
	"context"
	"time"
)

// ArticleQARepo 维护问答文章的不可撤销采纳事实。
type ArticleQARepo interface {
	GetAcceptedCommentID(context.Context, int64) (*int64, error)
	Accept(context.Context, int64, int64, time.Time) error
}
