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

func (u *PreferencesUsecase) GetCurrentPreferences(ctx context.Context, userID int64) (*bbsuserv1.GetCurrentPreferences_Resp_Preference, error) {
	reply, err := u.preferencesClient.GetCurrentPreferences(ctx, userID)
	if err != nil {
		return nil, err
	}
	var preference *bbsuserv1.GetCurrentPreferences_Resp_Preference
	if row := reply; row != nil {
		preference = &bbsuserv1.GetCurrentPreferences_Resp_Preference{
			UserId:      row.UserID,
			Language:    commonenums.Language(row.Language),
			Timezone:    row.Timezone,
			Theme:       row.Theme,
			MobileTheme: row.MobileTheme,
		}
	}
	return preference, nil
}

type UpdateCurrentPreferencesReq struct {
	UserID      int64
	Timezone    *string
	Theme       *string
	MobileTheme *string
	Language    *commonenums.Language
}

func (u *PreferencesUsecase) UpdateCurrentPreferences(ctx context.Context, req *UpdateCurrentPreferencesReq) (*bbsuserv1.UpdateCurrentPreferences_Resp_Preference, error) {
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
	var preference *bbsuserv1.UpdateCurrentPreferences_Resp_Preference
	if row := reply; row != nil {
		preference = &bbsuserv1.UpdateCurrentPreferences_Resp_Preference{
			UserId:      row.UserID,
			Language:    commonenums.Language(row.Language),
			Timezone:    row.Timezone,
			Theme:       row.Theme,
			MobileTheme: row.MobileTheme,
		}
	}
	return preference, nil
}
