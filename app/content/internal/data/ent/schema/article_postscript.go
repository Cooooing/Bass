package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticlePostscript 附言表
type ArticlePostscript struct {
	ent.Schema
}

func (ArticlePostscript) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("所属文章ID"),
		field.Text("content").Comment("附言内容").NotEmpty(),
		field.Int32("status").Comment("状态 0-正常 1-隐藏").Default(0),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (ArticlePostscript) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联附言 多对一
		edge.From("article", Article.Type).Ref("postscripts").Required().Unique().Field("article_id"),
	}
}
