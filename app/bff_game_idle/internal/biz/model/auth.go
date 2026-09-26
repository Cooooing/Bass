package model

import "time"

type LoginToken struct {
	AccessToken           string     `json:"access_token"`
	RefreshToken          string     `json:"refresh_token"`
	AccessTokenExpiresAt  *time.Time `json:"access_token_expires_at,omitempty"`
	RefreshTokenExpiresAt *time.Time `json:"refresh_token_expires_at,omitempty"`
	SessionExpiresAt      *time.Time `json:"session_expires_at,omitempty"`
	UserID                int64      `json:"user_id"`
	Name                  string     `json:"name"`
	Nickname              string     `json:"nickname"`
}
