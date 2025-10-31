package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticleVoteRecord 投票记录表
type ArticleVoteRecord struct {
	ent.Schema
}

func (ArticleVoteRecord) Fields() []ent.Field {
	return append([]ent.Field{
		field.Int("vote_id").Comment("所属投票ID"),
		field.Int("user_id").Comment("投票用户ID"),
		field.Int("option_index").Comment("投票选项索引"),
		field.Bool("anonymous").Comment("是否匿名").Default(false),
	}, pkg.TimeAuditFields()...)
}

func (ArticleVoteRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联投票 多对一
		edge.From("vote", ArticleVote.Type).Ref("records").Required().Unique().Field("vote_id"),
	}
}
