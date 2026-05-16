package repo

import (
	"context"
	"user/internal/biz/model"
)

type UserLocationRepo interface {
	GetByUserID(ctx context.Context, userID int64) (*model.UserLocation, error)
	Update(ctx context.Context, l *model.UserLocation) (*model.UserLocation, error)
}
