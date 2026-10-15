package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ArticleLottery 保存文章抽奖的时间和聚合报名人数。
type ArticleLottery struct {
	ent.Schema
}

func (ArticleLottery) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_lotteries"}, entsql.WithComments(true)}
}

func (ArticleLottery) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("文章 ID").Immutable().Unique(),
		field.Time("draw_at").Comment("计划开奖时间"),
		field.Time("drawn_at").Comment("实际开奖时间").Optional().Nillable(),
		field.Int64("entry_count").Comment("当前有效报名人数").Default(0),
	}
}

func (ArticleLottery) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
