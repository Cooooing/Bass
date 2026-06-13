package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.PreferencesClient = (*PreferencesClient)(nil)

type PreferencesClient struct {
	userClient *rpc.UserClient
}

func NewPreferencesClient(userClient *rpc.UserClient) repo.PreferencesClient {
	return &PreferencesClient{userClient: userClient}
}

func (r *PreferencesClient) GetCurrentPreferences(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Preferences.Get(ctx, &userv1.GetPreferences_Request{UserId: userID})
	if err != nil {
		return nil, err
	}
	preferences := reply.GetPreferences()
	var out *bbsuserv1.Preference
	if preferences != nil {
		out = &bbsuserv1.Preference{
			UserId:      preferences.GetUserId(),
			Language:    preferences.Language,
			Timezone:    preferences.Timezone,
			Theme:       preferences.Theme,
			MobileTheme: preferences.MobileTheme,
		}
	}
	return &bbsuserv1.GetCurrentPreferences_Reply{Preference: out}, nil
}

func (r *PreferencesClient) UpdateCurrentPreferences(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Preferences.Update(ctx, &userv1.UpdatePreferences_Request{
		UserId:      userID,
		Language:    req.Language,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
	})
	if err != nil {
		return nil, err
	}
	preferences := reply.GetPreferences()
	var out *bbsuserv1.Preference
	if preferences != nil {
		out = &bbsuserv1.Preference{
			UserId:      preferences.GetUserId(),
			Language:    preferences.Language,
			Timezone:    preferences.Timezone,
			Theme:       preferences.Theme,
			MobileTheme: preferences.MobileTheme,
		}
	}
	return &bbsuserv1.UpdateCurrentPreferences_Reply{Preference: out}, nil
}
