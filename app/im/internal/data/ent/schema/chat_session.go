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
		field.Int64("receiver_id").Comment("私聊接收者 ID (仅私聊有值)").Optional().Nillable(),
		field.Int64("group_id").Comment("群组ID (仅群聊有值)").Optional().Nillable(),
		field.Bool("is_muted").Comment("是否免打扰").Default(false),
		field.Bool("is_pinned").Comment("是否置顶").Default(false),
		field.Int64("last_read_message_id").Comment("已读最后一条消息 ID").Optional().Nillable(),
		field.Uint32("read_count").Comment("已读数").Default(0),

		/*
			以下字段，仅私聊时值有意义。
			session 表承担群聊私聊会话进度（已读）
			私聊信息（消息总数与最新消息）的管理（O(2)可以接受），单独建 group 同级表管理信息会增加额外复杂度，弊大于利
			群聊的信息（消息总数与最新消息）为避免写扩散上移到 group 表
		*/
		field.Uint32("message_count").Comment("私聊消息数").Default(0),
		field.Int64("last_message_id").Comment("最后消息 ID").Optional().Nillable(),
		field.Int64("created_by").Comment("创建人ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人ID").Nillable().Optional(),
	}
	return fields
}

func (ChatSession) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ChatSession) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联群组 多对一
		edge.From("group", ChatGroup.Type).Ref("group_sessions").Unique().Field("group_id"),
		// 关联消息 一对多
		edge.To("session_messages", ChatMessage.Type),
		// 关联最后一条群消息 一对一
		edge.From("last_message_of_session", ChatMessage.Type).Ref("last_message_session").Unique().Field("last_message_id"),
	}
}

func (ChatSession) Indexes() []ent.Index {
	return []ent.Index{
		// 保证每个用户对每个群组只有一个会话
		index.Fields("created_by", "group_id").Unique(),
		// 保证每个用户对每个私聊接收者只有一个会话
		index.Fields("created_by", "receiver_id").Unique(),
	}
}
