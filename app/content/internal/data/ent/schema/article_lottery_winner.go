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

// ArticleLotteryWinner 抽奖获奖记录
type ArticleLotteryWinner struct {
	ent.Schema
}

func (ArticleLotteryWinner) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_lottery_winners"},
		entsql.WithComments(true),
	}
}

func (ArticleLotteryWinner) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("lottery_id").Comment("所属抽奖ID"),
		field.Int64("user_id").Comment("获奖用户ID"),
		field.String("prize").Comment("奖品名称").NotEmpty(),
	}
}

func (ArticleLotteryWinner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ArticleLotteryWinner) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联抽奖 多对一
		edge.From("lottery", ArticleLottery.Type).Ref("winners").Required().Unique().Field("lottery_id"),
	}
}
