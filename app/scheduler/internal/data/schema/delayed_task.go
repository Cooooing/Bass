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

// DelayedTask 保存延迟任务当前配置。
type DelayedTask struct {
	ent.Schema
}

func (DelayedTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixScheduler.String() + "delayed_tasks"},
		entsql.WithComments(true),
	}
}

func (DelayedTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("代码内任务名").MaxLen(128).NotEmpty(),
		field.String("title").Comment("展示标题").MaxLen(128).NotEmpty(),
		field.String("description").Comment("任务说明").MaxLen(512).Default(""),
		field.Bool("enabled").Comment("是否启用").Default(false),
		field.Int32("timeout_seconds").Comment("单次执行超时秒数").Default(30).Positive(),
		field.Int32("stale_after_seconds").Comment("超过计划时间后的过期判断秒数，空值表示不过期").Optional().Nillable().Positive(),
		field.Int32("max_attempts").Comment("最大尝试次数").Default(3).Positive(),
		field.Enum("misfire_policy").Values(schedulerenum.TaskMisfirePolicyMap.EnumValues()...).Default(schedulerenum.TaskMisfirePolicyExecuteAll.String()).Comment("错过计划时间后的处理策略"),
		field.Int64("version").Comment("配置版本").Default(1),
	}
}

func (DelayedTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (DelayedTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique().StorageKey("scheduler_delayed_tasks_title_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("enabled", "id").StorageKey("scheduler_delayed_tasks_enabled_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("updated_at").StorageKey("scheduler_delayed_tasks_updated_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}

func (DelayedTask) Edges() []ent.Edge {
	return nil
}
