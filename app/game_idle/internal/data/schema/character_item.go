package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CharacterItem 定义角色物品余额。
type CharacterItem struct {
	ent.Schema
}

func (CharacterItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "character_items",
		},
		entsql.WithComments(true),
	}
}

func (CharacterItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("character_id").Comment("角色 ID"),
		field.String("item_id").Comment("物品编码").MaxLen(64),
		field.Int64("quantity").Comment("持有数量").NonNegative().Default(0),
		field.Int64("total_obtained").Comment("累计获取数量").NonNegative().Default(0),
		field.Int64("total_consumed").Comment("累计消耗数量").NonNegative().Default(0),
	}
}

func (CharacterItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (CharacterItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("character_id", "item_id").
			Unique().
			StorageKey("game_idle_character_items_character_item_unique"),
		index.Fields("item_id").
			StorageKey("game_idle_character_items_item_idx"),
	}
}

func (CharacterItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("character", Character.Type).
			Ref("items").
			Field("character_id").
			Unique().
			Required(),
		edge.From("item", Item.Type).
			Ref("character_items").
			Field("item_id").
			Unique().
			Required(),
	}
}
