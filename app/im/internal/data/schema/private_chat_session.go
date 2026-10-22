package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PrivateChatSession struct{ ent.Schema }

func (PrivateChatSession) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixIM.String() + "sessions"}, entsql.WithComments(true)}
}

func (PrivateChatSession) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("会话所属账号 ID"),
		field.Int64("peer_user_id").Comment("私聊对端账号 ID"),
		field.Int64("last_message_id").Comment("最后消息 ID").Optional().Nillable(),
		field.Time("last_message_at").Comment("最后消息时间").Optional().Nillable(),
		field.Int64("last_read_message_id").Comment("最后已读消息 ID").Optional().Nillable(),
		field.Uint32("unread_count").Comment("未读消息数").Default(0),
		field.Bool("is_muted").Comment("是否免打扰").Default(false),
		field.Bool("is_pinned").Comment("是否置顶").Default(false),
	}
}

func (PrivateChatSession) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (PrivateChatSession) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "peer_user_id").Unique(),
		index.Fields("user_id", "is_pinned", "last_message_at"),
	}
}
