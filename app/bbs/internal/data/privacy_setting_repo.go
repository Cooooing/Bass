package data

import (
	"bbs/internal/biz/repo"
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

func (r *PrivacySettingClient) GetCurrentPrivacySetting(ctx context.Context, userID int64) (*repo.PrivacySetting, error) {
	reply, err := r.userClient.PrivacySetting.Get(ctx, &userv1.GetPrivacySetting_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	setting := reply.GetPrivacySetting()
	var out *repo.PrivacySetting
	if setting != nil {
		out = &repo.PrivacySetting{
			UserID:             setting.GetUserId(),
			PublicPoints:       setting.PublicPoints,
			PublicFollowers:    setting.PublicFollowers,
			PublicArticles:     setting.PublicArticles,
			PublicComments:     setting.PublicComments,
			PublicOnlineStatus: setting.PublicOnlineStatus,
			PublicLocation:     setting.PublicLocation,
		}
	}
	return out, nil
}

func (r *PrivacySettingClient) UpdateCurrentPrivacySetting(ctx context.Context, req *repo.UpdateCurrentPrivacySettingReq) (*repo.PrivacySetting, error) {
	reply, err := r.userClient.PrivacySetting.Update(ctx, &userv1.UpdatePrivacySetting_Req{
		UserId:             req.UserID,
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
	setting := reply.GetPrivacySetting()
	var out *repo.PrivacySetting
	if setting != nil {
		out = &repo.PrivacySetting{
			UserID:             setting.GetUserId(),
			PublicPoints:       setting.PublicPoints,
			PublicFollowers:    setting.PublicFollowers,
			PublicArticles:     setting.PublicArticles,
			PublicComments:     setting.PublicComments,
			PublicOnlineStatus: setting.PublicOnlineStatus,
			PublicLocation:     setting.PublicLocation,
		}
	}
	return out, nil
}
