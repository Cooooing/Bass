package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	commonenums "common/proto/gen/common/enums"
	"context"
)

type PreferencesUsecase struct {
	preferencesClient repo.PreferencesClient
}

func NewPreferencesUsecase(preferencesClient repo.PreferencesClient) *PreferencesUsecase {
	return &PreferencesUsecase{preferencesClient: preferencesClient}
}

type GetCurrentPreferencesReq struct {
	UserID int64
}

type GetCurrentPreferencesResponse struct {
	Preference *bbsuserv1.GetCurrentPreferences_Response_Preference
}

func (u *PreferencesUsecase) GetCurrentPreferences(ctx context.Context, req *GetCurrentPreferencesReq) (*GetCurrentPreferencesResponse, error) {
	reply, err := u.preferencesClient.GetCurrentPreferences(ctx, &repo.GetCurrentPreferencesReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var preference *bbsuserv1.GetCurrentPreferences_Response_Preference
	if row := reply.Preference; row != nil {
		preference = &bbsuserv1.GetCurrentPreferences_Response_Preference{
			UserId:      row.UserID,
			Language:    commonenums.Language(row.Language),
			Timezone:    row.Timezone,
			Theme:       row.Theme,
			MobileTheme: row.MobileTheme,
		}
	}
	return &GetCurrentPreferencesResponse{Preference: preference}, nil
}

type UpdateCurrentPreferencesReq struct {
	UserID      int64
	Timezone    *string
	Theme       *string
	MobileTheme *string
	Language    *commonenums.Language
}

type UpdateCurrentPreferencesResponse struct {
	Preference *bbsuserv1.UpdateCurrentPreferences_Response_Preference
}

func (u *PreferencesUsecase) UpdateCurrentPreferences(ctx context.Context, req *UpdateCurrentPreferencesReq) (*UpdateCurrentPreferencesResponse, error) {
	var language *int32
	if req.Language != nil {
		value := int32(*req.Language)
		language = &value
	}
	reply, err := u.preferencesClient.UpdateCurrentPreferences(ctx, &repo.UpdateCurrentPreferencesReq{
		UserID:      req.UserID,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
		Language:    language,
	})
	if err != nil {
		return nil, err
	}
	var preference *bbsuserv1.UpdateCurrentPreferences_Response_Preference
	if row := reply.Preference; row != nil {
		preference = &bbsuserv1.UpdateCurrentPreferences_Response_Preference{
			UserId:      row.UserID,
			Language:    commonenums.Language(row.Language),
			Timezone:    row.Timezone,
			Theme:       row.Theme,
			MobileTheme: row.MobileTheme,
		}
	}
	return &UpdateCurrentPreferencesResponse{Preference: preference}, nil
}
