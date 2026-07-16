package repo

import "context"

type PreferencesClient interface {
	GetCurrentPreferences(ctx context.Context, req *GetCurrentPreferencesReq) (*GetCurrentPreferencesResponse, error)
	UpdateCurrentPreferences(ctx context.Context, req *UpdateCurrentPreferencesReq) (*UpdateCurrentPreferencesResponse, error)
}

type Preference struct {
	UserID      int64
	Language    int32
	Timezone    *string
	Theme       *string
	MobileTheme *string
}

type GetCurrentPreferencesReq struct {
	UserID int64
}

type GetCurrentPreferencesResponse struct {
	Preference *Preference
}

type UpdateCurrentPreferencesReq struct {
	UserID      int64
	Timezone    *string
	Theme       *string
	MobileTheme *string
	Language    *int32
}

type UpdateCurrentPreferencesResponse struct {
	Preference *Preference
}
