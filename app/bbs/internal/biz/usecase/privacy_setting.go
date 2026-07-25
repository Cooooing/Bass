package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"
)

type PrivacySettingUsecase struct {
	privacySettingClient repo.PrivacySettingClient
}

func NewPrivacySettingUsecase(
	privacySettingClient repo.PrivacySettingClient,
) *PrivacySettingUsecase {
	return &PrivacySettingUsecase{
		privacySettingClient: privacySettingClient,
	}
}

func (u *PrivacySettingUsecase) GetCurrentPrivacySetting(
	ctx context.Context,
	userID int64,
) (*bbsuserv1.GetCurrentPrivacySetting_Resp_PrivacySetting, error) {
	reply, err := u.privacySettingClient.GetCurrentPrivacySetting(ctx, userID)
	if err != nil {
		return nil, err
	}
	var setting *bbsuserv1.GetCurrentPrivacySetting_Resp_PrivacySetting
	if row := reply; row != nil {
		setting = &bbsuserv1.GetCurrentPrivacySetting_Resp_PrivacySetting{
			UserId:             row.UserID,
			PublicPoints:       row.PublicPoints,
			PublicFollowers:    row.PublicFollowers,
			PublicArticles:     row.PublicArticles,
			PublicComments:     row.PublicComments,
			PublicOnlineStatus: row.PublicOnlineStatus,
			PublicLocation:     row.PublicLocation,
		}
	}
	return setting, nil
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

func (u *PrivacySettingUsecase) UpdateCurrentPrivacySetting(
	ctx context.Context,
	req *UpdateCurrentPrivacySettingReq,
) (*bbsuserv1.UpdateCurrentPrivacySetting_Resp_PrivacySetting, error) {
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
	var setting *bbsuserv1.UpdateCurrentPrivacySetting_Resp_PrivacySetting
	if row := reply; row != nil {
		setting = &bbsuserv1.UpdateCurrentPrivacySetting_Resp_PrivacySetting{
			UserId:             row.UserID,
			PublicPoints:       row.PublicPoints,
			PublicFollowers:    row.PublicFollowers,
			PublicArticles:     row.PublicArticles,
			PublicComments:     row.PublicComments,
			PublicOnlineStatus: row.PublicOnlineStatus,
			PublicLocation:     row.PublicLocation,
		}
	}
	return setting, nil
}
