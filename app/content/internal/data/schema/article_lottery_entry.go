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

// ArticleLotteryEntry 保存账号对抽奖的有效报名。
type ArticleLotteryEntry struct {
	ent.Schema
}

func (ArticleLotteryEntry) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_lottery_entries"}, entsql.WithComments(true)}
}

func (ArticleLotteryEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("lottery_id").Comment("抽奖 ID").Immutable(),
		field.Int64("user_id").Comment("报名账号 ID").Immutable(),
	}
}

func (ArticleLotteryEntry) Indexes() []ent.Index {
	return []ent.Index{index.Fields("lottery_id", "user_id").Unique()}
}

func (ArticleLotteryEntry) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
