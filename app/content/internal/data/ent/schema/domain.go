package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Domain struct {
	ent.Schema
}

func (Domain) Fields() []ent.Field {
	return append([]ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("域名名称").NotEmpty(),
		field.String("description").Comment("域名描述").NotEmpty(),
		field.Int32("status").Comment("状态 0-正常，1-禁用").Default(0).Optional(),
		field.String("url").Comment("领域地址").Nillable().Optional(),
		field.String("icon").Comment("图标").Nillable().Optional(),
		field.Int32("tag_count").Comment("标签数").Default(0),
		field.Bool("is_nav").Comment("是否导航").Default(false).Optional(),
	}, pkg.TimeAuditFields()...)
}

func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联标签 一对多
		edge.To("tags", Tag.Type),
	}
}

func (Domain) Indexes() []ent.Index {
	return []ent.Index{}
}
