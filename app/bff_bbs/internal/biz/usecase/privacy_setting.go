package usecase

import (
	"bff_bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bff_bbs/v1/user"
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
			UserId:               row.UserID,
			PublicPoints:         row.PublicPoints,
			PublicFollowerList:   row.PublicFollowerList,
			PublicFollowingList:  row.PublicFollowingList,
			PublicArticleList:    row.PublicArticleList,
			PublicCommentList:    row.PublicCommentList,
			PublicOnlineStatus:   row.PublicOnlineStatus,
			PublicLocation:       row.PublicLocation,
			PublicMoonbreezeList: row.PublicMoonbreezeList,
		}
	}
	return setting, nil
}

type UpdateCurrentPrivacySettingReq struct {
	UserID               int64
	PublicPoints         *bool
	PublicFollowerList   *bool
	PublicFollowingList  *bool
	PublicArticleList    *bool
	PublicCommentList    *bool
	PublicOnlineStatus   *bool
	PublicLocation       *bool
	PublicMoonbreezeList *bool
}

func (u *PrivacySettingUsecase) UpdateCurrentPrivacySetting(
	ctx context.Context,
	req *UpdateCurrentPrivacySettingReq,
) (*bbsuserv1.UpdateCurrentPrivacySetting_Resp_PrivacySetting, error) {
	reply, err := u.privacySettingClient.UpdateCurrentPrivacySetting(ctx, &repo.UpdateCurrentPrivacySettingReq{
		UserID:               req.UserID,
		PublicPoints:         req.PublicPoints,
		PublicFollowerList:   req.PublicFollowerList,
		PublicFollowingList:  req.PublicFollowingList,
		PublicArticleList:    req.PublicArticleList,
		PublicCommentList:    req.PublicCommentList,
		PublicOnlineStatus:   req.PublicOnlineStatus,
		PublicLocation:       req.PublicLocation,
		PublicMoonbreezeList: req.PublicMoonbreezeList,
	})
	if err != nil {
		return nil, err
	}
	var setting *bbsuserv1.UpdateCurrentPrivacySetting_Resp_PrivacySetting
	if row := reply; row != nil {
		setting = &bbsuserv1.UpdateCurrentPrivacySetting_Resp_PrivacySetting{
			UserId:               row.UserID,
			PublicPoints:         row.PublicPoints,
			PublicFollowerList:   row.PublicFollowerList,
			PublicFollowingList:  row.PublicFollowingList,
			PublicArticleList:    row.PublicArticleList,
			PublicCommentList:    row.PublicCommentList,
			PublicOnlineStatus:   row.PublicOnlineStatus,
			PublicLocation:       row.PublicLocation,
			PublicMoonbreezeList: row.PublicMoonbreezeList,
		}
	}
	return setting, nil
}
