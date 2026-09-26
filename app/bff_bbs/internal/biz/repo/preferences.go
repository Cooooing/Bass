package repo

import "context"

type PreferencesClient interface {
	GetCurrentPreferences(ctx context.Context, userID int64) (*Preference, error)
	UpdateCurrentPreferences(ctx context.Context, req *UpdateCurrentPreferencesReq) (*Preference, error)
}

type Preference struct {
	UserID      int64
	Language    int32
	Timezone    *string
	Theme       *string
	MobileTheme *string
}

type UpdateCurrentPreferencesReq struct {
	UserID      int64
	Timezone    *string
	Theme       *string
	MobileTheme *string
	Language    *int32
}
