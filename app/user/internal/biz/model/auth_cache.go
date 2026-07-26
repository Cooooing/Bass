package model

import (
	"time"
	"user/internal/enum"
)

type VerificationCode struct {
	Type        enum.VerificationType `json:"type"`
	Account     string                `json:"account"`
	UserID      *int64                `json:"user_id,omitempty"`
	Code        string                `json:"code"`
	Attempts    int32                 `json:"attempts"`
	MaxAttempts int32                 `json:"max_attempts"`
	CreatedAt   *time.Time            `json:"created_at,omitempty"`
	ExpiresAt   *time.Time            `json:"expires_at,omitempty"`
}

type RegisterDraft struct {
	Name         string     `json:"name"`
	Nickname     *string    `json:"nickname,omitempty"`
	PasswordHash string     `json:"password_hash"`
	Email        *string    `json:"email,omitempty"`
	Phone        *string    `json:"phone,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
}
