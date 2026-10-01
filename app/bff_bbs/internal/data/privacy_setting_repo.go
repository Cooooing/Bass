package data

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.PrivacySettingClient = (*PrivacySettingClient)(nil)

type PrivacySettingClient struct {
	userClient *rpc.UserClient
}

func NewPrivacySettingClient(
	userClient *rpc.UserClient,
) repo.PrivacySettingClient {
	return &PrivacySettingClient{
		userClient: userClient,
	}
}

func (r *PrivacySettingClient) GetCurrentPrivacySetting(
	ctx context.Context,
	userID int64,
) (*repo.PrivacySetting, error) {
	reply, err := r.userClient.PrivacySetting.Get(ctx, &userv1.GetPrivacySetting_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	return r.privacySettingRepoModel(reply.GetPrivacySetting()), nil
}

func (r *PrivacySettingClient) MapPrivacySettings(
	ctx context.Context,
	userIDs []int64,
) (map[int64]*repo.PrivacySetting, error) {
	if len(userIDs) == 0 {
		return map[int64]*repo.PrivacySetting{}, nil
	}
	reply, err := r.userClient.PrivacySetting.Map(ctx, &userv1.MapPrivacySettings_Req{UserIds: userIDs})
	if err != nil {
		return nil, err
	}
	settings := make(map[int64]*repo.PrivacySetting, len(reply.GetPrivacySettings()))
	for userID, setting := range reply.GetPrivacySettings() {
		settings[userID] = r.privacySettingRepoModel(setting)
	}
	return settings, nil
}

func (r *PrivacySettingClient) UpdateCurrentPrivacySetting(
	ctx context.Context,
	req *repo.UpdateCurrentPrivacySettingReq,
) (*repo.PrivacySetting, error) {
	reply, err := r.userClient.PrivacySetting.Update(ctx, &userv1.UpdatePrivacySetting_Req{
		UserId:               req.UserID,
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
	return r.privacySettingRepoModel(reply.GetPrivacySetting()), nil
}

func (*PrivacySettingClient) privacySettingRepoModel(setting *userv1.PrivacySetting) *repo.PrivacySetting {
	if setting == nil {
		return nil
	}
	return &repo.PrivacySetting{
		UserID:               setting.GetUserId(),
		PublicPoints:         setting.PublicPoints,
		PublicFollowerList:   setting.PublicFollowerList,
		PublicFollowingList:  setting.PublicFollowingList,
		PublicArticleList:    setting.PublicArticleList,
		PublicCommentList:    setting.PublicCommentList,
		PublicOnlineStatus:   setting.PublicOnlineStatus,
		PublicLocation:       setting.PublicLocation,
		PublicMoonbreezeList: setting.PublicMoonbreezeList,
	}
}
