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
		field.Enum("restriction").Values(contentenum.ContentRestrictionMap.EnumValues()...).Default(string(contentenum.ContentRestrictionNone)).Comment("管理限制"),
		field.Int64("created_by").Comment("创建人ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人ID").Nillable().Optional(),
	}
	return fields
}

func (ArticlePostscript) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ArticlePostscript) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联附言 多对一
		edge.From("article", Article.Type).Ref("postscripts").Required().Unique().Field("article_id"),
	}
}
