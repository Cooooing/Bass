package repo

import (
	"context"
	"user/internal/biz/model"
)

type UserCheckinRepo interface {
	GetStatByUserID(ctx context.Context, userID int64) (*model.UserCheckinStat, error)
	UpsertRecord(ctx context.Context, record *model.UserCheckinRecord) (*model.UserCheckinRecord, error)
	UpdateStat(ctx context.Context, stat *model.UserCheckinStat) (*model.UserCheckinStat, error)
}
