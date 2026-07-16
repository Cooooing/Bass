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

type GetPrivacySettingByUserIDReq struct {
	UserID int64
}

type GetPrivacySettingByUserIDResponse struct {
	PrivacySetting *model.PrivacySetting
}

func (s *PrivacySettingUsecase) GetByUserID(ctx context.Context, req *GetPrivacySettingByUserIDReq) (*GetPrivacySettingByUserIDResponse, error) {
	privacySettingResp, err := s.privacySettingRepo.Get(ctx, &repo.PrivacySettingGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	return &GetPrivacySettingByUserIDResponse{PrivacySetting: privacySettingResp.PrivacySetting}, nil
}

type UpsertPrivacySettingByUserIDReq struct {
	PrivacySetting *model.PrivacySetting
}

type UpsertPrivacySettingByUserIDResponse struct {
	PrivacySetting *model.PrivacySetting
}

func (s *PrivacySettingUsecase) UpsertByUserID(ctx context.Context, req *UpsertPrivacySettingByUserIDReq) (*UpsertPrivacySettingByUserIDResponse, error) {
	privacySettingResp, err := s.privacySettingRepo.UpsertByUserID(ctx, &repo.PrivacySettingUpsertByUserIDReq{PrivacySetting: req.PrivacySetting})
	if err != nil {
		return nil, err
	}
	return &UpsertPrivacySettingByUserIDResponse{PrivacySetting: privacySettingResp.PrivacySetting}, nil
}
