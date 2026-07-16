package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Event 保存世界内关键行为和状态变化。
type Event struct{ ent.Schema }

func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "events"}, entsql.WithComments(true)}
}
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.String("type").Comment("事件类型").MaxLen(64).NotEmpty(),
		field.Int64("actor_player_id").Comment("玩家 ID").Nillable().Optional(),
		field.Int64("target_npc_id").Comment("目标 NPC ID").Nillable().Optional(),
		field.Int64("location_id").Comment("地点 ID").Nillable().Optional(),
		field.Int64("command_id").Comment("命令 ID").Nillable().Optional(),
		field.String("summary").Comment("事件摘要").MaxLen(512).Default(""),
		field.Text("content").Comment("事件内容").Default(""),
		field.JSON("effects", map[string]any{}).Comment("影响").Optional(),
		field.JSON("metadata", map[string]any{}).Comment("元数据").Optional(),
		field.Time("occurred_at").Comment("发生时间"),
		field.Time("created_at").Comment("创建时间"),
	}
}
func (Event) Indexes() []ent.Index {
	return []ent.Index{index.Fields("world_id", "type", "occurred_at").StorageKey("game_town_events_world_type_time_idx"), index.Fields("world_id", "actor_player_id", "occurred_at").StorageKey("game_town_events_world_actor_time_idx")}
}
func (Event) Edges() []ent.Edge { return nil }
