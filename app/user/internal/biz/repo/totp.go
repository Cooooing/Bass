package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
)

type TotpRepo interface {
	Get(ctx context.Context, req *TotpGetReq) (*model.Totp, error)
	List(ctx context.Context, req *TotpGetReq) ([]*model.Totp, error)
	Map(ctx context.Context, req *TotpGetReq) (map[int64]*model.Totp, error)
	Count(ctx context.Context, req *TotpGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *TotpGetReq) ([]*model.Totp, *common.PageReply, error)
	UpsertEnabledByUserID(ctx context.Context, userID int64, secret string) (*model.Totp, error)
	DisableByUserID(ctx context.Context, userID int64) (*model.Totp, error)
}

type TotpGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
	Enable  *bool
}
