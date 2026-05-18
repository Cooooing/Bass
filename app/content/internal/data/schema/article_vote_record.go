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

// ArticleVoteRecord 投票记录表
type ArticleVoteRecord struct {
	ent.Schema
}

func (ArticleVoteRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_vote_records"},
		entsql.WithComments(true),
	}
}

func (ArticleVoteRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("vote_id").Comment("所属投票ID"),
		field.Int64("user_id").Comment("投票用户ID"),
		field.Int32("option_index").Comment("投票选项索引"),
		field.Bool("anonymous").Comment("是否匿名").Default(false),
	}
}

func (ArticleVoteRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ArticleVoteRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联投票 多对一
		edge.From("vote", ArticleVote.Type).Ref("records").Required().Unique().Field("vote_id"),
	}
}
