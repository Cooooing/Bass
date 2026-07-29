package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	schedulerenum "scheduler/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ScheduledTaskExecutionRecord 保存定时任务执行记录。
type ScheduledTaskExecutionRecord struct {
	ent.Schema
}

func (ScheduledTaskExecutionRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixScheduler.String() + "scheduled_task_execution_records"},
		entsql.WithComments(true),
	}
}

func (ScheduledTaskExecutionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("scheduled_task_id").Comment("定时任务 ID"),
		field.Int64("scheduled_task_version").Comment("定时任务配置版本"),
		field.Enum("trigger_type").Values(schedulerenum.TaskTriggerTypeMap.EnumValues()...).Default(string(schedulerenum.TaskTriggerTypeSchedule)).Comment("触发类型"),
		field.String("schedule_key").Comment("调度消息唯一键").MaxLen(128).Default(""),
		field.Time("scheduled_at").Comment("计划执行时间"),
		field.Time("started_at").Comment("开始执行时间").Optional().Nillable(),
		field.Time("finished_at").Comment("结束执行时间").Optional().Nillable(),
		field.Int64("duration_ms").Comment("执行耗时毫秒").Optional().Nillable().NonNegative(),
		field.Enum("status").Values(schedulerenum.TaskExecutionStatusMap.EnumValues()...).Comment("执行状态"),
		field.Int32("attempt").Comment("当前尝试次数").Default(0).NonNegative(),
		field.Int32("max_attempts").Comment("最大尝试次数").Default(1).Positive(),
		field.Int32("timeout_seconds").Comment("执行超时秒数").Default(300).Positive(),
		field.Int32("stale_after_seconds").Comment("超过计划时间后的过期判断秒数，空值表示不过期").Optional().Nillable().Positive(),
		field.Enum("misfire_policy").Values(schedulerenum.TaskMisfirePolicyMap.EnumValues()...).Default(schedulerenum.TaskMisfirePolicyExecuteLatest.String()).Comment("错过计划时间后的处理策略"),
		field.String("worker_id").Comment("执行节点标识").MaxLen(128).Default(""),
		field.Text("payload").Comment("JSON 编码后的执行参数快照"),
		field.Text("last_error").Comment("最近错误摘要").Default(""),
		field.String("trace_id").Comment("链路追踪 ID").MaxLen(64).Default(""),
	}
}

func (ScheduledTaskExecutionRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (ScheduledTaskExecutionRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("scheduled_task_id", "trigger_type", "schedule_key").Unique().StorageKey("scheduler_scheduled_task_execution_records_task_schedule_key_unique"),
		index.Fields("scheduled_task_id", "status").StorageKey("scheduler_scheduled_task_execution_records_task_status_idx"),
		index.Fields("scheduled_task_id", "trigger_type", "scheduled_at").StorageKey("scheduler_scheduled_task_execution_records_task_trigger_scheduled_idx"),
		index.Fields("status", "updated_at").StorageKey("scheduler_scheduled_task_execution_records_status_updated_idx"),
	}
}

func (ScheduledTaskExecutionRecord) Edges() []ent.Edge {
	return nil
}
