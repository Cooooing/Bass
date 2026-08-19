package usecase

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"
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

func (s *PreferencesUsecase) GetByUserID(ctx context.Context, userID int64) (*model.Preferences, error) {
	preferences, err := s.preferencesRepo.Get(ctx, &repo.PreferencesGetReq{
		UserID: &userID,
	})
	if err != nil || preferences != nil {
		return preferences, err
	}
	return &model.Preferences{
		UserID:      userID,
		Language:    new(enum.LanguageZhCN),
		Timezone:    new("Asia/Shanghai"),
		Theme:       new("default"),
		MobileTheme: new("default"),
	}, nil
}

func (s *PreferencesUsecase) UpsertByUserID(ctx context.Context, preferences *model.Preferences) (*model.Preferences, error) {
	return s.preferencesRepo.UpsertByUserID(ctx, preferences)
}
