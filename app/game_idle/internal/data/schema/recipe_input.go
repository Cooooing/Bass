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

// RecipeInput 定义配方输入物品。
type RecipeInput struct {
	ent.Schema
}

func (RecipeInput) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "recipe_inputs",
		},
		entsql.WithComments(true),
	}
}

func (RecipeInput) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("配方输入编码").MaxLen(96).NotEmpty().Immutable().Unique(),
		field.String("recipe_id").Comment("配方编码").MaxLen(64),
		field.String("item_id").Comment("输入物品编码").MaxLen(64),
		field.Int64("quantity").Comment("输入数量").Positive(),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (RecipeInput) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (RecipeInput) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("recipe_id", "item_id").
			Unique().
			StorageKey("game_idle_recipe_inputs_recipe_item_active_unique").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("recipe_id", "sort").
			StorageKey("game_idle_recipe_inputs_recipe_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("item_id").
			StorageKey("game_idle_recipe_inputs_item_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (RecipeInput) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).
			Ref("inputs").
			Field("recipe_id").
			Unique().
			Required(),
		edge.From("item", Item.Type).
			Ref("recipe_inputs").
			Field("item_id").
			Unique().
			Required(),
	}
}
