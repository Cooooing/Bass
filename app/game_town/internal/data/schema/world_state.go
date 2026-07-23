package schema

import (
	"time"

	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type WorldState struct {
	ent.Schema
}

func (WorldState) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "world_states"},
		entsql.WithComments(true),
	}
}

func (WorldState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (WorldState) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("version").Comment("世界状态版本").Default(0),
		field.Uint64("event_sequence").Comment("已分配事件序号").Default(0),
		field.Uint64("agent_cursor").Comment("Agent 已消费事件序号").Default(0),
		field.Text("summary").Comment("内部世界叙事摘要").Default(""),
		field.String("current_arc").Comment("当前剧情阶段").MaxRuneLen(256).Default(""),
		field.Text("public_chronicle").Comment("玩家可知的公开纪事").Default(""),
		field.String("current_era").Comment("当前纪元或阶段").MaxRuneLen(128).Default(""),
		field.Time("world_time").Comment("已物化的世界时间").Default(time.Now),
		field.Time("time_anchor").Comment("计算世界时间的现实锚点").Default(time.Now),
		field.Uint32("time_scale").Comment("现实一小时对应的世界小时数").Default(24),
		field.Int64("rule_version").Comment("世界规则版本").Default(1),
		field.Time("next_tick_at").Comment("兼容旧世界演进调度").Nillable().Optional(),
		field.Time("next_due_at").Comment("下一次到期世界动作时间").Nillable().Optional(),
	}
}

func (WorldState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id").Unique().StorageKey("game_town_world_states_world_unique"),
		index.Fields("next_tick_at").StorageKey("game_town_world_states_next_tick_idx"),
		index.Fields("next_due_at").StorageKey("game_town_world_states_next_due_idx"),
	}
}
