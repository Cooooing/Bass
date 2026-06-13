package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"
)

type PreferencesUsecase struct {
	preferencesClient repo.PreferencesClient
}

func NewPreferencesUsecase(preferencesClient repo.PreferencesClient) *PreferencesUsecase {
	return &PreferencesUsecase{preferencesClient: preferencesClient}
}

func (u *PreferencesUsecase) GetCurrentPreferences(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	return u.preferencesClient.GetCurrentPreferences(ctx, req)
}

func (u *PreferencesUsecase) UpdateCurrentPreferences(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
	return u.preferencesClient.UpdateCurrentPreferences(ctx, req)
}
