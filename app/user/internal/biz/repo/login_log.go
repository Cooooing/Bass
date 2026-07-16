package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type LoginLogRepo interface {
	Create(ctx context.Context, req *LoginLogCreateReq) (*LoginLogCreateResponse, error)
	Get(ctx context.Context, req *LoginLogGetReq) (*LoginLogGetResponse, error)
	List(ctx context.Context, req *LoginLogGetReq) (*LoginLogListResponse, error)
	Map(ctx context.Context, req *LoginLogGetReq) (*LoginLogMapResponse, error)
	Count(ctx context.Context, req *LoginLogGetReq) (*LoginLogCountResponse, error)
	Page(ctx context.Context, req *LoginLogPageReq) (*LoginLogPageResponse, error)
}

type LoginLogCreateReq struct {
	Log *model.LoginLog
}

type LoginLogCreateResponse struct {
	Log *model.LoginLog
}

type LoginLogGetReq struct {
	ID          *int64
	IDs         []int64
	UserID      *int64
	UserIDs     []int64
	Status      *enum.LoginStatus
	IP          *string
	LastSuccess bool
}

type LoginLogGetResponse struct {
	Log *model.LoginLog
}

type LoginLogListResponse struct {
	Rows []*model.LoginLog
}

type LoginLogMapResponse struct {
	Rows map[int64]*model.LoginLog
}

type LoginLogCountResponse struct {
	Count int
}

type LoginLogPageReq struct {
	Page  PageReq
	Query LoginLogGetReq
}

type LoginLogPageResponse struct {
	Rows []*model.LoginLog
	Page PageResponse
}
