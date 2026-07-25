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

// RbacPermission 定义指定 realm 下可校验的权限点。
type RbacPermission struct {
	ent.Schema
}

func (RbacPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixUser.String() + "rbac_permissions"}, entsql.WithComments(true)}
}

func (RbacPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("realm").Values(commonenum.LoginRealmMap.EnumValues()...).Comment("权限域"),
		field.String("code").Comment("权限编码").MaxLen(128).NotEmpty(),
		field.String("name").Comment("权限名称").MaxLen(128).NotEmpty(),
		field.String("description").Comment("权限说明").MaxLen(512).Default(""),
		field.Bool("enabled").Comment("是否启用").Default(true),
	}
}

func (RbacPermission) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (RbacPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("realm", "code").Unique().StorageKey("user_rbac_permissions_realm_code_unique"),
		index.Fields("realm", "enabled").StorageKey("user_rbac_permissions_realm_enabled_idx"),
	}
}

func (RbacPermission) Edges() []ent.Edge { return nil }
