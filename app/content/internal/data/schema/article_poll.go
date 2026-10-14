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

// ArticlePollVote 保存账号不可修改的一次投票选择。
type ArticlePollVote struct {
	ent.Schema
}

func (ArticlePollVote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_poll_votes"},
		entsql.WithComments(true),
	}
}

func (ArticlePollVote) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("poll_id").Comment("投票 ID").Immutable(),
		field.Int64("user_id").Comment("投票账号 ID").Immutable(),
		field.JSON("option_ids", []int64{}).Comment("所选选项 ID"),
	}
}

func (ArticlePollVote) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("poll_id", "user_id").Unique(),
	}
}

func (ArticlePollVote) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
