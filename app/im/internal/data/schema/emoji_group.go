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

// EmojiGroup 是账号拥有的表情分组事实。
type EmojiGroup struct{ ent.Schema }

func (EmojiGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixIM.String() + "emoji_groups"}, entsql.WithComments(true)}
}

func (EmojiGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("分组所属账号 ID"),
		field.String("name").Comment("用户定义的分组名称").NotEmpty(),
		field.Int64("sort").Comment("同一账号内的展示排序").Default(0),
	}
}

func (EmojiGroup) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (EmojiGroup) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "name").Unique(),
		index.Fields("user_id", "sort"),
	}
}
