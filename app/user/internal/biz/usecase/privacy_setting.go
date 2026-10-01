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
		UserID:               userID,
		PublicPoints:         &public,
		PublicFollowerList:   &public,
		PublicFollowingList:  &public,
		PublicArticleList:    &public,
		PublicCommentList:    &public,
		PublicOnlineStatus:   &public,
		PublicLocation:       &public,
		PublicMoonbreezeList: &public,
	}, nil
}

// MapByUserIDs returns effective settings, including defaults for accounts that
// have not materialized their optional privacy row yet.
func (s *PrivacySettingUsecase) MapByUserIDs(ctx context.Context, userIDs []int64) (map[int64]*model.PrivacySetting, error) {
	result := make(map[int64]*model.PrivacySetting, len(userIDs))
	if len(userIDs) == 0 {
		return result, nil
	}
	rows, err := s.privacySettingRepo.Map(ctx, &repo.PrivacySettingGetReq{UserIDs: userIDs})
	if err != nil {
		return nil, err
	}
	public := true
	for _, userID := range userIDs {
		if setting := rows[userID]; setting != nil {
			result[userID] = setting
			continue
		}
		result[userID] = &model.PrivacySetting{
			UserID:               userID,
			PublicPoints:         &public,
			PublicFollowerList:   &public,
			PublicFollowingList:  &public,
			PublicArticleList:    &public,
			PublicCommentList:    &public,
			PublicOnlineStatus:   &public,
			PublicLocation:       &public,
			PublicMoonbreezeList: &public,
		}
	}
	return result, nil
}

func (s *PrivacySettingUsecase) UpsertByUserID(ctx context.Context, setting *model.PrivacySetting) (*model.PrivacySetting, error) {
	return s.privacySettingRepo.UpsertByUserID(ctx, setting)
}
