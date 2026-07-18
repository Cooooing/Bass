package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type LoginLogRepo interface {
	Create(ctx context.Context, log *model.LoginLog) (*model.LoginLog, error)
	Get(ctx context.Context, req *LoginLogGetReq) (*model.LoginLog, error)
	List(ctx context.Context, req *LoginLogGetReq) ([]*model.LoginLog, error)
	Map(ctx context.Context, req *LoginLogGetReq) (map[int64]*model.LoginLog, error)
	Count(ctx context.Context, req *LoginLogGetReq) (int, error)
	Page(ctx context.Context, req *LoginLogPageReq) (*LoginLogPageResp, error)
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

type LoginLogPageReq struct {
	Page  PageReq
	Query LoginLogGetReq
}

type LoginLogPageResp struct {
	Rows []*model.LoginLog
	Page PageResp
}
