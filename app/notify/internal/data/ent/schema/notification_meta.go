package schema

import (
	"common/pkg/constant"
	commonModel "common/pkg/model"
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
		entsql.Annotation{Table: constant.TablePrefixMsgCenter.String() + "notification_meta"},
	}
}

// Fields 定义表字段
func (NotificationMeta) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("uuid").Comment("唯一标识 幂等").Unique(),
		field.String("notification_type").Comment("通知类型"),
		field.Int64("sender_id").Comment("发送者ID"),
		field.JSON("meta", commonModel.Meta{}).Comment("通知元数据"),
		field.String("title").Comment("渲染标题").Default(""),
		field.String("content").Comment("渲染内容"),
		field.Bool("is_global").Comment("是否全站广播").Default(false),
		field.String("status").Comment("状态 1-正常"),
	}
	fields = append(fields, util.TimeAuditFields()...)
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
