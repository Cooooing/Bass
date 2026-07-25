package usecase

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type PreferencesUsecase struct {
	preferencesRepo repo.PreferencesRepo
}

func NewPreferencesUsecase(
	preferencesRepo repo.PreferencesRepo,
) *PreferencesUsecase {
	return &PreferencesUsecase{
		preferencesRepo: preferencesRepo,
	}
}

func (s *PreferencesUsecase) GetByUserID(
	ctx context.Context,
	userID int64,
) (*model.Preferences, error) {
	return s.preferencesRepo.Get(ctx, &repo.PreferencesGetReq{
		UserID: &userID,
	})
}

func (s *PreferencesUsecase) UpsertByUserID(
	ctx context.Context,
	preferences *model.Preferences,
) (*model.Preferences, error) {
	return s.preferencesRepo.UpsertByUserID(ctx, preferences)
}
