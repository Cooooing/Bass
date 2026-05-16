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

// NotificationRecord 通知记录表
type NotificationRecord struct {
	ent.Schema
}

func (NotificationRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "notification_record"},
	}
}

// Fields 定义表字段
func (NotificationRecord) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("notification_id").Comment("通知元数据ID"),
		field.Int64("receiver_id").Comment("接收者ID"),
		field.Time("read_time").Comment("已读时间").Optional().Nillable(),
	}
	return fields
}

func (NotificationRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("receiver_id", "notification_id"),
	}
}

func (NotificationRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联通知元数据 多对一
		edge.From("notification_meta", NotificationMeta.Type).Ref("notification_records").Required().Unique().Field("notification_id"),
	}
}
