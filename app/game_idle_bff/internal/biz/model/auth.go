package model

import "time"

type LoginToken struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *time.Time
	RefreshTokenExpiresAt *time.Time
	SessionExpiresAt      *time.Time
	UserID                int64
	Name                  string
	Nickname              string
}
