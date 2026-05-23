package repo

import (
	"context"
	"user/internal/biz/model"
)

type CheckinRepo interface {
	FindStatByUserID(ctx context.Context, userID int64) (*model.CheckinStat, error)
	UpsertRecord(ctx context.Context, record *model.CheckinRecord) (*model.CheckinRecord, error)
	UpsertStat(ctx context.Context, stat *model.CheckinStat) (*model.CheckinStat, error)
}
