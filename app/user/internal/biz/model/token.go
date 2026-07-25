package model

import (
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"time"
	"user/internal/enum"
)

type Token struct {
	Type      enum.TokenType        `json:"typ"`
	UserID    int64                 `json:"user_id"`
	SessionID string                `json:"sid"`
	Realm     commonenum.LoginRealm `json:"realm"`
	Name      string                `json:"name,omitempty"`
	Nickname  string                `json:"nickname,omitempty"`
	Language  commonenums.Language  `json:"language,omitempty"`
	Timezone  string                `json:"timezone,omitempty"`
	JTI       string                `json:"jti,omitempty"`
}

type TokenVerityCodeAccount struct {
	Account string `json:"account"`
}

type RefreshSession struct {
	SessionID            string
	UserID               int64
	Realm                commonenum.LoginRealm
	CurrentJTI           string
	CreatedAtUnix        int64
	LastSeenAtUnix       int64
	SessionExpiresAtUnix int64
	Client               LoginContext
}

type TokenPair struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
	SessionExpiresAt      time.Time
	SessionID             string
}
