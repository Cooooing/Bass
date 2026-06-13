package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.PrivacySettingClient = (*PrivacySettingClient)(nil)

type PrivacySettingClient struct {
	userClient *rpc.UserClient
}

func NewPrivacySettingClient(userClient *rpc.UserClient) repo.PrivacySettingClient {
	return &PrivacySettingClient{userClient: userClient}
}

func (r *PrivacySettingClient) GetCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.PrivacySetting.Get(ctx, &userv1.GetPrivacySetting_Request{UserId: userID})
	if err != nil {
		return nil, err
	}
	setting := reply.GetPrivacySetting()
	var out *bbsuserv1.PrivacySetting
	if setting != nil {
		out = &bbsuserv1.PrivacySetting{
			UserId:             setting.GetUserId(),
			PublicPoints:       setting.PublicPoints,
			PublicFollowers:    setting.PublicFollowers,
			PublicArticles:     setting.PublicArticles,
			PublicComments:     setting.PublicComments,
			PublicOnlineStatus: setting.PublicOnlineStatus,
			PublicLocation:     setting.PublicLocation,
		}
	}
	return &bbsuserv1.GetCurrentPrivacySetting_Reply{PrivacySetting: out}, nil
}

func (r *PrivacySettingClient) UpdateCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Request) (*bbsuserv1.UpdateCurrentPrivacySetting_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.PrivacySetting.Update(ctx, &userv1.UpdatePrivacySetting_Request{
		UserId:             userID,
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
	var out *bbsuserv1.PrivacySetting
	if setting != nil {
		out = &bbsuserv1.PrivacySetting{
			UserId:             setting.GetUserId(),
			PublicPoints:       setting.PublicPoints,
			PublicFollowers:    setting.PublicFollowers,
			PublicArticles:     setting.PublicArticles,
			PublicComments:     setting.PublicComments,
			PublicOnlineStatus: setting.PublicOnlineStatus,
			PublicLocation:     setting.PublicLocation,
		}
	}
	return &bbsuserv1.UpdateCurrentPrivacySetting_Reply{PrivacySetting: out}, nil
}
