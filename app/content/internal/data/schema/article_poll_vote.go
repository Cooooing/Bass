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

// ArticlePollVote 保存账号不可修改的一次投票选择。
type ArticlePollVote struct {
	ent.Schema
}

func (ArticlePollVote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "article_poll_votes"},
		entsql.WithComments(true),
	}
}

func (ArticlePollVote) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("poll_id").Comment("投票 ID").Immutable(),
		field.Int64("user_id").Comment("投票账号 ID").Immutable(),
		field.JSON("option_ids", []int64{}).Comment("所选选项 ID"),
	}
}

func (ArticlePollVote) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("poll_id", "user_id").Unique(),
	}
}

func (ArticlePollVote) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
