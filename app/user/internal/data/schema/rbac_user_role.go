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

// RbacUserRole 记录账号被授予的角色。
type RbacUserRole struct {
	ent.Schema
}

func (RbacUserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixUser.String() + "rbac_user_roles"}, entsql.WithComments(true)}
}

func (RbacUserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID"),
		field.Int64("role_id").Comment("角色 ID"),
		field.Int64("granted_by").Comment("授权人账号 ID"),
		field.Time("expires_at").Comment("授权过期时间，空表示永久").Optional().Nillable(),
	}
}

func (RbacUserRole) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (RbacUserRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "role_id").Unique().StorageKey("user_rbac_user_roles_user_role_unique"),
		index.Fields("role_id").StorageKey("user_rbac_user_roles_role_idx"),
		index.Fields("user_id", "expires_at").StorageKey("user_rbac_user_roles_user_expires_idx"),
	}
}

func (RbacUserRole) Edges() []ent.Edge { return nil }