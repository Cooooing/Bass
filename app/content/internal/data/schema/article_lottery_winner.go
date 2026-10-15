package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ArticleLotteryWinner 保存不可变的抽奖中奖结果。
type ArticleLotteryWinner struct {
	ent.Schema
}

func (ArticleLotteryWinner) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_lottery_winners"}, entsql.WithComments(true)}
}

func (ArticleLotteryWinner) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("lottery_id").Comment("抽奖 ID").Immutable(),
		field.Int64("prize_id").Comment("奖项 ID").Immutable(),
		field.Int64("user_id").Comment("中奖账号 ID").Immutable(),
		field.Int32("sort").Comment("奖项内中奖顺序"),
	}
}

func (ArticleLotteryWinner) Indexes() []ent.Index {
	return []ent.Index{index.Fields("lottery_id", "user_id").Unique(), index.Fields("prize_id", "sort").Unique()}
}

func (ArticleLotteryWinner) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
