package model

type VerificationCode struct {
	Type          string `json:"type"`
	Account       string `json:"account"`
	Code          string `json:"code"`
	Attempts      int32  `json:"attempts"`
	MaxAttempts   int32  `json:"max_attempts"`
	CreatedAtUnix int64  `json:"created_at_unix"`
	ExpiresAtUnix int64  `json:"expires_at_unix"`
}

type RegisterDraft struct {
	Name          string  `json:"name"`
	Nickname      *string `json:"nickname,omitempty"`
	PasswordHash  string  `json:"password_hash"`
	Email         *string `json:"email,omitempty"`
	Phone         *string `json:"phone,omitempty"`
	CreatedAtUnix int64   `json:"created_at_unix"`
	ExpiresAtUnix int64   `json:"expires_at_unix"`
}
