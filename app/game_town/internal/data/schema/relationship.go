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

// Relationship 保存玩家与 NPC 的关系。
type Relationship struct{ ent.Schema }

func (Relationship) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "relationships"}, entsql.WithComments(true)}
}
func (Relationship) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("player_id").Comment("玩家 ID"),
		field.Int64("npc_id").Comment("NPC ID"),
		field.Int32("affinity").Comment("亲近程度").Default(50),
		field.Int32("trust").Comment("信任程度").Default(50),
		field.Int32("tension").Comment("冲突程度").Default(0),
		field.JSON("custom_metrics", map[string]any{}).Comment("自定义关系指标").Optional(),
		field.Time("last_interaction_at").Comment("最近互动时间").Nillable().Optional(),
	}
}
func (Relationship) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }
func (Relationship) Indexes() []ent.Index {
	return []ent.Index{index.Fields("world_id", "player_id", "npc_id").Unique().StorageKey("game_town_relationships_world_player_npc_unique")}
}
func (Relationship) Edges() []ent.Edge { return nil }
