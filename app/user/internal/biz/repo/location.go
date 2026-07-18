package repo

import (
	"context"
	"user/internal/biz/model"
)

type LocationRepo interface {
	Get(ctx context.Context, req *LocationGetReq) (*model.Location, error)
	List(ctx context.Context, req *LocationGetReq) ([]*model.Location, error)
	Map(ctx context.Context, req *LocationGetReq) (map[int64]*model.Location, error)
	Count(ctx context.Context, req *LocationGetReq) (int, error)
	Page(ctx context.Context, req *LocationPageReq) (*LocationPageResp, error)
	UpsertByUserID(ctx context.Context, location *model.Location) (*model.Location, error)
	Update(ctx context.Context, location *model.Location) (*model.Location, error)
}

type LocationGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type LocationPageReq struct {
	Page  PageReq
	Query LocationGetReq
}

type LocationPageResp struct {
	Rows []*model.Location
	Page PageResp
}
