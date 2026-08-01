package repo

import (
	"context"
	"time"
)

type DelayedTaskClient interface {
	RegisterPublishScheduledArticle(ctx context.Context, articleID int64, publishAt time.Time) error
	CancelPublishScheduledArticle(ctx context.Context, articleID int64) error
}
