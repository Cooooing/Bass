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

// ArticleLotteryParticipant 抽奖参与记录
type ArticleLotteryParticipant struct {
	ent.Schema
}

func (ArticleLotteryParticipant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_lottery_participants"},
		entsql.WithComments(true),
	}
}

func (ArticleLotteryParticipant) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("lottery_id").Comment("所属抽奖ID"),
		field.Int64("user_id").Comment("参与用户ID"),
	}
}

func (ArticleLotteryParticipant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ArticleLotteryParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联抽奖 多对一
		edge.From("lottery", ArticleLottery.Type).Ref("participants").Required().Unique().Field("lottery_id"),
	}
}
