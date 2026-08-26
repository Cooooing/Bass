package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	gameenum "game_idle/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Item 定义挂机游戏物品，货币、战利品、消耗品、技能书、装备、资源都统一作为物品。
type Item struct {
	ent.Schema
}

func (Item) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "items",
		},
		entsql.WithComments(true),
	}
}

func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("物品编码").MaxLen(64).NotEmpty().Immutable().Unique(),
		field.String("name").Comment("物品名称").MaxRuneLen(128).NotEmpty(),
		field.Enum("type").
			Values(gameenum.ItemTypeValues()...).
			Default(gameenum.ItemTypeResource.String()).
			Comment("物品类型"),
		field.String("description").Comment("物品描述").MaxRuneLen(1024).Default(""),
		field.Bool("enabled").Comment("是否启用").Default(true),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Item) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "enabled", "sort").
			StorageKey("game_idle_items_type_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipe_inputs", RecipeInput.Type),
		edge.To("recipe_outputs", RecipeOutput.Type),
		edge.To("character_items", CharacterItem.Type),
	}
}
