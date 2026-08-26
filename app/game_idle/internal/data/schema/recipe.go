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

// Recipe 定义独立配方，行动通过绑定关系引用配方。
type Recipe struct {
	ent.Schema
}

func (Recipe) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "recipes",
		},
		entsql.WithComments(true),
	}
}

func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("配方编码").MaxLen(64).NotEmpty().Immutable().Unique(),
		field.String("name").Comment("配方名称").MaxRuneLen(128).NotEmpty(),
		field.String("description").Comment("配方描述").MaxRuneLen(1024).Default(""),
		field.Enum("type").
			Values(gameenum.RecipeTypeValues()...).
			Default(gameenum.RecipeTypeNormal.String()).
			Comment("配方类型"),
		field.Int32("generation_times").Comment("产出生成次数").Positive().Default(1),
		field.Bool("enabled").Comment("是否启用").Default(true),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (Recipe) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Recipe) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "enabled", "sort").
			StorageKey("game_idle_recipes_type_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("enabled", "sort").
			StorageKey("game_idle_recipes_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (Recipe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("action_recipes", ActionRecipe.Type),
		edge.To("inputs", RecipeInput.Type),
		edge.To("outputs", RecipeOutput.Type),
	}
}
