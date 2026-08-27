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

// Region 定义挂机游戏区域，区域只作为行动的前端展示分类。
type Region struct {
	ent.Schema
}

func (Region) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "regions",
		},
		entsql.WithComments(true),
	}
}

func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("区域编码").MaxLen(64).NotEmpty().Immutable().Unique(),
		field.String("name").Comment("区域名称").MaxRuneLen(128).NotEmpty(),
		field.String("description").Comment("区域描述").MaxRuneLen(1024).Default(""),
		field.Enum("action_kind").
			Values(gameenum.ActionKindValues()...).
			Default(gameenum.ActionKindForaging.String()).
			Comment("归属行动类型"),
		field.Bool("enabled").Comment("是否启用").Default(true),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (Region) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Region) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("action_kind", "enabled", "sort").
			StorageKey("game_idle_regions_kind_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("enabled", "sort").
			StorageKey("game_idle_regions_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actions", Action.Type),
	}
}
