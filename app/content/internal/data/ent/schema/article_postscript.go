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
	return append([]ent.Field{
		field.Int("article_id").Comment("所属文章ID"),
		field.Text("content").Comment("附言内容").NotEmpty(),
	}, pkg.TimeAuditFields()...)
}

func (ArticlePostscript) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联附言 多对一
		edge.From("article", Article.Type).Ref("postscripts").Required().Unique().Field("article_id"),
	}
}
