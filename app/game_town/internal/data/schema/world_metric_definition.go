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

// WorldMetricDefinition 保存世界可演化指标定义。
type WorldMetricDefinition struct{ ent.Schema }

func (WorldMetricDefinition) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "world_metric_definitions"}, entsql.WithComments(true)}
}
func (WorldMetricDefinition) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.String("key").Comment("指标键").MaxLen(64).NotEmpty(),
		field.String("name").Comment("指标名称").MaxLen(128).NotEmpty(),
		field.Text("description").Comment("指标说明"),
		field.Int32("min_value").Comment("最小值"),
		field.Int32("max_value").Comment("最大值"),
		field.Int32("initial_value").Comment("初始值"),
	}
}
func (WorldMetricDefinition) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }
func (WorldMetricDefinition) Indexes() []ent.Index {
	return []ent.Index{index.Fields("world_id", "key").Unique().StorageKey("game_town_world_metric_definitions_world_key_unique")}
}
func (WorldMetricDefinition) Edges() []ent.Edge { return nil }
