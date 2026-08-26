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

// RecipeOutput 定义配方产出物品。
type RecipeOutput struct {
	ent.Schema
}

func (RecipeOutput) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "recipe_outputs",
		},
		entsql.WithComments(true),
	}
}

func (RecipeOutput) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("配方产出编码").MaxLen(96).NotEmpty().Immutable().Unique(),
		field.String("recipe_id").Comment("配方编码").MaxLen(64),
		field.String("item_id").Comment("产出物品编码").MaxLen(64),
		field.Int64("min_quantity").Comment("最小数量").Positive(),
		field.Int64("max_quantity").Comment("最大数量").Positive(),
		field.Int32("weight").Comment("随机产出权重").Positive().Default(1),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (RecipeOutput) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (RecipeOutput) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("recipe_id", "item_id").
			Unique().
			StorageKey("game_idle_recipe_outputs_recipe_item_active_unique").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("recipe_id", "sort").
			StorageKey("game_idle_recipe_outputs_recipe_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("item_id").
			StorageKey("game_idle_recipe_outputs_item_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (RecipeOutput) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).
			Ref("outputs").
			Field("recipe_id").
			Unique().
			Required(),
		edge.From("item", Item.Type).
			Ref("recipe_outputs").
			Field("item_id").
			Unique().
			Required(),
	}
}
