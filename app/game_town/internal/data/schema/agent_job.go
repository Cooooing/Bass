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

// AgentJob 定义持久化模型任务。
type AgentJob struct {
	ent.Schema
}

func (AgentJob) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "agent_jobs",
		},
		entsql.WithComments(true),
	}
}

func (AgentJob) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("source_event_id").Comment("来源事件 ID"),
		field.Enum("type").Values(gameenum.AgentJobTypeMap.EnumValues()...).Comment("任务类型"),
		field.Enum("priority").Values(gameenum.AgentJobPriorityMap.EnumValues()...).Comment("任务优先级"),
		field.String("lane_key").Comment("串行执行通道").MaxLen(128).NotEmpty(),
		field.Enum("status").Values(gameenum.AgentJobStatusMap.EnumValues()...).Default(string(gameenum.AgentJobStatusQueued)).Comment("任务状态"),
		field.Int64("world_version").Comment("创建任务时的世界版本"),
		field.Int64("npc_id").Comment("关联 NPC ID").Nillable().Optional(),
		field.Int32("attempt_count").Comment("模型调用次数").Default(0),
		field.Time("available_at").Comment("下次可执行时间"),
		field.Time("started_at").Comment("本次开始时间").Nillable().Optional(),
		field.Time("finished_at").Comment("最终完成时间").Nillable().Optional(),
		field.String("error_summary").Comment("最近错误摘要").MaxLen(512).Default(""),
	}
}

func (AgentJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (AgentJob) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source_event_id", "type").
			Unique().
			StorageKey("game_town_agent_jobs_source_event_type_unique"),
		index.Fields("status", "priority", "available_at", "id").
			StorageKey("game_town_agent_jobs_schedule_idx"),
		index.Fields("world_id", "lane_key", "status").
			StorageKey("game_town_agent_jobs_world_lane_status_idx"),
	}
}
