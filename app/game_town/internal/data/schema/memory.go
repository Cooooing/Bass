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

// Memory 保存 NPC 关于玩家的记忆。
type Memory struct{ ent.Schema }

func (Memory) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "memories"}, entsql.WithComments(true)}
}
func (Memory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("player_id").Comment("玩家 ID"),
		field.Int64("npc_id").Comment("NPC ID"),
		field.String("type").Comment("记忆类型").MaxLen(32).NotEmpty(),
		field.Text("content").Comment("记忆内容"),
		field.Int32("importance").Comment("重要程度").Default(50),
		field.Int64("source_event_id").Comment("来源事件 ID").Nillable().Optional(),
		field.Time("last_recalled_at").Comment("最近召回时间").Nillable().Optional(),
		field.Time("expires_at").Comment("过期时间").Nillable().Optional(),
	}
}
func (Memory) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}, utilent.SoftDeleteMixin{}}
}
func (Memory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "player_id", "npc_id", "type").StorageKey("game_town_memories_world_player_npc_type_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("source_event_id").StorageKey("game_town_memories_source_event_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
func (Memory) Edges() []ent.Edge { return nil }
