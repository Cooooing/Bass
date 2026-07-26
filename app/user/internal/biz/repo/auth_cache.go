package repo

import (
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/enum"
)

type AuthCacheRepo interface {
	SaveCode(ctx context.Context, code *model.VerificationCode, ttl time.Duration) error
	GetCode(ctx context.Context, codeType enum.VerificationType, account string) (*model.VerificationCode, error)
	IncrCodeAttempts(ctx context.Context, codeType enum.VerificationType, account string) (int64, error)
	DeleteCode(ctx context.Context, codeType enum.VerificationType, account string) error

	SaveRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string, draft *model.RegisterDraft, ttl time.Duration) error
	GetRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string) (*model.RegisterDraft, error)
	DeleteRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string) error

	SaveSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration) error
	GetSession(ctx context.Context, sessionID string) (*model.RefreshSession, error)
	TouchSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration) error
	RotateSessionJTI(ctx context.Context, sessionID string, oldJTI string, newJTI string, lastSeenAt *time.Time, ttl time.Duration) (bool, error)
	DeleteSession(ctx context.Context, userID int64, sessionID string) error
	DeleteUserSessions(ctx context.Context, userID int64) error

	SaveRbacPermissions(ctx context.Context, realm string, userID int64, permissions []string, ttl time.Duration) error
	GetRbacPermissions(ctx context.Context, realm string, userID int64) ([]string, bool, error)
	DeleteRbacPermissions(ctx context.Context, realm string, userID int64) error
	DeleteUserRbacPermissions(ctx context.Context, userID int64) error
	DeleteRealmRbacPermissions(ctx context.Context, realm string) error
}
