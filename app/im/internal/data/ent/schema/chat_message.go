package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"im/internal/enum"

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
		field.Int64("session_id").Comment("所属会话 ID").Optional().Nillable(),
		field.Enum("type").Values(enum.MessageTypeMap.EnumValues()...).Default(string(enum.MessageTypeNormal)).Comment("消息内容类型"),
		field.Text("content").Comment("消息内容"),
		field.Enum("status").Values(enum.MessageStatusMap.EnumValues()...).Default(string(enum.MessageStatusNormal)).Comment("消息状态"),
		field.Int64("created_by").Comment("创建人ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人ID").Nillable().Optional(),
	}
	return fields
}

func (ChatMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联群组 多对一
		edge.From("group", ChatGroup.Type).Ref("group_messages").Field("group_id").Unique(),
		// 关联群组最后一条消息 一对一
		edge.To("last_message_group", ChatGroup.Type).Unique(),
		// 关联私聊 多对一
		edge.From("session", ChatSession.Type).Ref("session_messages").Field("session_id").Unique(),
		// 关联私聊最后一条消息 一对一
		edge.To("last_message_session", ChatSession.Type).Unique(),
	}
}

func (ChatMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sender_id"),
		index.Fields("group_id"),
		index.Fields("session_id"),
	}
}
