package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TaskVersion 保存定时任务配置历史快照。
type TaskVersion struct {
	ent.Schema
}

func (TaskVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixScheduler.String() + "task_versions"},
		entsql.WithComments(true),
	}
}

func (TaskVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("task_id").Comment("任务 ID"),
		field.Int64("version").Comment("任务配置版本"),
		field.String("name").Comment("代码内任务名称快照").MaxLen(128).NotEmpty(),
		field.String("title").Comment("展示名称快照").MaxLen(128).NotEmpty(),
		field.String("description").Comment("任务说明快照").MaxLen(512).Default(""),
		field.Bool("enabled").Comment("是否启用调度"),
		field.String("cron_spec").Comment("cron 表达式").MaxLen(128).NotEmpty(),
		field.Text("payload").Comment("JSON 编码后的调度参数"),
		field.Int32("timeout_seconds").Comment("单次执行等待超时时间，单位秒").Positive(),
		field.Bool("allow_overlap").Comment("是否允许同一任务重叠执行"),
		field.Bool("alert_enabled").Comment("是否启用告警"),
	}
}

func (TaskVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (TaskVersion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task_id", "version").Unique().StorageKey("scheduler_task_versions_task_version_unique"),
		index.Fields("task_id").StorageKey("scheduler_task_versions_task_idx"),
	}
}

func (TaskVersion) Edges() []ent.Edge { return nil }
