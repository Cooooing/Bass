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

// TaskExecutionRecord 保存定时任务执行记录。
type TaskExecutionRecord struct {
	ent.Schema
}

func (TaskExecutionRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixScheduler.String() + "task_execution_records",
		},
		entsql.WithComments(true),
	}
}

func (TaskExecutionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("task_id").Comment("任务 ID"),
		field.Int64("task_version").Comment("执行时使用的任务配置版本"),
		field.Time("scheduled_at").Comment("理论调度时间"),
		field.Time("started_at").Comment("实际开始时间").Optional().Nillable(),
		field.Time("finished_at").Comment("实际结束时间").Optional().Nillable(),
		field.Int64("duration_ms").Comment("执行耗时，单位毫秒").Optional().Nillable().NonNegative(),
		field.Enum("status").
			Values(schedulerenum.TaskExecutionStatusMap.EnumValues()...).
			Comment("执行状态"),
		field.Enum("trigger_type").
			Values(schedulerenum.TaskTriggerTypeMap.EnumValues()...).
			Default(string(schedulerenum.TaskTriggerTypeSchedule)).
			Comment("触发类型"),
		field.String("worker_id").Comment("触发本次执行的 scheduler 实例").MaxLen(128).Default(""),
		field.Text("payload").Comment("JSON 编码后的执行参数快照"),
		field.Text("last_error").Comment("最近一次错误摘要").Default(""),
		field.String("trace_id").Comment("链路追踪 ID").MaxLen(64).Default(""),
	}
}

func (TaskExecutionRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (TaskExecutionRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task_id", "scheduled_at").Unique().StorageKey("scheduler_task_execution_records_task_period_unique"),
		index.Fields("task_id", "status").StorageKey("scheduler_task_execution_records_task_status_idx"),
		index.Fields("task_id", "trigger_type", "scheduled_at").StorageKey("scheduler_task_execution_records_task_trigger_scheduled_idx"),
		index.Fields("status", "updated_at").StorageKey("scheduler_task_execution_records_status_updated_idx"),
	}
}

func (TaskExecutionRecord) Edges() []ent.Edge {
	return nil
}
