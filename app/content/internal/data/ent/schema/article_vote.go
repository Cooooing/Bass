package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticleVote 投票表
type ArticleVote struct {
	ent.Schema
}

func (ArticleVote) Fields() []ent.Field {
	return append([]ent.Field{
		field.Int("article_id").Comment("所属文章ID"),
		field.JSON("vote_options", []string{}).Comment("投票选项").Optional(),
		field.JSON("vote_counts", []int{}).Comment("各选项票数").Optional(),
		field.Bool("vote_multiple").Comment("是否允许多选").Default(false),
		field.Bool("vote_anonymous").Comment("是否匿名投票").Default(true),
		field.Int("total_count").Comment("总投票数").Default(0),
		field.Time("end_at").Comment("投票截止时间").Optional(),
	}, pkg.TimeAuditFields()...)
}

func (ArticleVote) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联文章 多对一
		edge.From("article", Article.Type).Ref("votes").Required().Unique().Field("article_id"),
		// 关联投票记录 一对多
		edge.To("records", ArticleVoteRecord.Type),
	}
}
