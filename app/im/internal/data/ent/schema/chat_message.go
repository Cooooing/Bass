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
		field.String("session_id").Comment("会话全局id").NotEmpty(),
		field.Int64("sender_id").Comment("发送者id"),
		field.Int64("target_id").Comment("接收者id（用户id或者群组id）").Optional(),
		field.Int32("type").Comment("消息类型").Default(1),
		field.Text("content").Comment("消息内容").NotEmpty(),
		field.Int32("status").Comment("状态: 0-正常, 1-撤回").Default(0),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	fields = append(fields, pkg.UsernameAuditFields()...)
	return fields
}

func (ChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		// 消息所属的群 (M2O)
		edge.From("group", ChatGroup.Type).Ref("messages").Unique().Field("target_id"),
		// 作为群组的最后一条消息 (O2O 反向)
		edge.From("last_msg_of_group", ChatGroup.Type).Ref("last_message").Unique(),
		// 作为私聊会话的最后一条消息 (O2O 反向)
		edge.From("last_msg_of_session", ChatSession.Type).Ref("last_message").Unique(),
	}
}

func (ChatMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("session_id", "id"),
	}
}
