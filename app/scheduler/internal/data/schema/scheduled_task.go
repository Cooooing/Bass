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

// ScheduledTask 保存定时任务当前配置。
type ScheduledTask struct {
	ent.Schema
}

func (ScheduledTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixScheduler.String() + "scheduled_tasks"},
		entsql.WithComments(true),
	}
}

func (ScheduledTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("代码内任务名").MaxLen(128).NotEmpty(),
		field.String("title").Comment("展示标题").MaxLen(128).NotEmpty(),
		field.String("description").Comment("任务说明").MaxLen(512).Default(""),
		field.Bool("enabled").Comment("是否启用调度").Default(false),
		field.String("cron_spec").Comment("cron 表达式").MaxLen(128).NotEmpty(),
		field.Text("payload").Comment("JSON 编码后的默认执行参数"),
		field.Int32("timeout_seconds").Comment("单次执行超时秒数").Default(300).Positive(),
		field.Int32("stale_after_seconds").Comment("超过计划时间后的过期判断秒数，空值表示不过期").Optional().Nillable().Positive(),
		field.Int32("max_attempts").Comment("最大尝试次数").Default(1).Positive(),
		field.Enum("misfire_policy").Values(schedulerenum.TaskMisfirePolicyMap.EnumValues()...).Default(schedulerenum.TaskMisfirePolicyExecuteLatest.String()).Comment("错过计划时间后的处理策略"),
		field.Bool("allow_overlap").Comment("是否允许重叠执行").Default(false),
		field.Int64("version").Comment("配置版本").Default(1),
	}
}

func (ScheduledTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (ScheduledTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique().StorageKey("scheduler_scheduled_tasks_title_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("enabled", "id").StorageKey("scheduler_scheduled_tasks_enabled_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("updated_at").StorageKey("scheduler_scheduled_tasks_updated_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}

func (ScheduledTask) Edges() []ent.Edge {
	return nil
}
