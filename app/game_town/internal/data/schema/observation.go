package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"game_town/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Observation struct {
	ent.Schema
}

func (Observation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "observations"},
		entsql.WithComments(true),
	}
}

func (Observation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (Observation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.Int64("event_id"),
		field.Int64("npc_id").Optional().Nillable(),
		field.Int64("player_id").Optional().Nillable(),
		field.Enum("source").Values(enum.ObservationSourceMap.EnumValues()...),
		field.Enum("certainty").Values(enum.KnowledgeCertaintyMap.EnumValues()...),
		field.Text("summary"),
		field.Float("salience"),
		field.Time("observed_at"),
		field.Time("world_time"),
	}
}

func (Observation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "player_id", "event_id"),
		index.Fields("world_id", "npc_id", "event_id"),
	}
}
