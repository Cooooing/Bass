package repo

import (
	"context"

	"notify/internal/biz/model"
)

type ContentClient interface {
	GetArticle(ctx context.Context, articleID int64) (*model.ContentArticle, error)
	GetComment(ctx context.Context, commentID int64) (*model.ContentComment, error)
	ListArticleWatchers(ctx context.Context, articleID int64) ([]int64, error)
	ListPollParticipants(ctx context.Context, articleID int64) ([]int64, error)
	ListLotteryParticipants(ctx context.Context, articleID int64) ([]int64, error)
}
