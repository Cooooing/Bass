package repo

import (
	"context"
	"user/internal/biz/model"
)

type LocationRepo interface {
	Get(ctx context.Context, req *LocationGetReq) (*LocationGetResponse, error)
	List(ctx context.Context, req *LocationGetReq) (*LocationListResponse, error)
	Map(ctx context.Context, req *LocationGetReq) (*LocationMapResponse, error)
	Count(ctx context.Context, req *LocationGetReq) (*LocationCountResponse, error)
	Page(ctx context.Context, req *LocationPageReq) (*LocationPageResponse, error)
	UpsertByUserID(ctx context.Context, req *LocationUpsertByUserIDReq) (*LocationUpsertByUserIDResponse, error)
	Update(ctx context.Context, req *LocationUpdateReq) (*LocationUpdateResponse, error)
}

type LocationGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type LocationGetResponse struct {
	Location *model.Location
}

type LocationListResponse struct {
	Rows []*model.Location
}

type LocationMapResponse struct {
	Rows map[int64]*model.Location
}

type LocationCountResponse struct {
	Count int
}

type LocationPageReq struct {
	Page  PageReq
	Query LocationGetReq
}

type LocationPageResponse struct {
	Rows []*model.Location
	Page PageResponse
}

type LocationUpsertByUserIDReq struct {
	Location *model.Location
}

type LocationUpsertByUserIDResponse struct {
	Location *model.Location
}

type LocationUpdateReq struct {
	Location *model.Location
}

type LocationUpdateResponse struct {
	Location *model.Location
}
