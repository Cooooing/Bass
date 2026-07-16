package repo

import (
	"context"
	"user/internal/biz/model"
)

type PrivacySettingRepo interface {
	Get(ctx context.Context, req *PrivacySettingGetReq) (*PrivacySettingGetResponse, error)
	List(ctx context.Context, req *PrivacySettingGetReq) (*PrivacySettingListResponse, error)
	Map(ctx context.Context, req *PrivacySettingGetReq) (*PrivacySettingMapResponse, error)
	Count(ctx context.Context, req *PrivacySettingGetReq) (*PrivacySettingCountResponse, error)
	Page(ctx context.Context, req *PrivacySettingPageReq) (*PrivacySettingPageResponse, error)
	UpsertByUserID(ctx context.Context, req *PrivacySettingUpsertByUserIDReq) (*PrivacySettingUpsertByUserIDResponse, error)
	Update(ctx context.Context, req *PrivacySettingUpdateReq) (*PrivacySettingUpdateResponse, error)
}

type PrivacySettingGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type PrivacySettingGetResponse struct {
	PrivacySetting *model.PrivacySetting
}

type PrivacySettingListResponse struct {
	Rows []*model.PrivacySetting
}

type PrivacySettingMapResponse struct {
	Rows map[int64]*model.PrivacySetting
}

type PrivacySettingCountResponse struct {
	Count int
}

type PrivacySettingPageReq struct {
	Page  PageReq
	Query PrivacySettingGetReq
}

type PrivacySettingPageResponse struct {
	Rows []*model.PrivacySetting
	Page PageResponse
}

type PrivacySettingUpsertByUserIDReq struct {
	PrivacySetting *model.PrivacySetting
}

type PrivacySettingUpsertByUserIDResponse struct {
	PrivacySetting *model.PrivacySetting
}

type PrivacySettingUpdateReq struct {
	PrivacySetting *model.PrivacySetting
}

type PrivacySettingUpdateResponse struct {
	PrivacySetting *model.PrivacySetting
}
