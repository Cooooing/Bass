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

// ScheduledTaskVersion 保存定时任务配置历史快照。
type ScheduledTaskVersion struct {
	ent.Schema
}

func (ScheduledTaskVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixScheduler.String() + "scheduled_task_versions"},
		entsql.WithComments(true),
	}
}

func (ScheduledTaskVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("scheduled_task_id").Comment("定时任务 ID"),
		field.Int64("version").Comment("配置版本"),
		field.String("name").Comment("代码内任务名称快照").MaxLen(128).NotEmpty(),
		field.String("title").Comment("展示标题快照").MaxLen(128).NotEmpty(),
		field.String("description").Comment("任务说明快照").MaxLen(512).Default(""),
		field.Bool("enabled").Comment("是否启用调度快照"),
		field.String("cron_spec").Comment("cron 表达式快照").MaxLen(128).NotEmpty(),
		field.Text("payload").Comment("JSON 编码后的默认执行参数快照"),
		field.Int32("timeout_seconds").Comment("单次执行超时秒数快照").Positive(),
		field.Int32("stale_after_seconds").Comment("过期判断秒数快照，空值表示不过期").Optional().Nillable().Positive(),
		field.Int32("max_attempts").Comment("最大尝试次数快照").Positive(),
		field.Enum("misfire_policy").Values(schedulerenum.TaskMisfirePolicyMap.EnumValues()...).Comment("错过计划时间后的处理策略快照"),
		field.Bool("allow_overlap").Comment("是否允许重叠执行快照"),
	}
}

func (ScheduledTaskVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (ScheduledTaskVersion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("scheduled_task_id", "version").Unique().StorageKey("scheduler_scheduled_task_versions_task_version_unique"),
		index.Fields("scheduled_task_id").StorageKey("scheduler_scheduled_task_versions_task_idx"),
	}
}

func (ScheduledTaskVersion) Edges() []ent.Edge {
	return nil
}
