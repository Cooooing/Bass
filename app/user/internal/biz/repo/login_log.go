package repo

import (
	"common/api/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type LoginLogRepo interface {
	Create(ctx context.Context, log *model.LoginLog) (*model.LoginLog, error)
	FindLastSuccessByUserID(ctx context.Context, userID int64) (*model.LoginLog, error)
	List(ctx context.Context, req *LoginLogGetReq) ([]*model.LoginLog, error)
	Page(ctx context.Context, page *common.PageRequest, req *LoginLogGetReq) ([]*model.LoginLog, *common.PageReply, error)
}

type LoginLogGetReq struct {
	UserID  *int64
	Account *string
	Status  *enum.LoginStatus
	IP      *string
}
