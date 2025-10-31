package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticleLotteryWinner 抽奖获奖记录
type ArticleLotteryWinner struct {
	ent.Schema
}

func (ArticleLotteryWinner) Fields() []ent.Field {
	return append([]ent.Field{
		field.Int("lottery_id").Comment("所属抽奖ID"),
		field.Int("user_id").Comment("获奖用户ID"),
		field.String("prize").Comment("奖品名称").NotEmpty(),
	}, pkg.TimeAuditFields()...)
}

func (ArticleLotteryWinner) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联抽奖 多对一
		edge.From("lottery", ArticleLottery.Type).Ref("winners").Required().Unique().Field("lottery_id"),
	}
}
