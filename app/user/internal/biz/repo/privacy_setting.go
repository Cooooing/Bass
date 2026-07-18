package repo

import (
	"context"
	"user/internal/biz/model"
)

type PrivacySettingRepo interface {
	Get(ctx context.Context, req *PrivacySettingGetReq) (*model.PrivacySetting, error)
	List(ctx context.Context, req *PrivacySettingGetReq) ([]*model.PrivacySetting, error)
	Map(ctx context.Context, req *PrivacySettingGetReq) (map[int64]*model.PrivacySetting, error)
	Count(ctx context.Context, req *PrivacySettingGetReq) (int, error)
	Page(ctx context.Context, req *PrivacySettingPageReq) (*PrivacySettingPageResp, error)
	UpsertByUserID(ctx context.Context, setting *model.PrivacySetting) (*model.PrivacySetting, error)
	Update(ctx context.Context, setting *model.PrivacySetting) (*model.PrivacySetting, error)
}

type PrivacySettingGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type PrivacySettingPageReq struct {
	Page  PageReq
	Query PrivacySettingGetReq
}

type PrivacySettingPageResp struct {
	Rows []*model.PrivacySetting
	Page PageResp
}
