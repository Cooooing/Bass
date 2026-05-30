package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type PreferencesUsecase struct {
	preferencesRepo repo.PreferencesRepo
}

func NewPreferencesUsecase(preferencesRepo repo.PreferencesRepo) *PreferencesUsecase {
	return &PreferencesUsecase{preferencesRepo: preferencesRepo}
}

func (u *PreferencesUsecase) GetCurrentPreferences(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	return u.preferencesRepo.GetCurrentPreferences(ctx, req)
}

func (u *PreferencesUsecase) UpdateCurrentPreferences(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
	return u.preferencesRepo.UpdateCurrentPreferences(ctx, req)
}
