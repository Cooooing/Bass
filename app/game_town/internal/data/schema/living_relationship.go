package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	gameenum "game_town/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LivingRelationship struct {
	ent.Schema
}

func (LivingRelationship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "relationships",
		},
		entsql.WithComments(true),
	}
}

func (LivingRelationship) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (LivingRelationship) Fields() []ent.Field {

	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.Enum("source_type").Values(gameenum.EntityTypeMap.EnumValues()...),
		field.Int64("source_id"),
		field.Enum("target_type").Values(gameenum.EntityTypeMap.EnumValues()...),
		field.Int64("target_id"),
		field.JSON("metrics", map[string]float64{}),
		field.JSON("tags", []string{}).Optional(),
		field.Int64("version").Default(0),
	}
}

func (LivingRelationship) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "source_type", "source_id", "target_type", "target_id").Unique().StorageKey("game_town_relationships_direction_unique"),
		index.Fields("world_id", "target_type", "target_id").StorageKey("game_town_relationships_target_idx"),
	}
}
