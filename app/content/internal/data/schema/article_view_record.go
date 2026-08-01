package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ArticleViewRecord 记录登录用户的文章浏览历史。
type ArticleViewRecord struct {
	ent.Schema
}

func (ArticleViewRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixContent.String() + "article_view_records",
		},
		entsql.WithComments(true),
	}
}

func (ArticleViewRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ArticleViewRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("文章 ID"),
		field.Int64("user_id").Comment("浏览用户 ID"),
		field.String("ip").Comment("浏览 IP").Optional().Nillable(),
		field.String("user_agent").Comment("浏览 User-Agent").Optional().Nillable(),
		field.String("browser_fingerprint").Comment("浏览器指纹").Optional().Nillable(),
		field.Time("viewed_at").Comment("最近浏览时间"),
	}
}

func (ArticleViewRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).Ref("view_records").Field("article_id").Required().Unique(),
	}
}

func (ArticleViewRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id", "user_id").Unique(),
		index.Fields("user_id", "viewed_at"),
	}
}
