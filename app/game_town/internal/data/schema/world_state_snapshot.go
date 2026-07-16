package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// WorldStateSnapshot 保存世界状态快照。
type WorldStateSnapshot struct{ ent.Schema }

func (WorldStateSnapshot) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "world_state_snapshots"}, entsql.WithComments(true)}
}
func (WorldStateSnapshot) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("tick_count").Comment("演化次数"),
		field.String("current_arc").Comment("当前阶段").MaxLen(256).Default(""),
		field.JSON("metrics", map[string]any{}).Comment("当前指标值").Optional(),
		field.Text("summary").Comment("状态摘要").Default(""),
		field.Int64("reason_event_id").Comment("原因事件 ID").Nillable().Optional(),
		field.Time("created_at").Comment("创建时间"),
	}
}
func (WorldStateSnapshot) Indexes() []ent.Index {
	return []ent.Index{index.Fields("world_id", "tick_count").StorageKey("game_town_world_state_snapshots_world_tick_idx")}
}
func (WorldStateSnapshot) Edges() []ent.Edge { return nil }
