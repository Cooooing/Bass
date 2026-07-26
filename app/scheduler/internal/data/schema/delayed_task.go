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

type DelayedTask struct{ ent.Schema }

func (DelayedTask) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{
		Table: constant.TablePrefixScheduler.String() + "delayed_tasks",
	}, entsql.WithComments(true)}
}

func (DelayedTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("idempotency_key").Comment("幂等键").MaxLen(256).NotEmpty(),
		field.String("task_name").Comment("代码内任务名称").MaxLen(128).NotEmpty(),
		field.Text("payload").Comment("JSON 编码后的执行参数"),
		field.Time("execute_at").Comment("期望执行时间"),
		field.Time("next_run_at").Comment("下次执行时间"),
		field.Enum("status").Values(schedulerenum.DelayedTaskStatusMap.EnumValues()...).Default(string(schedulerenum.DelayedTaskStatusPending)).Comment("任务状态"),
		field.Int32("attempt").Comment("已尝试次数").Default(0).NonNegative(),
		field.Int32("max_attempts").Comment("最大尝试次数").Default(3).Positive(),
		field.Int32("timeout_seconds").Comment("单次执行超时秒数").Default(30).Positive(),
		field.String("locked_by").Comment("锁定 worker").MaxLen(128).Default(""),
		field.Time("lock_expires_at").Comment("锁过期时间").Optional().Nillable(),
		field.Time("started_at").Comment("开始时间").Optional().Nillable(),
		field.Time("finished_at").Comment("结束时间").Optional().Nillable(),
		field.Text("last_error").Comment("最近错误摘要").Default(""),
	}
}

func (DelayedTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (DelayedTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("idempotency_key").Unique().StorageKey("scheduler_delayed_tasks_idempotency_unique"),
		index.Fields("status", "next_run_at").StorageKey("scheduler_delayed_tasks_status_next_run_idx"),
		index.Fields("task_name", "created_at").StorageKey("scheduler_delayed_tasks_name_created_idx"),
	}
}

func (DelayedTask) Edges() []ent.Edge {
	return nil
}
