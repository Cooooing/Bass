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

type WorldRule struct {
	ent.Schema
}

func (WorldRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "world_rules",
		},
		entsql.WithComments(true),
	}
}

func (WorldRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (WorldRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("version").Comment("规则版本").Default(1),
		field.JSON("rules", map[string]any{}).Comment("结构化世界规则"),
	}
}

func (WorldRule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id").Unique().StorageKey("game_town_world_rules_world_unique"),
	}
}
