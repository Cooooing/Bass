package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RbacRolePermission 记录角色与权限的绑定关系。
type RbacRolePermission struct {
	ent.Schema
}

func (RbacRolePermission) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixUser.String() + "rbac_role_permissions"}, entsql.WithComments(true)}
}

func (RbacRolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("role_id").Comment("角色 ID"),
		field.Int64("permission_id").Comment("权限 ID"),
	}
}

func (RbacRolePermission) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (RbacRolePermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "permission_id").Unique().StorageKey("user_rbac_role_permissions_role_permission_unique"),
		index.Fields("permission_id").StorageKey("user_rbac_role_permissions_permission_idx"),
	}
}

func (RbacRolePermission) Edges() []ent.Edge { return nil }