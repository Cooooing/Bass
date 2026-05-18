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

// NotificationMeta 通知元数据表
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
		field.String("uuid").Comment("唯一标识 幂等").Unique(),
		field.Enum("event_type").Values(notifyenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.JSON("meta", []byte{}).Comment("事件 payload JSON"),
		field.String("title").Comment("渲染标题").Default(""),
		field.String("content").Comment("渲染内容"),
		field.Bool("is_global").Comment("是否全站广播").Default(false),
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
		index.Fields("event_type", "status"),
	}
}

func (NotificationMeta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notification_records", NotificationRecord.Type),
	}
}
