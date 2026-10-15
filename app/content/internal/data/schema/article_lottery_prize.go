package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ArticleLotteryPrize 保存抽奖奖项配置。
type ArticleLotteryPrize struct {
	ent.Schema
}

func (ArticleLotteryPrize) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_lottery_prizes"}, entsql.WithComments(true)}
}

func (ArticleLotteryPrize) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("lottery_id").Comment("抽奖 ID").Immutable(),
		field.String("name").Comment("奖项名称").NotEmpty(),
		field.String("description").Comment("奖项说明").Optional().Nillable(),
		field.Int32("winner_limit").Comment("中奖名额"),
		field.Int32("sort").Comment("奖项排序"),
	}
}

func (ArticleLotteryPrize) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
