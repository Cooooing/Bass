package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"
)

type PrivacySettingUsecase struct {
	privacySettingClient repo.PrivacySettingClient
}

func NewPrivacySettingUsecase(privacySettingClient repo.PrivacySettingClient) *PrivacySettingUsecase {
	return &PrivacySettingUsecase{privacySettingClient: privacySettingClient}
}

func (u *PrivacySettingUsecase) GetCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error) {
	return u.privacySettingClient.GetCurrentPrivacySetting(ctx, req)
}

func (u *PrivacySettingUsecase) UpdateCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Request) (*bbsuserv1.UpdateCurrentPrivacySetting_Reply, error) {
	return u.privacySettingClient.UpdateCurrentPrivacySetting(ctx, req)
}
