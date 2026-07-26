package schema

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RbacAuditLog 记录权限配置变更审计事实。
type RbacAuditLog struct {
	ent.Schema
}

func (RbacAuditLog) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{
		Table: constant.TablePrefixUser.String() + "rbac_audit_logs",
	}, entsql.WithComments(true)}
}

func (RbacAuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("operator_id").Comment("操作人账号 ID"),
		field.Enum("realm").Values(commonenum.LoginRealmMap.EnumValues()...).Comment("权限域"),
		field.String("action").Comment("操作动作").MaxLen(64).NotEmpty(),
		field.Int64("user_id").Comment("目标账号 ID").Optional().Nillable(),
		field.Int64("role_id").Comment("目标角色 ID").Optional().Nillable(),
		field.Int64("permission_id").Comment("目标权限 ID").Optional().Nillable(),
		field.Text("detail").Comment("操作详情").Default(""),
	}
}

func (RbacAuditLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (RbacAuditLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("operator_id", "created_at").StorageKey("user_rbac_audit_logs_operator_created_idx"),
		index.Fields("realm", "created_at").StorageKey("user_rbac_audit_logs_realm_created_idx"),
	}
}

func (RbacAuditLog) Edges() []ent.Edge {
	return nil
}
