package usecase

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type PrivacySettingUsecase struct {
	privacySettingRepo repo.PrivacySettingRepo
}

func NewPrivacySettingUsecase(
	privacySettingRepo repo.PrivacySettingRepo,
) *PrivacySettingUsecase {
	return &PrivacySettingUsecase{
		privacySettingRepo: privacySettingRepo,
	}
}

func (s *PrivacySettingUsecase) GetByUserID(ctx context.Context, userID int64) (*model.PrivacySetting, error) {
	setting, err := s.privacySettingRepo.Get(ctx, &repo.PrivacySettingGetReq{
		UserID: &userID,
	})
	if err != nil || setting != nil {
		return setting, err
	}
	public := true
	return &model.PrivacySetting{
		UserID:             userID,
		PublicPoints:       &public,
		PublicFollowers:    &public,
		PublicArticles:     &public,
		PublicComments:     &public,
		PublicOnlineStatus: &public,
		PublicLocation:     &public,
	}, nil
}

func (s *PrivacySettingUsecase) UpsertByUserID(ctx context.Context, setting *model.PrivacySetting) (*model.PrivacySetting, error) {
	return s.privacySettingRepo.UpsertByUserID(ctx, setting)
}
