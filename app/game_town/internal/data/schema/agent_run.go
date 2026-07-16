package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AgentRun 保存 Agent 调用输入输出，便于调试和回放。
type AgentRun struct{ ent.Schema }

func (AgentRun) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "agent_runs"}, entsql.WithComments(true)}
}
func (AgentRun) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID").Nillable().Optional(),
		field.Int64("agent_config_id").Comment("Agent 配置 ID").Nillable().Optional(),
		field.String("run_type").Comment("运行类型").MaxLen(64).NotEmpty(),
		field.Int64("command_id").Comment("命令 ID").Nillable().Optional(),
		field.Int64("event_id").Comment("事件 ID").Nillable().Optional(),
		field.Int64("npc_id").Comment("NPC ID").Nillable().Optional(),
		field.String("model").Comment("模型名称").MaxLen(128).Default(""),
		field.JSON("input_json", map[string]any{}).Comment("输入 JSON").Optional(),
		field.JSON("output_json", map[string]any{}).Comment("输出 JSON").Optional(),
		field.String("status").Comment("运行状态").MaxLen(32).NotEmpty(),
		field.String("error_summary").Comment("错误摘要").MaxLen(512).Default(""),
		field.Int64("latency_ms").Comment("耗时毫秒").Default(0),
		field.Time("created_at").Comment("创建时间"),
	}
}
func (AgentRun) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "run_type", "created_at").StorageKey("game_town_agent_runs_world_type_time_idx"),
		index.Fields("agent_config_id", "created_at").StorageKey("game_town_agent_runs_config_time_idx"),
	}
}
func (AgentRun) Edges() []ent.Edge { return nil }
