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

type ChatSession struct {
	ent.Schema
}

func (ChatSession) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixIM.String() + "chat_sessions"},
	}
}

func (ChatSession) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("所属用户id"),
		field.Int64("receiver_id").Comment("私聊接收者ID (仅私聊有值)").Optional().Nillable(),
		field.Int64("group_id").Comment("群组ID (仅群聊有值)").Optional().Nillable(),
		field.Int64("last_read_message_id").Comment("已读最后一条消息id").Default(0),
		field.Int32("read_count").Comment("已读数").Default(0),
		field.Int32("message_count").Comment("消息总数（仅私聊有意义）").Default(0),
		field.Bool("is_muted").Comment("是否免打扰").Default(false),
		field.Bool("is_pinned").Comment("是否置顶").Default(false),

		field.Int64("last_message_id").Comment("最后一条消息id 仅私聊，群聊避免写扩散存储在群组表中").Unique().Optional().Nillable(),
		field.String("last_message_content").Comment("最后消息内容 仅私聊，群聊避免写扩散存储在群组表中").Optional().Nillable(),
		field.Time("last_message_at").Comment("最后消息时间 仅私聊，群聊避免写扩散存储在群组表中").Optional().Nillable(),
	}
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (ChatSession) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联群组 多对一
		edge.From("group", ChatGroup.Type).Ref("sessions").Unique().Field("group_id"),
		// 关联最后一条会话消息 一对一
		edge.From("last_message_of_session", ChatMessage.Type).Ref("last_message_session").Unique().Field("last_message_id"),
	}
}

func (ChatSession) Indexes() []ent.Index {
	return []ent.Index{
		// 保证每个用户对每个私聊对象只有一个会话
		index.Fields("user_id", "receiver_id").Unique(),
		// 保证每个用户对每个群组只有一个会话
		index.Fields("user_id", "group_id").Unique(),
	}
}
