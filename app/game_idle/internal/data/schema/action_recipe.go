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

// ActionRecipe 定义行动与配方的绑定关系。
type ActionRecipe struct {
	ent.Schema
}

func (ActionRecipe) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "action_recipes",
		},
		entsql.WithComments(true),
	}
}

func (ActionRecipe) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("行动配方绑定编码").MaxLen(128).NotEmpty().Immutable().Unique(),
		field.String("action_id").Comment("行动编码").MaxLen(64),
		field.String("recipe_id").Comment("配方编码").MaxLen(64),
		field.Bool("enabled").Comment("是否启用").Default(true),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (ActionRecipe) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (ActionRecipe) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("action_id", "recipe_id").
			Unique().
			StorageKey("game_idle_action_recipes_action_recipe_active_unique").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("action_id", "enabled", "sort").
			StorageKey("game_idle_action_recipes_action_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("recipe_id").
			StorageKey("game_idle_action_recipes_recipe_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (ActionRecipe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action", Action.Type).
			Ref("recipes").
			Field("action_id").
			Unique().
			Required(),
		edge.From("recipe", Recipe.Type).
			Ref("action_recipes").
			Field("recipe_id").
			Unique().
			Required(),
	}
}
