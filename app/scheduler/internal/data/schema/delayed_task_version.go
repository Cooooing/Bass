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

// DelayedTaskVersion 保存延迟任务配置历史快照。
type DelayedTaskVersion struct {
	ent.Schema
}

func (DelayedTaskVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixScheduler.String() + "delayed_task_versions"},
		entsql.WithComments(true),
	}
}

func (DelayedTaskVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("delayed_task_id").Comment("延迟任务 ID"),
		field.Int64("version").Comment("配置版本"),
		field.String("task_key").Comment("任务配置唯一键快照").MaxLen(128).NotEmpty(),
		field.String("handler_name").Comment("代码内任务处理器快照").MaxLen(128).NotEmpty(),
		field.String("title").Comment("展示标题快照").MaxLen(128).NotEmpty(),
		field.String("description").Comment("任务说明快照").MaxLen(512).Default(""),
		field.Bool("enabled").Comment("是否启用快照"),
		field.Int32("timeout_seconds").Comment("单次执行超时秒数快照").Positive(),
		field.Int32("stale_after_seconds").Comment("过期判断秒数快照，空值表示不过期").Optional().Nillable().Positive(),
		field.Int32("max_attempts").Comment("最大尝试次数快照").Positive(),
		field.Enum("misfire_policy").Values(schedulerenum.TaskMisfirePolicyMap.EnumValues()...).Comment("错过计划时间后的处理策略快照"),
	}
}

func (DelayedTaskVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (DelayedTaskVersion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("delayed_task_id", "version").Unique().StorageKey("scheduler_delayed_task_versions_task_version_unique"),
		index.Fields("delayed_task_id").StorageKey("scheduler_delayed_task_versions_task_idx"),
	}
}

func (DelayedTaskVersion) Edges() []ent.Edge {
	return nil
}
