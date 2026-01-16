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

// ChatMessage 消息实体定义
type ChatMessage struct {
	ent.Schema
}

func (ChatMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixIM.String() + "chat_messages"},
		entsql.WithComments(true),
	}
}

func (ChatMessage) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("sender_id").Comment("发送者ID"),
		field.Int64("receiver_id").Comment("私聊接收者ID (仅私聊有值)").Optional().Nillable(),
		field.Int64("group_id").Comment("群组ID (仅群聊有值)").Optional().Nillable(),
		field.Int32("type").Comment("消息内容类型").Default(1),
		field.Text("content").Comment("消息内容"),
		field.Int32("status").Comment("状态: 1-正常, 2-撤回").Default(1),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	fields = append(fields, pkg.UsernameAuditFields()...)
	return fields
}

func (ChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联群组 多对一
		edge.From("group", ChatGroup.Type).Ref("group_messages").Field("group_id").Unique(),
		// 关联群组最后一条消息 一对一
		edge.To("last_message_group", ChatGroup.Type).Unique(),
		// 关联私聊 多对一
		edge.From("session", ChatSession.Type).Ref("session_messages").Field("receiver_id").Unique(),
		// 关联私聊最后一条消息 一对一
		edge.To("last_message_session", ChatSession.Type).Unique(),
	}
}

func (ChatMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sender_id"),
		index.Fields("group_id"),
	}
}
