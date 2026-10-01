package repo

import "context"

type PrivacySettingClient interface {
	GetCurrentPrivacySetting(ctx context.Context, userID int64) (*PrivacySetting, error)
	MapPrivacySettings(ctx context.Context, userIDs []int64) (map[int64]*PrivacySetting, error)
	UpdateCurrentPrivacySetting(ctx context.Context, req *UpdateCurrentPrivacySettingReq) (*PrivacySetting, error)
}

type PrivacySetting struct {
	UserID             int64
	PublicPoints       *bool
	PublicFollowers    *bool
	PublicFollowing    *bool
	PublicArticles     *bool
	PublicComments     *bool
	PublicOnlineStatus *bool
	PublicLocation     *bool
	PublicBreezemoons  *bool
}

type UpdateCurrentPrivacySettingReq struct {
	UserID             int64
	PublicPoints       *bool
	PublicFollowers    *bool
	PublicFollowing    *bool
	PublicArticles     *bool
	PublicComments     *bool
	PublicOnlineStatus *bool
	PublicLocation     *bool
	PublicBreezemoons  *bool
}
