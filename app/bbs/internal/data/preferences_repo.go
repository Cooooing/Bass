package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	commonenums "common/proto/gen/common/enums"
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

func (r *PreferencesClient) GetCurrentPreferences(ctx context.Context, req *repo.GetCurrentPreferencesReq) (*repo.GetCurrentPreferencesResponse, error) {
	reply, err := r.userClient.Preferences.Get(ctx, &userv1.GetPreferences_Request{UserId: req.UserID})
	if err != nil {
		return nil, err
	}
	preferences := reply.GetPreferences()
	var out *repo.Preference
	if preferences != nil {
		out = &repo.Preference{
			UserID:      preferences.GetUserId(),
			Language:    int32(preferences.GetLanguage()),
			Timezone:    preferences.Timezone,
			Theme:       preferences.Theme,
			MobileTheme: preferences.MobileTheme,
		}
	}
	return &repo.GetCurrentPreferencesResponse{Preference: out}, nil
}

func (r *PreferencesClient) UpdateCurrentPreferences(ctx context.Context, req *repo.UpdateCurrentPreferencesReq) (*repo.UpdateCurrentPreferencesResponse, error) {
	updateReq := &userv1.UpdatePreferences_Request{
		UserId:      req.UserID,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
	}
	if req.Language != nil {
		language := commonenums.Language(*req.Language)
		updateReq.Language = &language
	}
	reply, err := r.userClient.Preferences.Update(ctx, updateReq)
	if err != nil {
		return nil, err
	}
	preferences := reply.GetPreferences()
	var out *repo.Preference
	if preferences != nil {
		out = &repo.Preference{
			UserID:      preferences.GetUserId(),
			Language:    int32(preferences.GetLanguage()),
			Timezone:    preferences.Timezone,
			Theme:       preferences.Theme,
			MobileTheme: preferences.MobileTheme,
		}
	}
	return &repo.UpdateCurrentPreferencesResponse{Preference: out}, nil
}
