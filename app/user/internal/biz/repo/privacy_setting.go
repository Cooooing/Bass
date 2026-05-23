package repo

import (
	"context"
	"user/internal/biz/model"
)

type PrivacySettingRepo interface {
	FindByUserID(ctx context.Context, userID int64) (*model.PrivacySetting, error)
	UpsertByUserID(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error)
	Update(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error)
}
