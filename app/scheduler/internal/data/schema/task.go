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

// Task 保存当前生效的定时任务配置。
type Task struct {
	ent.Schema
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixScheduler.String() + "tasks"},
		entsql.WithComments(true),
	}
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("代码内任务名称").MaxLen(128).NotEmpty(),
		field.String("title").Comment("展示名称").MaxLen(128).NotEmpty(),
		field.String("description").Comment("任务说明").MaxLen(512).Default(""),
		field.Bool("enabled").Comment("是否启用调度").Default(false),
		field.String("cron_spec").Comment("cron 表达式").MaxLen(128).NotEmpty(),
		field.Text("payload").Comment("JSON 编码后的调度参数"),
		field.Int32("timeout_seconds").Comment("单次执行等待超时时间，单位秒").Default(300).Positive(),
		field.Bool("allow_overlap").Comment("是否允许同一任务重叠执行").Default(false),
		field.Bool("alert_enabled").Comment("是否启用告警").Default(true),
		field.Int64("version").Comment("任务配置版本").Default(1),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique().StorageKey("scheduler_tasks_title_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("enabled", "id").StorageKey("scheduler_tasks_enabled_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}

func (Task) Edges() []ent.Edge { return nil }
