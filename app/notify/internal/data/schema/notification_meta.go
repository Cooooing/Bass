package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	notifyenum "notify/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotificationMeta 站内信内容表。
type NotificationMeta struct {
	ent.Schema
}

func (NotificationMeta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "notification_meta"},
	}
}

// Fields 定义表字段
func (NotificationMeta) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("title").Comment("站内信标题").Default(""),
		field.String("content").Comment("站内信内容"),
		field.Enum("status").Values(notifyenum.NotificationStatusMap.EnumValues()...).Default(string(notifyenum.NotificationStatusNormal)).Comment("状态"),
	}
}

func (NotificationMeta) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationMeta) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "id"),
	}
}

func (NotificationMeta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notification_records", NotificationRecord.Type),
	}
}
