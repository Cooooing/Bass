package usecase

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type PreferencesUsecase struct {
	preferencesRepo repo.PreferencesRepo
}

func NewPreferencesUsecase(preferencesRepo repo.PreferencesRepo) *PreferencesUsecase {
	return &PreferencesUsecase{preferencesRepo: preferencesRepo}
}

type GetPreferencesByUserIDReq struct {
	UserID int64
}

type GetPreferencesByUserIDResponse struct {
	Preferences *model.Preferences
}

func (s *PreferencesUsecase) GetByUserID(ctx context.Context, req *GetPreferencesByUserIDReq) (*GetPreferencesByUserIDResponse, error) {
	preferencesResp, err := s.preferencesRepo.Get(ctx, &repo.PreferencesGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	return &GetPreferencesByUserIDResponse{Preferences: preferencesResp.Preferences}, nil
}

type UpsertPreferencesByUserIDReq struct {
	Preferences *model.Preferences
}

type UpsertPreferencesByUserIDResponse struct {
	Preferences *model.Preferences
}

func (s *PreferencesUsecase) UpsertByUserID(ctx context.Context, req *UpsertPreferencesByUserIDReq) (*UpsertPreferencesByUserIDResponse, error) {
	preferencesResp, err := s.preferencesRepo.UpsertByUserID(ctx, &repo.PreferencesUpsertByUserIDReq{Preferences: req.Preferences})
	if err != nil {
		return nil, err
	}
	return &UpsertPreferencesByUserIDResponse{Preferences: preferencesResp.Preferences}, nil
}
