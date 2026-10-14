package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ArticleQA 保存问答文章一次性的采纳事实。
type ArticleQA struct {
	ent.Schema
}

func (ArticleQA) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_qas"},
		entsql.WithComments(true),
	}
}

func (ArticleQA) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("文章 ID").Immutable().Unique(),
		field.Int64("accepted_comment_id").Comment("已采纳评论 ID").Immutable().Unique(),
		field.Time("accepted_at").Comment("采纳时间"),
	}
}

func (ArticleQA) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }
