package repo

import "context"

type PrivacySettingClient interface {
	GetCurrentPrivacySetting(ctx context.Context, req *GetCurrentPrivacySettingReq) (*GetCurrentPrivacySettingResponse, error)
	UpdateCurrentPrivacySetting(ctx context.Context, req *UpdateCurrentPrivacySettingReq) (*UpdateCurrentPrivacySettingResponse, error)
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

type GetCurrentPrivacySettingReq struct {
	UserID int64
}

type GetCurrentPrivacySettingResponse struct {
	PrivacySetting *PrivacySetting
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

type UpdateCurrentPrivacySettingResponse struct {
	PrivacySetting *PrivacySetting
}
