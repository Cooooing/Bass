package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return append([]ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("创建用户id"),
		field.String("name").Comment("标签名称").NotEmpty(),
		field.String("description").Comment("标签描述"),
		field.Int64("domain_id").Comment("所属领域id").Optional().Nillable(),
		field.Int32("status").Comment("标签状态：0-正常，1-禁用").Default(0),
		field.Int32("article_count").Comment("文章数").Default(0),
	}, pkg.TimeAuditFields()...)
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联文章 多对多
		edge.From("article", Article.Type).Ref("tags"),
		// 关联领域 多对一
		edge.From("domain", Domain.Type).Ref("tags").Field("domain_id").Unique(),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{}
}
