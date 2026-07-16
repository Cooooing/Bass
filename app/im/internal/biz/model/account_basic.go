package model

import "time"

type AccountBasic struct {
	ID            int64
	Name          string
	Nickname      *string
	URL           *string
	AvatarURL     *string
	Introduction  *string
	MBTI          string
	Status        string
	FollowCount   *int32
	FollowerCount *int32
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
