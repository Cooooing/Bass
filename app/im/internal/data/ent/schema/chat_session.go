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
		field.Int64("target_id").Comment("对方id或群id"),
		field.Int32("type").Comment("会话类型: 1-私聊, 2-群聊").Default(1),
		field.Int64("last_message_id").Comment("最后一条消息id（仅私聊，群聊避免扩散存储在群组表中）").Default(0),
		field.String("last_message_content").Comment("消息内容（仅私聊，群聊避免扩散存储在群组表中）").Default(""),
		field.Int64("last_read_msg_id").Comment("已读最后一条消息id").Default(0),
		field.Int32("read_count").Comment("已读数").Default(0),
		field.Int32("unread_count").Comment("未读数").Default(0),
		field.Bool("is_muted").Comment("是否免打扰").Default(false),
		field.Bool("is_pinned").Comment("是否置顶").Default(false),
	}
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (ChatSession) Edges() []ent.Edge {
	return []ent.Edge{
		// 会话对应的群组 (M2O)
		edge.From("group", ChatGroup.Type).Ref("sessions").Unique().Field("target_id"),
		// 会话最后一条消息 (O2O)
		edge.To("last_message", ChatMessage.Type).Unique().Field("last_message_id"),
	}
}

func (ChatSession) Indexes() []ent.Index {
	return []ent.Index{
		// 保证每个用户对每个目标只有一个会话
		index.Fields("user_id", "target_id").Unique(),
		// 用于列表按更新时间排序渲染
		index.Fields("user_id", "updated_at"),
	}
}
