package usecase

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type PrivacySettingUsecase struct {
	privacySettingRepo repo.PrivacySettingRepo
}

func NewPrivacySettingUsecase(privacySettingRepo repo.PrivacySettingRepo) *PrivacySettingUsecase {
	return &PrivacySettingUsecase{privacySettingRepo: privacySettingRepo}
}

func (s *PrivacySettingUsecase) GetByUserID(ctx context.Context, userID int64) (*model.PrivacySetting, error) {
	return s.privacySettingRepo.Get(ctx, &repo.PrivacySettingGetReq{UserID: &userID})
}

func (s *PrivacySettingUsecase) UpsertByUserID(ctx context.Context, privacySetting *model.PrivacySetting) (*model.PrivacySetting, error) {
	return s.privacySettingRepo.UpsertByUserID(ctx, privacySetting)
}
