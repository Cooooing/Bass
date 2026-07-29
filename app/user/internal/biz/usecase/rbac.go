package usecase

import (
	commonenum "common/pkg/enum"
	"context"
	"strings"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"

	"github.com/samber/lo"
)

type RbacUsecase struct {
	accountRepo   repo.AccountRepo
	rbacRepo      repo.RbacRepo
	authCacheRepo repo.AuthCacheRepo
}

func NewRbacUsecase(
	accountRepo repo.AccountRepo,
	rbacRepo repo.RbacRepo,
	authCacheRepo repo.AuthCacheRepo,
) *RbacUsecase {
	return &RbacUsecase{
		accountRepo:   accountRepo,
		rbacRepo:      rbacRepo,
		authCacheRepo: authCacheRepo,
	}
}

func (u *RbacUsecase) UpsertRole(ctx context.Context, row *model.RbacRole) (*model.RbacRole, error) {
	saved, err := u.rbacRepo.UpsertRole(ctx, row)
	if err != nil {
		return nil, err
	}
	_ = u.authCacheRepo.DeleteRealmRbacPermissions(ctx, saved.Realm.String())
	return saved, nil
}

func (u *RbacUsecase) UpsertPermission(ctx context.Context, row *model.RbacPermission) (*model.RbacPermission, error) {
	saved, err := u.rbacRepo.UpsertPermission(ctx, row)
	if err != nil {
		return nil, err
	}
	_ = u.authCacheRepo.DeleteRealmRbacPermissions(ctx, saved.Realm.String())
	return saved, nil
}

func (u *RbacUsecase) BindRolePermission(ctx context.Context, roleID int64, permissionID int64) error {
	role, err := u.rbacRepo.GetRole(ctx, roleID)
	if err != nil {
		return err
	}
	if err := u.rbacRepo.BindRolePermission(ctx, roleID, permissionID); err != nil {
		return err
	}
	u.deleteRoleRealmCache(ctx, role)
	return nil
}

func (u *RbacUsecase) UnbindRolePermission(ctx context.Context, roleID int64, permissionID int64) error {
	role, err := u.rbacRepo.GetRole(ctx, roleID)
	if err != nil {
		return err
	}
	if err := u.rbacRepo.UnbindRolePermission(ctx, roleID, permissionID); err != nil {
		return err
	}
	u.deleteRoleRealmCache(ctx, role)
	return nil
}

func (u *RbacUsecase) GrantRole(ctx context.Context, userID int64, roleID int64, grantedBy int64, expiresAt *time.Time) error {
	role, err := u.rbacRepo.GetRole(ctx, roleID)
	if err != nil {
		return err
	}
	if err := u.rbacRepo.GrantRole(ctx, userID, roleID, grantedBy, expiresAt); err != nil {
		return err
	}
	u.deleteUserRoleCache(ctx, userID, role)
	return nil
}

func (u *RbacUsecase) RevokeRole(ctx context.Context, userID int64, roleID int64) error {
	role, err := u.rbacRepo.GetRole(ctx, roleID)
	if err != nil {
		return err
	}
	if err := u.rbacRepo.RevokeRole(ctx, userID, roleID); err != nil {
		return err
	}
	u.deleteUserRoleCache(ctx, userID, role)
	return nil
}

func (u *RbacUsecase) CheckPermission(ctx context.Context, userID int64, realm commonenum.LoginRealm, permissionCode string) (bool, error) {
	code := strings.TrimSpace(permissionCode)
	account, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: &userID,
	})
	if err != nil {
		return false, err
	}
	if account == nil || account.Status == nil || *account.Status != enum.AccountStatusNormal {
		return false, nil
	}
	if permissions, ok, err := u.authCacheRepo.GetRbacPermissions(ctx, realm.String(), userID); err == nil && ok {
		return lo.Contains(permissions, code), nil
	}
	permissions, err := u.rbacRepo.PermissionCodes(ctx, userID, realm, new(time.Now()))
	if err != nil {
		return false, err
	}
	_ = u.authCacheRepo.SaveRbacPermissions(ctx, realm.String(), userID, permissions, 5*time.Minute)
	return lo.Contains(permissions, code), nil
}

func (u *RbacUsecase) deleteRoleRealmCache(ctx context.Context, role *model.RbacRole) {
	if role == nil {
		return
	}
	_ = u.authCacheRepo.DeleteRealmRbacPermissions(ctx, role.Realm.String())
}

func (u *RbacUsecase) deleteUserRoleCache(ctx context.Context, userID int64, role *model.RbacRole) {
	if role == nil {
		_ = u.authCacheRepo.DeleteUserRbacPermissions(ctx, userID)
		return
	}
	_ = u.authCacheRepo.DeleteRbacPermissions(ctx, role.Realm.String(), userID)
}
