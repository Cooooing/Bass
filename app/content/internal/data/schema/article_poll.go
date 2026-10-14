package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticlePoll 保存发布后不可修改的投票配置。
type ArticlePoll struct {
	ent.Schema
}

func (ArticlePoll) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_polls"},
		entsql.WithComments(true),
	}
}

func (ArticlePoll) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("文章 ID").Immutable().Unique(),
		field.Time("ends_at").Comment("截止时间"),
		field.Int32("max_choices").Comment("每人最多选择数"),
		field.Int64("participant_count").Comment("参与人数").Default(0),
		field.Int64("total_vote_count").Comment("总票数").Default(0),
	}
}

func (ArticlePoll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("options", ArticlePollOption.Type),
		edge.To("votes", ArticlePollVote.Type),
	}
}

func (ArticlePoll) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
