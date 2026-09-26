package repo

import "context"

type PrivacySettingClient interface {
	GetCurrentPrivacySetting(ctx context.Context, userID int64) (*PrivacySetting, error)
	UpdateCurrentPrivacySetting(ctx context.Context, req *UpdateCurrentPrivacySettingReq) (*PrivacySetting, error)
}

type PrivacySetting struct {
	UserID             int64
	PublicPoints       *bool
	PublicFollowers    *bool
	PublicArticles     *bool
	PublicComments     *bool
	PublicOnlineStatus *bool
	PublicLocation     *bool
}

type UpdateCurrentPrivacySettingReq struct {
	UserID             int64
	PublicPoints       *bool
	PublicFollowers    *bool
	PublicArticles     *bool
	PublicComments     *bool
	PublicOnlineStatus *bool
	PublicLocation     *bool
}
