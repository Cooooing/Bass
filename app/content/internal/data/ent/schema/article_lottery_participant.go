package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticleLotteryParticipant 抽奖参与记录
type ArticleLotteryParticipant struct {
	ent.Schema
}

func (ArticleLotteryParticipant) Fields() []ent.Field {
	return append([]ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("lottery_id").Comment("所属抽奖ID"),
		field.Int64("user_id").Comment("参与用户ID"),
	}, pkg.TimeAuditFields()...)
}

func (ArticleLotteryParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联抽奖 多对一
		edge.From("lottery", ArticleLottery.Type).Ref("participants").Required().Unique().Field("lottery_id"),
	}
}
