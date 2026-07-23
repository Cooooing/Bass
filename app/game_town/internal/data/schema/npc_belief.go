package schema

import (
	"common/pkg/constant"
	gameenum "game_town/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type NpcBelief struct {
	ent.Schema
}

func (NpcBelief) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "npc_beliefs"},
		entsql.WithComments(true),
	}
}

func (NpcBelief) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.Int64("npc_id"),
		field.Int64("claim_id"),
		field.Int64("source_npc_id").Nillable().Optional(),
		field.Int64("source_player_id").Nillable().Optional(),
		field.Int64("source_event_id").Nillable().Optional(),
		field.Enum("stance").Values(gameenum.BeliefStanceMap.EnumValues()...).Default(string(gameenum.BeliefStanceBelieves)),
		field.Float("confidence").Default(0.5),
		field.Time("learned_at"),
		field.Time("updated_at"),
	}
}

func (NpcBelief) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "npc_id", "claim_id").Unique().StorageKey("game_town_npc_beliefs_npc_claim_unique"),
		index.Fields("world_id", "npc_id", "confidence").StorageKey("game_town_npc_beliefs_npc_confidence_idx"),
	}
}
