package data

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.PreferencesRepo = (*PreferencesRepo)(nil)

type PreferencesRepo struct {
	userClient *rpc.UserClient
}

func NewPreferencesRepo(userClient *rpc.UserClient) repo.PreferencesRepo {
	return &PreferencesRepo{userClient: userClient}
}

func (r *PreferencesRepo) GetCurrentPreferences(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
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

func (r *PreferencesRepo) UpdateCurrentPreferences(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
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
