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

type GetCurrentPrivacySettingReq struct {
	UserID int64
}

type GetCurrentPrivacySettingResponse struct {
	PrivacySetting *bbsuserv1.GetCurrentPrivacySetting_Response_PrivacySetting
}

func (u *PrivacySettingUsecase) GetCurrentPrivacySetting(ctx context.Context, req *GetCurrentPrivacySettingReq) (*GetCurrentPrivacySettingResponse, error) {
	reply, err := u.privacySettingClient.GetCurrentPrivacySetting(ctx, &repo.GetCurrentPrivacySettingReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var setting *bbsuserv1.GetCurrentPrivacySetting_Response_PrivacySetting
	if row := reply.PrivacySetting; row != nil {
		setting = &bbsuserv1.GetCurrentPrivacySetting_Response_PrivacySetting{
			UserId:             row.UserID,
			PublicPoints:       row.PublicPoints,
			PublicFollowers:    row.PublicFollowers,
			PublicArticles:     row.PublicArticles,
			PublicComments:     row.PublicComments,
			PublicOnlineStatus: row.PublicOnlineStatus,
			PublicLocation:     row.PublicLocation,
		}
	}
	return &GetCurrentPrivacySettingResponse{PrivacySetting: setting}, nil
}

type UpdateCurrentPrivacySettingReq struct {
	UserID             int64
	PublicPoints       *bool
	PublicFollowers    *bool
	PublicArticles     *bool
	PublicComments     *bool
	PublicOnlineStatus *bool
	PublicLocation     *bool
}

type UpdateCurrentPrivacySettingResponse struct {
	PrivacySetting *bbsuserv1.UpdateCurrentPrivacySetting_Response_PrivacySetting
}

func (u *PrivacySettingUsecase) UpdateCurrentPrivacySetting(ctx context.Context, req *UpdateCurrentPrivacySettingReq) (*UpdateCurrentPrivacySettingResponse, error) {
	reply, err := u.privacySettingClient.UpdateCurrentPrivacySetting(ctx, &repo.UpdateCurrentPrivacySettingReq{
		UserID:             req.UserID,
		PublicPoints:       req.PublicPoints,
		PublicFollowers:    req.PublicFollowers,
		PublicArticles:     req.PublicArticles,
		PublicComments:     req.PublicComments,
		PublicOnlineStatus: req.PublicOnlineStatus,
		PublicLocation:     req.PublicLocation,
	})
	if err != nil {
		return nil, err
	}
	var setting *bbsuserv1.UpdateCurrentPrivacySetting_Response_PrivacySetting
	if row := reply.PrivacySetting; row != nil {
		setting = &bbsuserv1.UpdateCurrentPrivacySetting_Response_PrivacySetting{
			UserId:             row.UserID,
			PublicPoints:       row.PublicPoints,
			PublicFollowers:    row.PublicFollowers,
			PublicArticles:     row.PublicArticles,
			PublicComments:     row.PublicComments,
			PublicOnlineStatus: row.PublicOnlineStatus,
			PublicLocation:     row.PublicLocation,
		}
	}
	return &UpdateCurrentPrivacySettingResponse{PrivacySetting: setting}, nil
}
