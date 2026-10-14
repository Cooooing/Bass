package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ArticlePollOption 保存投票选项及其聚合票数。
type ArticlePollOption struct {
	ent.Schema
}

func (ArticlePollOption) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_poll_options"},
		entsql.WithComments(true),
	}
}

func (ArticlePollOption) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("poll_id").Comment("投票 ID").Immutable(),
		field.String("content").Comment("选项内容").NotEmpty(),
		field.Int32("sort").Comment("排序"),
		field.Int32("vote_count").Comment("票数").Default(0),
	}
}

func (ArticlePollOption) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
