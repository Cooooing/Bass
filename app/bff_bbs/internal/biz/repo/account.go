package repo

import (
	"context"
	"time"

	"bff_bbs/internal/biz/model"
)

type AccountClient interface {
	GetCurrentAccount(ctx context.Context, userID int64) (*model.Account, error)
	GetProfile(ctx context.Context, name string) (*model.AccountProfile, *time.Time, error)
	MapProfileAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.AccountProfile, error)
	UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*model.AccountProfile, error)
	UpdatePasswordAccount(ctx context.Context, req *UpdatePasswordAccountReq) error
	UpdateEmailAccount(ctx context.Context, req *UpdateEmailAccountReq) error
	UpdatePhoneAccount(ctx context.Context, req *UpdatePhoneAccountReq) error
	AvatarAccount(ctx context.Context, name string) (*AvatarAccountResp, error)
}

type UpdateProfileAccountReq struct {
	UserID            int64
	AvatarAssetID     *int64
	BackgroundAssetID *int64
	Nickname          *string
	URL               *string
	Introduction      *string
	MBTI              *string
}

type UpdatePasswordAccountReq struct {
	UserID      int64
	OldPassword string
	NewPassword string
}

type UpdateEmailAccountReq struct {
	UserID int64
	Email  string
	Code   string
}

type UpdatePhoneAccountReq struct {
	UserID int64
	Phone  string
	Code   string
}
type AvatarAccountResp struct {
	Data        []byte
	ContentType string
}
