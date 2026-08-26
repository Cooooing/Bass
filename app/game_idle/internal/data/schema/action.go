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

// Action 定义队列可执行行动；区域只作为展示分类。
type Action struct {
	ent.Schema
}

func (Action) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "actions",
		},
		entsql.WithComments(true),
	}
}

func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("行动编码").MaxLen(64).NotEmpty().Immutable().Unique(),
		field.String("name").Comment("行动名称").MaxRuneLen(128).NotEmpty(),
		field.String("description").Comment("行动描述").MaxRuneLen(1024).Default(""),
		field.String("region_id").Comment("展示区域编码").MaxLen(64).Nillable().Optional(),
		field.Enum("action_kind").
			Values(gameenum.ActionKindValues()...).
			Default(gameenum.ActionKindForaging.String()).
			Comment("行动类型"),
		field.Enum("ability_id").
			Values(gameenum.AbilityValues()...).
			Comment("行动等级对应能力编码").
			Nillable().
			Optional(),
		field.Int32("required_ability_level").Comment("要求能力等级").NonNegative().Default(0),
		field.Int32("duration_seconds").Comment("基础耗时秒数").Positive(),
		field.Int64("exp_reward").Comment("完成后获得经验").NonNegative().Default(0),
		field.Bool("enabled").Comment("是否启用").Default(true),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (Action) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Action) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("region_id", "enabled", "sort").
			StorageKey("game_idle_actions_region_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("action_kind", "enabled", "sort").
			StorageKey("game_idle_actions_kind_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("ability_id"),
		index.Fields("deleted_at"),
	}
}

func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("region", Region.Type).
			Ref("actions").
			Field("region_id").
			Unique(),
		edge.To("recipes", ActionRecipe.Type),
	}
}
