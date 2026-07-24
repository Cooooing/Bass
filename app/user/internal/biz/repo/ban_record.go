package repo

import (
	"context"
	"user/internal/biz/model"
)

type BanRecordRepo interface {
	Create(ctx context.Context, record *model.BanRecord) (*model.BanRecord, error)
	Get(ctx context.Context, id int64) (*model.BanRecord, error)
	LatestByUserID(ctx context.Context, userID int64) (*model.BanRecord, error)
}
