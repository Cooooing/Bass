package repo

import (
	"context"
	"user/internal/biz/model"
)

type TotpRepo interface {
	Get(ctx context.Context, req *TotpGetReq) (*model.Totp, error)
	List(ctx context.Context, req *TotpGetReq) ([]*model.Totp, error)
	Map(ctx context.Context, req *TotpGetReq) (map[int64]*model.Totp, error)
	Count(ctx context.Context, req *TotpGetReq) (int, error)
	Page(ctx context.Context, req *TotpPageReq) (*TotpPageResp, error)
	UpsertEnabledByUserID(ctx context.Context, req *TotpUpsertEnabledByUserIDReq) (*model.Totp, error)
	DisableByUserID(ctx context.Context, userID int64) (*model.Totp, error)
}

type TotpGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
	Enable  *bool
}

type TotpPageReq struct {
	Page  PageReq
	Query TotpGetReq
}

type TotpPageResp struct {
	Rows []*model.Totp
	Page PageResp
}

type TotpUpsertEnabledByUserIDReq struct {
	UserID int64
	Secret string
}
