package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	gameenum "game_town/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/pgvector/pgvector-go"
)

type NpcMemory struct {
	ent.Schema
}

func (NpcMemory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "npc_memories",
		},
		entsql.WithComments(true),
	}
}

func (NpcMemory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NpcMemory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.Int64("npc_id"),
		field.Int64("source_event_id").Nillable().Optional(),
		field.Int64("source_observation_id").Nillable().Optional(),
		field.Enum("kind").Values(gameenum.MemoryKindMap.EnumValues()...),
		field.Text("content"),
		field.Float("importance").Default(0.5),
		field.Time("occurred_world_time"),
		field.Time("last_recalled_at").Nillable().Optional(),
		field.Other("embedding", pgvector.Vector{}).SchemaType(map[string]string{dialect.Postgres: "vector(1024)"}).Optional().Nillable(),
		field.String("embedding_model").MaxLen(128).Default(""),
		field.Enum("embedding_status").Values(gameenum.EmbeddingStatusMap.EnumValues()...).Default(gameenum.EmbeddingStatusPending.String()),
		field.String("embedding_error").MaxRuneLen(512).Default(""),
	}
}

func (NpcMemory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "npc_id", "occurred_world_time").StorageKey("game_town_npc_memories_npc_time_idx"),
		index.Fields("source_event_id", "npc_id").StorageKey("game_town_npc_memories_event_npc_idx"),
		index.Fields("embedding_status", "id").StorageKey("game_town_npc_memories_embedding_status_idx"),
	}
}
