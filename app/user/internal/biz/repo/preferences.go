package repo

import (
	"context"
	"user/internal/biz/model"
)

type PreferencesRepo interface {
	Get(ctx context.Context, req *PreferencesGetReq) (*model.Preferences, error)
	List(ctx context.Context, req *PreferencesGetReq) ([]*model.Preferences, error)
	Map(ctx context.Context, req *PreferencesGetReq) (map[int64]*model.Preferences, error)
	Count(ctx context.Context, req *PreferencesGetReq) (int, error)
	Page(ctx context.Context, req *PreferencesPageReq) (*PreferencesPageResp, error)
	UpsertByUserID(ctx context.Context, preferences *model.Preferences) (*model.Preferences, error)
	Update(ctx context.Context, preferences *model.Preferences) (*model.Preferences, error)
}

type PreferencesGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type PreferencesPageReq struct {
	Page  PageReq
	Query PreferencesGetReq
}

type PreferencesPageResp struct {
	Rows []*model.Preferences
	Page PageResp
}
