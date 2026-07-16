package repo

import "context"

type AccountClient interface {
	GetCurrentAccount(ctx context.Context, req *GetCurrentAccountReq) (*GetCurrentAccountResponse, error)
	GetProfileAccount(ctx context.Context, req *GetProfileAccountReq) (*GetProfileAccountResponse, error)
	UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*UpdateProfileAccountResponse, error)
	AvatarAccount(ctx context.Context, req *AvatarAccountReq) (*AvatarAccountResponse, error)
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
	CreatedAt     string
	UpdatedAt     string
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

type GetCurrentAccountReq struct {
	UserID int64
}

type GetCurrentAccountResponse struct {
	Account *Account
}

type GetProfileAccountReq struct {
	UserID int64
}

type GetProfileAccountResponse struct {
	Profile *AccountProfile
}

type UpdateProfileAccountReq struct {
	UserID       int64
	AvatarURL    *string
	Nickname     *string
	URL          *string
	Introduction *string
	MBTI         *int32
}

type UpdateProfileAccountResponse struct {
	Profile *AccountProfile
}

type AvatarAccountReq struct {
	Name string
}

type AvatarAccountResponse struct {
	Data        []byte
	ContentType string
}
