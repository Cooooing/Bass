package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Group 用户组实体定义
type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field {
	return append([]ent.Field{
		field.String("name").Comment("用户组名").NotEmpty(),
		field.String("endpoint").Comment("端点").NotEmpty(),
		field.String("description").Comment("描述").Optional(),
		field.String("module").Comment("模块").Optional(),
	}, pkg.TimeAuditFields()...)
}

func (Group) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "endpoint").Unique(),
	}
}
