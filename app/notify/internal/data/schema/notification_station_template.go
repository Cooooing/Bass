package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotificationStationTemplate 保存站内信模板字段。
type NotificationStationTemplate struct {
	ent.Schema
}

func (NotificationStationTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_station_template",
		},
		entsql.WithComments(true),
	}
}

func (NotificationStationTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("rule_id").Comment("通知规则 ID"),
		field.String("title_template").Comment("标题模板"),
		field.String("content_template").Comment("内容模板"),
	}
}

func (NotificationStationTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (NotificationStationTemplate) Indexes() []ent.Index {
	return []ent.Index{index.Fields("rule_id").Unique()}
}

func (NotificationStationTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rule", NotificationRule.Type).Ref("station_template").Required().Unique().Field("rule_id"),
	}
}
