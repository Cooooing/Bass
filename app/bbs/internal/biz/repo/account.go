package repo

import (
	"context"
	"time"
)

type AccountClient interface {
	GetCurrentAccount(ctx context.Context, userID int64) (*Account, error)
	GetProfileAccount(ctx context.Context, userID int64) (*AccountProfile, error)
	UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*AccountProfile, error)
	AvatarAccount(ctx context.Context, name string) (*AvatarAccountResp, error)
}

type AccountProfile struct {
	ID            int64
	Name          string
	Nickname      *string
	URL           *string
	AvatarURL     *string
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
	UserID       int64
	AvatarURL    *string
	Nickname     *string
	URL          *string
	Introduction *string
	MBTI         *int32
}
type AvatarAccountResp struct {
	Data        []byte
	ContentType string
}
