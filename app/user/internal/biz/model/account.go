package model

import (
	"time"
	"user/internal/enum"
)

type Account struct {
	ID            int64
	Name          string
	Nickname      *string
	Password      string
	Email         *string
	Phone         *string
	URL           *string
	AvatarAssetID *int64
	Introduction  *string
	Mbti          *enum.MBTI
	Status        *enum.AccountStatus
	FollowCount   *int32
	FollowerCount *int32
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type AccountProfileUpdate struct {
	UserID        int64
	AvatarAssetID *int64
	Nickname      *string
	URL           *string
	Introduction  *string
	Mbti          *enum.MBTI
	ClearMBTI     bool
}

type AccountAvailability struct {
	Name           *string
	Email          *string
	Phone          *string
	NameAvailable  bool
	EmailAvailable bool
	PhoneAvailable bool
}
