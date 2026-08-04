package repo

import (
	"context"
	"time"
)

type AccountClient interface {
	GetCurrentAccount(ctx context.Context, userID int64) (*Account, error)
	GetProfileAccount(ctx context.Context, userID int64) (*AccountProfile, error)
	UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*AccountProfile, error)
	UpdatePasswordAccount(ctx context.Context, req *UpdatePasswordAccountReq) error
	UpdateEmailAccount(ctx context.Context, req *UpdateEmailAccountReq) error
	UpdatePhoneAccount(ctx context.Context, req *UpdatePhoneAccountReq) error
	AvatarAccount(ctx context.Context, name string) (*AvatarAccountResp, error)
}

type AccountProfile struct {
	ID            int64
	Name          string
	Nickname      *string
	URL           *string
	AvatarURL     *string
	AvatarAssetID *int64
	Introduction  *string
	Status        int32
	MBTI          int32
	FollowCount   *int32
	FollowerCount *int32
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type AccountContact struct {
	UserID int64
	Email  *string
	Phone  *string
}

type Account struct {
	Profile *AccountProfile
	Contact *AccountContact
}

type UpdateProfileAccountReq struct {
	UserID        int64
	AvatarAssetID *int64
	Nickname      *string
	URL           *string
	Introduction  *string
	MBTI          *int32
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
