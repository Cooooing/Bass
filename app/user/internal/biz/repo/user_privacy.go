package repo

import (
	"context"
	"user/internal/biz/model"
)

type UserPrivacyRepo interface {
	GetByUserID(ctx context.Context, userID int64) (*model.UserPrivacy, error)
	Update(ctx context.Context, p *model.UserPrivacy) (*model.UserPrivacy, error)
}
