package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	contentenum "content/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticlePostscript 附言表
type ArticlePostscript struct {
	ent.Schema
}

func (ArticlePostscript) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_postscripts"},
		entsql.WithComments(true),
	}
}

func (ArticlePostscript) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("所属文章ID"),
		field.Text("content").Comment("附言内容").NotEmpty(),
		field.Enum("status").Values(contentenum.ArticlePostscriptStatusMap.EnumValues()...).Default(string(contentenum.ArticlePostscriptStatusNormal)).Comment("状态"),
	}
	return fields
}

func (ArticlePostscript) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.UserAuditMixin{},
	}
}

func (ArticlePostscript) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联附言 多对一
		edge.From("article", Article.Type).Ref("postscripts").Required().Unique().Field("article_id"),
	}
}
