package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"time"
	"user/internal/biz/model"
	bizrepo "user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/rbacpermission"
	"user/internal/data/gen/rbacrole"
	"user/internal/data/gen/rbacrolepermission"
	"user/internal/data/gen/rbacuserrole"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.RbacRepo = (*RbacRepo)(nil)

type RbacRepo struct{ db *gen.Client }

func NewRbacRepo(
	db *gen.Client,
) bizrepo.RbacRepo {
	return &RbacRepo{
		db: db,
	}
}

func (r *RbacRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *RbacRepo) GetRole(ctx context.Context, roleID int64) (*model.RbacRole, error) {
	row, err := r.getClient(ctx).RbacRole.Query().Where(rbacrole.ID(roleID)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rbacRoleToModel(row), nil
}

func (r *RbacRepo) UpsertRole(ctx context.Context, row *model.RbacRole) (*model.RbacRole, error) {
	client := r.getClient(ctx)
	if row.ID != 0 {
		saved, err := client.RbacRole.UpdateOneID(row.ID).SetRealm(rbacrole.Realm(row.Realm)).SetCode(row.Code).SetName(row.Name).SetDescription(row.Description).SetBuiltIn(row.BuiltIn).SetEnabled(row.Enabled).Save(ctx)
		if err != nil {
			return nil, err
		}
		return rbacRoleToModel(saved), nil
	}
	existing, err := client.RbacRole.Query().Where(rbacrole.RealmEQ(rbacrole.Realm(row.Realm)), rbacrole.Code(row.Code)).First(ctx)
	if err != nil && !gen.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		saved, err := client.RbacRole.UpdateOneID(existing.ID).SetName(row.Name).SetDescription(row.Description).SetBuiltIn(row.BuiltIn).SetEnabled(row.Enabled).Save(ctx)
		if err != nil {
			return nil, err
		}
		return rbacRoleToModel(saved), nil
	}
	saved, err := client.RbacRole.Create().SetRealm(rbacrole.Realm(row.Realm)).SetCode(row.Code).SetName(row.Name).SetDescription(row.Description).SetBuiltIn(row.BuiltIn).SetEnabled(row.Enabled).Save(ctx)
	if err != nil {
		return nil, err
	}
	return rbacRoleToModel(saved), nil
}

func (r *RbacRepo) UpsertPermission(ctx context.Context, row *model.RbacPermission) (*model.RbacPermission, error) {
	client := r.getClient(ctx)
	if row.ID != 0 {
		saved, err := client.RbacPermission.UpdateOneID(row.ID).SetRealm(rbacpermission.Realm(row.Realm)).SetCode(row.Code).SetName(row.Name).SetDescription(row.Description).SetEnabled(row.Enabled).Save(ctx)
		if err != nil {
			return nil, err
		}
		return rbacPermissionToModel(saved), nil
	}
	existing, err := client.RbacPermission.Query().Where(rbacpermission.RealmEQ(rbacpermission.Realm(row.Realm)), rbacpermission.Code(row.Code)).First(ctx)
	if err != nil && !gen.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		saved, err := client.RbacPermission.UpdateOneID(existing.ID).SetName(row.Name).SetDescription(row.Description).SetEnabled(row.Enabled).Save(ctx)
		if err != nil {
			return nil, err
		}
		return rbacPermissionToModel(saved), nil
	}
	saved, err := client.RbacPermission.Create().SetRealm(rbacpermission.Realm(row.Realm)).SetCode(row.Code).SetName(row.Name).SetDescription(row.Description).SetEnabled(row.Enabled).Save(ctx)
	if err != nil {
		return nil, err
	}
	return rbacPermissionToModel(saved), nil
}

func (r *RbacRepo) BindRolePermission(ctx context.Context, roleID int64, permissionID int64) error {
	exists, err := r.getClient(ctx).RbacRolePermission.Query().Where(rbacrolepermission.RoleID(roleID), rbacrolepermission.PermissionID(permissionID)).Exist(ctx)
	if err != nil || exists {
		return err
	}
	return r.getClient(ctx).RbacRolePermission.Create().SetRoleID(roleID).SetPermissionID(permissionID).Exec(ctx)
}

func (r *RbacRepo) UnbindRolePermission(ctx context.Context, roleID int64, permissionID int64) error {
	_, err := r.getClient(ctx).RbacRolePermission.Delete().Where(rbacrolepermission.RoleID(roleID), rbacrolepermission.PermissionID(permissionID)).Exec(ctx)
	return err
}

