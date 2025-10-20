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
		field.String("name").Comment("域名名称").NotEmpty(),
		field.String("description").Comment("域名描述").NotEmpty(),
		field.Int("status").Comment("状态 0-正常，1-禁用"),
		field.String("url").Comment("领域地址"),
		field.String("icon").Comment("图标"),
		field.Int("tag_count").Comment("标签数").Default(0),
		field.Bool("is_nav").Comment("是否导航").Default(false),
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
