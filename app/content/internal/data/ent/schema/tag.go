package schema

import (
	"common/pkg"
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "tags"},
		entsql.WithComments(true),
	}
}

func (Tag) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("标签名称").NotEmpty(),
		field.String("description").Comment("标签描述").Optional().Nillable(),
		field.Int64("domain_id").Comment("所属领域id").Optional().Nillable(),
		field.Int32("status").Comment("标签状态：1-正常，2-禁用").Default(1),
		field.Int32("article_count").Comment("文章数").Default(0),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
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
