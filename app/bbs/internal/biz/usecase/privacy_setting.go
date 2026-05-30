package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type PrivacySettingUsecase struct {
	privacySettingRepo repo.PrivacySettingRepo
}

func NewPrivacySettingUsecase(privacySettingRepo repo.PrivacySettingRepo) *PrivacySettingUsecase {
	return &PrivacySettingUsecase{privacySettingRepo: privacySettingRepo}
}

func (u *PrivacySettingUsecase) GetCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error) {
	return u.privacySettingRepo.GetCurrentPrivacySetting(ctx, req)
}

func (u *PrivacySettingUsecase) UpdateCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Request) (*bbsuserv1.UpdateCurrentPrivacySetting_Reply, error) {
	return u.privacySettingRepo.UpdateCurrentPrivacySetting(ctx, req)
}