func (r *RbacRepo) GrantRole(ctx context.Context, userID int64, roleID int64, grantedBy int64, expiresAt *time.Time) error {
	exists, err := r.getClient(ctx).RbacUserRole.Query().Where(rbacuserrole.UserID(userID), rbacuserrole.RoleID(roleID)).Exist(ctx)
	if err != nil {
		return err
	}
	if exists {
		update := r.getClient(ctx).RbacUserRole.Update().Where(rbacuserrole.UserID(userID), rbacuserrole.RoleID(roleID)).SetGrantedBy(grantedBy)
		if expiresAt == nil {
			update.ClearExpiresAt()
		} else {
			update.SetExpiresAt(*expiresAt)
		}
		_, err = update.Save(ctx)
		return err
	}
	create := r.getClient(ctx).RbacUserRole.Create().SetUserID(userID).SetRoleID(roleID).SetGrantedBy(grantedBy)
	if expiresAt != nil {
		create.SetExpiresAt(*expiresAt)
	}
	return create.Exec(ctx)
}

func (r *RbacRepo) RevokeRole(ctx context.Context, userID int64, roleID int64) error {
	_, err := r.getClient(ctx).RbacUserRole.Delete().Where(rbacuserrole.UserID(userID), rbacuserrole.RoleID(roleID)).Exec(ctx)
	return err
}

func (r *RbacRepo) HasPermission(ctx context.Context, userID int64, realm commonenum.LoginRealm, permissionCode string, now time.Time) (bool, error) {
	codes, err := r.PermissionCodes(ctx, userID, realm, now)
	if err != nil {
		return false, err
	}
	for _, code := range codes {
		if code == permissionCode {
			return true, nil
		}
	}
	return false, nil
}

func (r *RbacRepo) PermissionCodes(ctx context.Context, userID int64, realm commonenum.LoginRealm, now time.Time) ([]string, error) {
	client := r.getClient(ctx)
	userRoles, err := client.RbacUserRole.Query().Where(rbacuserrole.UserID(userID), rbacuserrole.Or(rbacuserrole.ExpiresAtIsNil(), rbacuserrole.ExpiresAtGT(now))).All(ctx)
	if err != nil || len(userRoles) == 0 {
		return nil, err
	}
	roleIDs := make([]int64, 0, len(userRoles))
	for _, row := range userRoles {
		roleIDs = append(roleIDs, row.RoleID)
	}
	roles, err := client.RbacRole.Query().Where(rbacrole.IDIn(roleIDs...), rbacrole.RealmEQ(rbacrole.Realm(realm)), rbacrole.Enabled(true)).All(ctx)
	if err != nil || len(roles) == 0 {
		return nil, err
	}
	roleIDs = roleIDs[:0]
	for _, row := range roles {
		roleIDs = append(roleIDs, row.ID)
	}
	links, err := client.RbacRolePermission.Query().Where(rbacrolepermission.RoleIDIn(roleIDs...)).All(ctx)
	if err != nil || len(links) == 0 {
		return nil, err
	}
	permissionIDs := make([]int64, 0, len(links))
	for _, row := range links {
		permissionIDs = append(permissionIDs, row.PermissionID)
	}
	permissions, err := client.RbacPermission.Query().Where(rbacpermission.IDIn(permissionIDs...), rbacpermission.RealmEQ(rbacpermission.Realm(realm)), rbacpermission.Enabled(true)).All(ctx)
	if err != nil {
		return nil, err
	}
	codes := make([]string, 0, len(permissions))
	seen := make(map[string]struct{}, len(permissions))
	for _, row := range permissions {
		if _, ok := seen[row.Code]; ok {
			continue
		}
		seen[row.Code] = struct{}{}
		codes = append(codes, row.Code)
	}
	return codes, nil
}

func rbacRoleToModel(row *gen.RbacRole) *model.RbacRole {
	if row == nil {
		return nil
	}
	return &model.RbacRole{
		ID:          row.ID,
		Realm:       commonenum.LoginRealm(row.Realm),
		Code:        row.Code,
		Name:        row.Name,
		Description: row.Description,
		BuiltIn:     row.BuiltIn,
		Enabled:     row.Enabled,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func rbacPermissionToModel(row *gen.RbacPermission) *model.RbacPermission {
	if row == nil {
		return nil
	}
	return &model.RbacPermission{
		ID:          row.ID,
		Realm:       commonenum.LoginRealm(row.Realm),
		Code:        row.Code,
		Name:        row.Name,
		Description: row.Description,
		Enabled:     row.Enabled,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}
