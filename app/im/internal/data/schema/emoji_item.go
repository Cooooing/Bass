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

// EmojiItem 是分组中引用的平台图片资产。
type EmojiItem struct{ ent.Schema }

func (EmojiItem) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixIM.String() + "emoji_items"}, entsql.WithComments(true)}
}

func (EmojiItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("表情项所属账号 ID"),
		field.Int64("group_id").Comment("可为空的所属表情分组 ID").Optional().Nillable(),
		field.Int64("asset_id").Comment("平台图片资产 ID"),
		field.String("name").Comment("可选表情别名").Optional(),
		field.Int64("sort").Comment("分组内展示排序").Default(0),
	}
}

func (EmojiItem) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (EmojiItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("group_id", "asset_id").Unique(),
		index.Fields("user_id", "group_id", "sort"),
	}
}
