package repo

import (
	"context"
	"user/internal/biz/model"
)

type PreferencesRepo interface {
	Get(ctx context.Context, req *PreferencesGetReq) (*PreferencesGetResponse, error)
	List(ctx context.Context, req *PreferencesGetReq) (*PreferencesListResponse, error)
	Map(ctx context.Context, req *PreferencesGetReq) (*PreferencesMapResponse, error)
	Count(ctx context.Context, req *PreferencesGetReq) (*PreferencesCountResponse, error)
	Page(ctx context.Context, req *PreferencesPageReq) (*PreferencesPageResponse, error)
	UpsertByUserID(ctx context.Context, req *PreferencesUpsertByUserIDReq) (*PreferencesUpsertByUserIDResponse, error)
	Update(ctx context.Context, req *PreferencesUpdateReq) (*PreferencesUpdateResponse, error)
}

type PreferencesGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type PreferencesGetResponse struct {
	Preferences *model.Preferences
}

type PreferencesListResponse struct {
	Rows []*model.Preferences
}

type PreferencesMapResponse struct {
	Rows map[int64]*model.Preferences
}

type PreferencesCountResponse struct {
	Count int
}

type PreferencesPageReq struct {
	Page  PageReq
	Query PreferencesGetReq
}

type PreferencesPageResponse struct {
	Rows []*model.Preferences
	Page PageResponse
}

type PreferencesUpsertByUserIDReq struct {
	Preferences *model.Preferences
}

type PreferencesUpsertByUserIDResponse struct {
	Preferences *model.Preferences
}

type PreferencesUpdateReq struct {
	Preferences *model.Preferences
}

type PreferencesUpdateResponse struct {
	Preferences *model.Preferences
}
