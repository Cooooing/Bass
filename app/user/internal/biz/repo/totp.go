package repo

import (
	"context"
	"user/internal/biz/model"
)

type TotpRepo interface {
	Get(ctx context.Context, req *TotpGetReq) (*TotpGetResponse, error)
	List(ctx context.Context, req *TotpGetReq) (*TotpListResponse, error)
	Map(ctx context.Context, req *TotpGetReq) (*TotpMapResponse, error)
	Count(ctx context.Context, req *TotpGetReq) (*TotpCountResponse, error)
	Page(ctx context.Context, req *TotpPageReq) (*TotpPageResponse, error)
	UpsertEnabledByUserID(ctx context.Context, req *TotpUpsertEnabledByUserIDReq) (*TotpUpsertEnabledByUserIDResponse, error)
	DisableByUserID(ctx context.Context, req *TotpDisableByUserIDReq) (*TotpDisableByUserIDResponse, error)
}

type TotpGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
	Enable  *bool
}

type TotpGetResponse struct {
	Totp *model.Totp
}

type TotpListResponse struct {
	Rows []*model.Totp
}

type TotpMapResponse struct {
	Rows map[int64]*model.Totp
}

type TotpCountResponse struct {
	Count int
}

type TotpPageReq struct {
	Page  PageReq
	Query TotpGetReq
}

type TotpPageResponse struct {
	Rows []*model.Totp
	Page PageResponse
}

type TotpUpsertEnabledByUserIDReq struct {
	UserID int64
	Secret string
}

type TotpUpsertEnabledByUserIDResponse struct {
	Totp *model.Totp
}

type TotpDisableByUserIDReq struct {
	UserID int64
}

type TotpDisableByUserIDResponse struct {
	Totp *model.Totp
}
