package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"time"
	"user/internal/biz/model"
)

type RbacRepo interface {
	GetRole(ctx context.Context, roleID int64) (*model.RbacRole, error)
	UpsertRole(ctx context.Context, row *model.RbacRole) (*model.RbacRole, error)
	UpsertPermission(ctx context.Context, row *model.RbacPermission) (*model.RbacPermission, error)
	BindRolePermission(ctx context.Context, roleID int64, permissionID int64) error
	UnbindRolePermission(ctx context.Context, roleID int64, permissionID int64) error
	GrantRole(ctx context.Context, userID int64, roleID int64, grantedBy int64, expiresAt *time.Time) error
	RevokeRole(ctx context.Context, userID int64, roleID int64) error
	HasAnyRole(ctx context.Context, userID int64, realm commonenum.LoginRealm, now *time.Time) (bool, error)
	HasPermission(ctx context.Context, userID int64, realm commonenum.LoginRealm, permissionCode string, now *time.Time) (bool, error)
	PermissionCodes(ctx context.Context, userID int64, realm commonenum.LoginRealm, now *time.Time) ([]string, error)
}
