package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticleLottery 抽奖表
type ArticleLottery struct {
	ent.Schema
}

func (ArticleLottery) Fields() []ent.Field {
	return append([]ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("所属文章ID"),
		field.JSON("prizes", []string{}).Comment("奖品列表").Optional(),
		field.Time("start_at").Comment("抽奖开始时间").Optional(),
		field.Time("end_at").Comment("抽奖结束时间").Optional(),
		field.Int32("status").Comment("状态 0-未开始 1-进行中 2-已结束").Default(0),
	}, pkg.TimeAuditFields()...)
}

func (ArticleLottery) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联文章 多对一
		edge.From("article", Article.Type).Ref("lotteries").Required().Unique().Field("article_id"),
		// 关联参与用户 一对多
		edge.To("participants", ArticleLotteryParticipant.Type),
		// 关联中奖用户 一对多
		edge.To("winners", ArticleLotteryWinner.Type),
	}
}
