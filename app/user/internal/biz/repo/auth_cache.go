package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/enum"
)

type AuthCacheRepo interface {
	SaveCode(ctx context.Context, code *model.VerificationCode, ttl time.Duration) error
	GetCode(ctx context.Context, req *VerificationCodeKeyReq) (*model.VerificationCode, error)
	IncrCodeAttempts(ctx context.Context, req *VerificationCodeKeyReq) (int64, error)
	DeleteCode(ctx context.Context, req *VerificationCodeKeyReq) error

	SaveRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string, draft *model.RegisterDraft, ttl time.Duration) error
	GetRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string) (*model.RegisterDraft, error)
	DeleteRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string) error

	SaveSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration, maxSessions int) error
	GetSession(ctx context.Context, realm commonenum.LoginRealm, sessionID string) (*model.RefreshSession, error)
	TouchSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration) error
	RotateSessionJTI(ctx context.Context, realm commonenum.LoginRealm, sessionID string, oldJTI string, newJTI string, lastSeenAt *time.Time, ttl time.Duration) (bool, error)
	DeleteSession(ctx context.Context, realm commonenum.LoginRealm, userID int64, sessionID string) error
	DeleteUserSessions(ctx context.Context, userID int64) error

	SaveRbacPermissions(ctx context.Context, realm string, userID int64, permissions []string, ttl time.Duration) error
	GetRbacPermissions(ctx context.Context, realm string, userID int64) ([]string, bool, error)
	DeleteRbacPermissions(ctx context.Context, realm string, userID int64) error
	DeleteUserRbacPermissions(ctx context.Context, userID int64) error
	DeleteRealmRbacPermissions(ctx context.Context, realm string) error
}

type VerificationCodeKeyReq struct {
	Type    enum.VerificationType
	Account string
	UserID  *int64
}
