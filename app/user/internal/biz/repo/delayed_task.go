package repo

import (
	"context"
	"time"
)

type DelayedTaskClient interface {
	RegisterUnbanExpired(ctx context.Context, userID int64, banRecordID int64, executeAt time.Time) error
}
