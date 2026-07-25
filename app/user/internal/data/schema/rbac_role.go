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

// RbacRole 定义指定 realm 下的角色。
type RbacRole struct {
	ent.Schema
}

func (RbacRole) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{
		Table: constant.TablePrefixUser.String() + "rbac_roles",
	}, entsql.WithComments(true)}
}

func (RbacRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("realm").Values(commonenum.LoginRealmMap.EnumValues()...).Comment("权限域"),
		field.String("code").Comment("角色编码").MaxLen(128).NotEmpty(),
		field.String("name").Comment("角色名称").MaxLen(128).NotEmpty(),
		field.String("description").Comment("角色说明").MaxLen(512).Default(""),
		field.Bool("built_in").Comment("是否内置角色").Default(false),
		field.Bool("enabled").Comment("是否启用").Default(true),
	}
}

func (RbacRole) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (RbacRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("realm", "code").Unique().StorageKey("user_rbac_roles_realm_code_unique"),
		index.Fields("realm", "enabled").StorageKey("user_rbac_roles_realm_enabled_idx"),
	}
}

func (RbacRole) Edges() []ent.Edge {
	return nil
}
