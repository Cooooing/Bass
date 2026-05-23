package repo

import (
	"context"
	"user/internal/biz/model"
)

type LocationRepo interface {
	FindByUserID(ctx context.Context, userID int64) (*model.Location, error)
	UpsertByUserID(ctx context.Context, l *model.Location) (*model.Location, error)
	Update(ctx context.Context, l *model.Location) (*model.Location, error)
}
