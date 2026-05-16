package schema

import (
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	utilent "common/pkg/util/ent"
	"common/pkg/enum"
	"common/pkg/util"

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
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("uuid").Comment("唯一标识 幂等").Unique(),
		field.Enum("event_type").Comment("事件类型").
			GoType(enum.EventType("")),
		field.Int64("sender_id").Comment("发送者ID"),
		field.JSON("meta", &v1.NotificationMeta{}).Comment("通知元数据"),
		field.String("title").Comment("渲染标题").Default(""),
		field.String("content").Comment("渲染内容"),
		field.Bool("is_global").Comment("是否全站广播").Default(false),
		field.Enum("status").Comment("状态").
			GoType(enum.NotificationStatus("")),
	}
	return fields
}

func (NotificationMeta) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationMeta) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type", "status"),
		index.Fields("sender_id"),
	}
}

func (NotificationMeta) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联通知记录 一对多
		edge.To("notification_records", NotificationRecord.Type),
	}
}
