package model

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
