package schema

import (
	"common/pkg"
	"common/pkg/constant"

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
		field.Int32("notification_type").Comment("通知类型"),
		field.Int64("sender_id").Comment("发送者ID"),
		field.JSON("meta", map[string]any{}).Comment("通知元数据"),
		field.String("content").Comment("渲染内容"),
		field.Int32("status").Comment("状态 0-正常 1-被取消"),
	}
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (NotificationMeta) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("notification_type", "status"),
		index.Fields("sender_id"),
	}
}

func (NotificationMeta) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联通知记录 一对多
		edge.To("notification_records", NotificationRecord.Type),
	}
}
